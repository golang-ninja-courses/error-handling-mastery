package docker

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorImplementsHelpfulMethods(t *testing.T) {
	type helpfulMethods interface {
		IsPullAccessDeniedError() bool
		IsNoSuchContainerError() bool
		IsContainerNotRunningError() bool
	}
	assert.Implements(t, (*helpfulMethods)(nil), new(Error))
}

func TestDocker_RunContainer(t *testing.T) {
	isPullAccessDeniedError := func(err error) bool {
		var i interface {
			IsPullAccessDeniedError() bool
		}
		return err != nil && errors.As(err, &i) && i.IsPullAccessDeniedError()
	}

	cases := []struct {
		name                    string
		image                   string
		errMock                 error
		isPullAccessDeniedError bool
	}{
		{
			name:                    "no error",
			image:                   "redis",
			errMock:                 nil,
			isPullAccessDeniedError: false,
		},
		{
			name:                    "not `pull access denied` error",
			image:                   "-k redis",
			errMock:                 errors.New(`unknown shorthand flag: 'k' in -k. See 'docker run --help'`),
			isPullAccessDeniedError: false,
		},
		{
			// $ docker run rabbitmq:4-management
			// docker: Error response from daemon: pull access denied for rabbitmq:4-management
			name:                    "`pull access denied` error",
			image:                   "rabbitmq:4-management",
			errMock:                 errors.New(`Error response from daemon: pull access denied for rabbitmq:4-management`),
			isPullAccessDeniedError: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var d Docker
			err := d.RunContainer(ctx, newExecutorMock("run", tt.errMock), tt.image)
			assert.Equal(t, tt.isPullAccessDeniedError, isPullAccessDeniedError(err))
		})
	}
}

func TestDocker_StopContainer(t *testing.T) {
	isNoSuchContainerError := func(err error) bool {
		var i interface {
			IsNoSuchContainerError() bool
		}
		return err != nil && errors.As(err, &i) && i.IsNoSuchContainerError()
	}

	cases := []struct {
		name                   string
		containerID            string
		errMock                error
		isNoSuchContainerError bool
	}{
		{
			name:                   "no error",
			containerID:            "7aeae4613083",
			errMock:                nil,
			isNoSuchContainerError: false,
		},
		{
			name:                   "not `no such container` error",
			containerID:            "-k 7aeae4613083",
			errMock:                errors.New(`unknown shorthand flag: 'k' in -k. See 'docker stop --help'`),
			isNoSuchContainerError: false,
		},
		{
			// $ docker stop unknown
			// Error response from daemon: No such container: unknown
			name:                   "`no such container` error",
			containerID:            "unknown",
			errMock:                errors.New(`Error response from daemon: No such container: unknown`),
			isNoSuchContainerError: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var d Docker
			err := d.StopContainer(ctx, newExecutorMock("stop", tt.errMock), tt.containerID)
			assert.Equal(t, tt.isNoSuchContainerError, isNoSuchContainerError(err))
		})
	}
}

func TestDocker_ExecContainerCmd(t *testing.T) {
	isContainerNotRunningError := func(err error) bool {
		var i interface {
			IsContainerNotRunningError() bool
		}
		return err != nil && errors.As(err, &i) && i.IsContainerNotRunningError()
	}

	cases := []struct {
		name                       string
		containerID                string
		errMock                    error
		isContainerNotRunningError bool
	}{
		{
			name:                       "no error",
			containerID:                "7aeae4613083",
			errMock:                    nil,
			isContainerNotRunningError: false,
		},
		{
			name:                       "not `container is not running` error",
			containerID:                "-k 7aeae4613083",
			errMock:                    errors.New(`unknown shorthand flag: 'k' in -k. See 'docker exec --help'`),
			isContainerNotRunningError: false,
		},
		{
			// $ docker exec -it 7aeae4613083 /bin/bash
			// Error response from daemon: Container 7aeae4613083 is not running
			name:                       "`container is not running` error",
			containerID:                "7aeae4613083",
			errMock:                    errors.New(`Error response from daemon: Container 7aeae4613083 is not running`),
			isContainerNotRunningError: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			var d Docker
			err := d.ExecContainerCmd(ctx, newExecutorMock("exec", tt.errMock), tt.containerID, "/bin/bash")
			assert.Equal(t, tt.isContainerNotRunningError, isContainerNotRunningError(err))
		})
	}
}

type executorMock struct {
	expectedCmd string
	expectedErr error
}

var _ Executor = (*executorMock)(nil)

func newExecutorMock(cmd string, err error) executorMock {
	return executorMock{
		expectedCmd: cmd,
		expectedErr: err,
	}
}

func (e executorMock) Exec(ctx context.Context, cmd string, args ...interface{}) error {
	if e.expectedCmd == cmd {
		return e.expectedErr
	}
	return nil
}
