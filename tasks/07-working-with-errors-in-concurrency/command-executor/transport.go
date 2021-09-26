package commandexecutor

import "context"

//go:generate mockgen -source=$GOFILE -destination=mocks/transport_generated.go -package commandexecutormocks ITransport

type ProtoCommand struct {
	ID string
}

type ProtoCommandStatus int

const (
	ProtoCommandStatusUnknownError ProtoCommandStatus = iota
	ProtoCommandStatusTimeoutError
	ProtoCommandStatusUnsupportedCommandError
	ProtoCommandStatusSuccess
)

type ProtoCommandResult struct {
	ID     string
	Status ProtoCommandStatus
}

type ITransport interface {
	Context() context.Context
	Recv() (*ProtoCommand, error)
	Send(c *ProtoCommandResult) error
}
