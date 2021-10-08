package commandexecutor_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/goleak"

	ce "github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/tasks/07-working-with-errors-in-concurrency/command-executor"
	cemocks "github.com/www-golang-courses-ru/advanced-dealing-with-errors-in-go/tasks/07-working-with-errors-in-concurrency/command-executor/mocks"
)

// go test -race -count 10 .

type TestServerSuite struct {
	suite.Suite

	ctx    context.Context
	cancel context.CancelFunc

	ctrl *gomock.Controller

	t     *cemocks.MockITransport
	exec  *cemocks.MockICmdExecutor
	mColl *cemocks.MockIMetricsCollector
	srv   *ce.Server
}

func TestServer(t *testing.T) {
	suite.Run(t, new(TestServerSuite))
}

func (s *TestServerSuite) SetupTest() {
	s.ctx, s.cancel = context.WithCancel(context.Background())

	s.ctrl = gomock.NewController(s.T())

	s.t = cemocks.NewMockITransport(s.ctrl)
	s.exec = cemocks.NewMockICmdExecutor(s.ctrl)
	s.mColl = cemocks.NewMockIMetricsCollector(s.ctrl)
	s.srv = ce.NewServer(s.exec, s.mColl)
}

func (s *TestServerSuite) TearDownTest() {
	s.cancel()
	s.srv.Wait()
	s.ctrl.Finish()
	goleak.VerifyNone(s.T())
}

func (s *TestServerSuite) TestFlow() {
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	s.t.EXPECT().Context().AnyTimes().Return(ctx)

	const (
		n             = 10
		executionTime = time.Second
	)

	// Эмулируем получение от транспорта n команд и ждём момента,
	// чтобы послать завершающий io.EOF.
	var i int
	s.t.EXPECT().Recv().Times(n + 1).DoAndReturn(func() (*ce.ProtoCommand, error) {
		i++
		if i == n+1 {
			<-ctx.Done()
			return nil, io.EOF
		}
		return &ce.ProtoCommand{ID: strconv.Itoa(i)}, nil
	})

	// Отправляем в транспорт n результатов команд и после этого
	// сигнализируем, что транспорту можно послать нам io.EOF.
	protoResults := make([]*ce.ProtoCommandResult, 0, n)
	s.t.EXPECT().Send(gomock.Any()).Times(n).DoAndReturn(func(c *ce.ProtoCommandResult) error {
		protoResults = append(protoResults, c)
		if len(protoResults) == n {
			cancel()
		}
		return nil
	})

	resultsCh := make(chan ce.CommandResult)
	s.exec.EXPECT().ResultsCh().AnyTimes().Return(resultsCh)
	s.exec.EXPECT().Exec(gomock.Any()).Times(n).DoAndReturn(func(cid ce.CommandID) error {
		var err error
		var data interface{}

		switch cid.S() {
		case "1":
			data = 1 // Успешная команда.
		case "2":
			// Успешная команда без результата.
		case "3":
			err = errors.New("unknown error")
		case "4":
			err = fmt.Errorf("too long: %w", ce.ErrCommandTimeout)
		case "5":
			err = fmt.Errorf("confused: %w", ce.ErrUnsupportedCommand)
		case "6":
			data = 2
		case "7", "8", "9", "10":
		default:
			panic("unhandled command id:" + cid.S())
		}

		go func() {
			res := ce.CommandResult{
				ID:   cid,
				Err:  err,
				Data: data,
			}

			select {
			case <-ctx.Done():
				return
			case <-time.After(executionTime):
			}

			select {
			case <-ctx.Done():
			case resultsCh <- res:
			}
		}()

		return nil
	})

	// Ожидаем 2 успешных команды с данными в результате.
	s.mColl.EXPECT().CountCommand(gomock.Any()).Times(2).DoAndReturn(func(cid ce.CommandID) error {
		s.Require().NotEmpty(cid)
		return nil
	})
	// Ожидаем 3 команды, завершившихся с ошибкой.
	s.mColl.EXPECT().CountError(gomock.Any()).Times(3).DoAndReturn(func(err error) error {
		s.Require().NotNil(err)
		return nil
	})

	start := time.Now()
	{
		err := s.srv.ProcessCommandsStream(s.t)
		s.Require().NoError(err)
		s.srv.Wait()
	}
	elapsed := time.Since(start)

	s.Run("commands was executed concurrently", func() {
		s.LessOrEqual(elapsed, executionTime*3) // С погрешностью на среду выполнения.
	})

	s.Run("status processing is ok", func() {
		s.ElementsMatch(protoResults, []*ce.ProtoCommandResult{
			{ID: "1", Status: ce.ProtoCommandStatusSuccess},
			{ID: "2", Status: ce.ProtoCommandStatusSuccess},
			{ID: "3", Status: ce.ProtoCommandStatusUnknownError},
			{ID: "4", Status: ce.ProtoCommandStatusTimeoutError},
			{ID: "5", Status: ce.ProtoCommandStatusUnsupportedCommandError},
			{ID: "6", Status: ce.ProtoCommandStatusSuccess},
			{ID: "7", Status: ce.ProtoCommandStatusSuccess},
			{ID: "8", Status: ce.ProtoCommandStatusSuccess},
			{ID: "9", Status: ce.ProtoCommandStatusSuccess},
			{ID: "10", Status: ce.ProtoCommandStatusSuccess},
		})
	})
}

