package commandexecutor

import (
	"sync"
)

//go:generate mockgen -source=$GOFILE -destination=mocks/server_generated.go -package commandexecutormocks ICmdExecutor,IMetricsCollector

type CommandID string

func (c CommandID) S() string {
	return string(c)
}

type CommandResult struct {
	ID   CommandID
	Data interface{}
	Err  error
}

type ICmdExecutor interface {
	// ResultsCh возвращает канал с потоком результатов выполненных команд.
	ResultsCh() <-chan CommandResult
	// Exec выполняет команду, результат приходит в ResultsCh.
	Exec(cid CommandID) error
}

type IMetricsCollector interface {
	CountCommand(cid CommandID) error
	CountError(err error) error
}

type Server struct {
	executor ICmdExecutor
	metrics  IMetricsCollector
	wg       sync.WaitGroup
}

func NewServer(exec ICmdExecutor, mColl IMetricsCollector) *Server {
	// Реализуй меня.
	return nil
}

func (s *Server) Wait() {
	s.wg.Wait()
}

// ProcessCommandsStream получает команды из транспорта, отдаёт их на выполнение,
// а параллельно с этим: слушает канал результатов команд, считает по ним метрики,
// на основе CommandResult формирует ProtoCommandResult и отправляет назад в транспорт.
func (s *Server) ProcessCommandsStream(t ITransport) error {
	ctx := t.Context()
	resultsCh := s.executor.ResultsCh()

	// Реализуй цикл обработки результатов команд из resultsCh и отправки их в транспорт.
	// Реализуй цикл получения команд из транспорта.

	_, _ = ctx, resultsCh
	return nil
}
