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
	// Context возвращает контекст текущего соединения.
	Context() context.Context

	// Recv позволяет принять команду. При успешном отсоединении клиента вернёт ошибку io.EOF.
	// Любая другая ошибка будет постоянной и при её появлении нужно завершать обработчик
	// транспорта в принципе. Вызов Recv блокирующий.
	//
	// Безопасно иметь одну горутину, вызывающую Recv и другую – вызывающую Send.
	// Но опасно вызывать Recv в разных горутинах.
	Recv() (*ProtoCommand, error)

	// Send позволяет отправить результат команды. Любая ошибка будет постоянной и
	// аналогичной той, что придёт в очередном Recv. Вызов Send блокирующий.
	//
	// Безопасно иметь одну горутину, вызывающую Recv и другую – вызывающую Send.
	// Но опасно вызывать Send в разных горутинах.
	Send(c *ProtoCommandResult) error
}
