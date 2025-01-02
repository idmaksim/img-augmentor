package model

import tea "github.com/charmbracelet/bubbletea"

type MessageHandler interface {
	Handle(Model) (Model, tea.Cmd)
}

type (
	KeyMessage struct {
		msg tea.KeyMsg
	}
	ProcessFinishedMessage struct{}
	ErrorMessage           struct {
		err error
	}
)

func NewKeyHandler(msg tea.KeyMsg) MessageHandler {
	return &KeyMessage{msg: msg}
}

func NewProcessFinishedHandler() MessageHandler {
	return &ProcessFinishedMessage{}
}

func NewErrorHandler(err error) MessageHandler {
	return &ErrorMessage{err: err}
}
