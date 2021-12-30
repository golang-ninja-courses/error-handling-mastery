package chain

import (
	"fmt"
	"io"
)

type Message struct {
	ID string
}

func ProcessMessage() error { // Почини возвращаемую цепочку ошибок!
	msg := readMessageFromQueue()

	if err := process(msg); err != nil {
		return fmt.Errorf("cannot process msg: %w", err)
	}

	return nil
}

func readMessageFromQueue() Message {
	return Message{ID: "8fbad38c-c5c5-11eb-b876-1e00d13a7870"}
}

func process(msg Message) error {
	if err := saveMsg(msg); err != nil {
		return fmt.Errorf("cannot write data: %v", err)
	}
	return nil
}

type saveMsgError struct {
	id  string
	err error
}

func (w *saveMsgError) Error() string {
	return fmt.Sprintf("save msg %q error: %v", w.id, w.err)
}

func saveMsg(m Message) error {
	if true {
		return &saveMsgError{
			id:  m.ID,
			err: fmt.Errorf("%w", io.ErrShortWrite),
		}
	}
	return nil
}