func (s *TestServerSuite) TestRecvError() {
	s.exec.EXPECT().ResultsCh().AnyTimes().Return(make(chan ce.CommandResult))
	s.t.EXPECT().Context().AnyTimes().Return(s.ctx)
	s.t.EXPECT().Recv().Times(1).Return(nil, errors.New("network error"))

	err := s.srv.ProcessCommandsStream(s.t)
	s.Error(err)
}

func (s *TestServerSuite) TestExecError() {
	const cid = ce.CommandID("1")

	s.t.EXPECT().Context().AnyTimes().Return(s.ctx)
	s.t.EXPECT().Recv().Times(1).Return(&ce.ProtoCommand{ID: cid.S()}, nil)

	s.exec.EXPECT().ResultsCh().AnyTimes().Return(make(chan ce.CommandResult))
	s.exec.EXPECT().Exec(cid).Return(errors.New("exec error"))

	err := s.srv.ProcessCommandsStream(s.t)
	s.Error(err)
}

func (s *TestServerSuite) TestErrFromCountError() {
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	const cid = ce.CommandID("1")

	s.t.EXPECT().Context().AnyTimes().Return(ctx)
	s.t.EXPECT().Recv().Times(1).Return(&ce.ProtoCommand{ID: cid.S()}, nil)
	s.t.EXPECT().Recv().Times(1).DoAndReturn(func() (*ce.ProtoCommand, error) {
		<-ctx.Done()
		return nil, io.EOF
	})

	resultsCh := make(chan ce.CommandResult)
	s.exec.EXPECT().ResultsCh().AnyTimes().Return(resultsCh)

	cmdErr := errors.New("broken pipe")
	s.exec.EXPECT().Exec(cid).DoAndReturn(func(_ ce.CommandID) error {
		go func() {
			res := ce.CommandResult{
				ID:  cid,
				Err: cmdErr,
			}
			select {
			case <-ctx.Done():
			case resultsCh <- res:
			}
		}()
		return nil
	})

	s.mColl.EXPECT().CountError(cmdErr).DoAndReturn(func(_ error) error {
		cancel()
		return errors.New("counter error")
	})

	_ = s.srv.ProcessCommandsStream(s.t)
}

func (s *TestServerSuite) TestErrFromCountCommand() {
	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	const cid = ce.CommandID("1")

	s.t.EXPECT().Context().AnyTimes().Return(ctx)
	s.t.EXPECT().Recv().Times(1).Return(&ce.ProtoCommand{ID: cid.S()}, nil)
	s.t.EXPECT().Recv().Times(1).DoAndReturn(func() (*ce.ProtoCommand, error) {
		<-ctx.Done()
		return nil, io.EOF
	})

	resultsCh := make(chan ce.CommandResult)
	s.exec.EXPECT().ResultsCh().AnyTimes().Return(resultsCh)

	s.exec.EXPECT().Exec(cid).DoAndReturn(func(_ ce.CommandID) error {
		go func() {
			res := ce.CommandResult{
				ID:   cid,
				Data: 42,
			}
			select {
			case <-ctx.Done():
			case resultsCh <- res:
			}
		}()
		return nil
	})

	s.mColl.EXPECT().CountCommand(cid).DoAndReturn(func(_ ce.CommandID) error {
		cancel()
		return errors.New("counter error")
	})

	_ = s.srv.ProcessCommandsStream(s.t)
}
