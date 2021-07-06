package docker

import "context"

type Executor interface {
	Exec(ctx context.Context, cmd string, args ...interface{}) error
}

type Docker struct{}

func (d *Docker) RunContainer(ctx context.Context, e Executor, image string) error {
	if err := e.Exec(ctx, "run", image); err != nil {
		return newDockerError(err)
	}
	return nil
}

func (d *Docker) StopContainer(ctx context.Context, e Executor, containerID string) error {
	if err := e.Exec(ctx, "stop", containerID); err != nil {
		return newDockerError(err)
	}
	return nil
}

func (d *Docker) ExecContainerCmd(ctx context.Context, e Executor, containerID, cmd string) error {
	if err := e.Exec(ctx, "exec", containerID, cmd); err != nil {
		return newDockerError(err)
	}
	return nil
}
