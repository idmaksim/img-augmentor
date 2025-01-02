package model

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/idmaksim/img-augmentor/internal/augmentor"
	"github.com/idmaksim/img-augmentor/internal/keymap"
)

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

func (k *KeyMessage) Handle(m Model) (Model, tea.Cmd) {
	switch {
	case key.Matches(k.msg, keymap.Keys.Quit):
		return m, tea.Quit
	case key.Matches(k.msg, keymap.Keys.Up):
		return m.moveCursorUp(), nil
	case key.Matches(k.msg, keymap.Keys.Down):
		return m.moveCursorDown(), nil
	case key.Matches(k.msg, keymap.Keys.Enter):
		return m.startProcessing()
	default:
		return m, nil
	}
}

func (p *ProcessFinishedMessage) Handle(m Model) (Model, tea.Cmd) {
	m.IsProcessing, m.Selected = false, nil
	return m, nil
}

func (e *ErrorMessage) Handle(m Model) (Model, tea.Cmd) {
	m.Err, m.IsProcessing = e.err, false
	return m, nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var handler MessageHandler
	switch msg := msg.(type) {
	case tea.KeyMsg:
		handler = NewKeyHandler(msg)
	case ProcessFinishedMessage:
		handler = NewProcessFinishedHandler()
	case error:
		handler = NewErrorHandler(msg)
	default:
		return m, nil
	}
	return handler.Handle(m)
}

func (m Model) moveCursorUp() Model {
	if m.Cursor > 0 {
		m.Cursor--
	}
	return m
}

func (m Model) moveCursorDown() Model {
	if m.Cursor < len(m.Files)-1 {
		m.Cursor++
	}
	return m
}

func (m Model) startProcessing() (Model, tea.Cmd) {
	m.Selected = m.Files[m.Cursor]
	m.IsProcessing = true
	m.Augmentor = augmentor.New("data", 10)

	return m, func() tea.Msg {
		if err := m.Augmentor.ProcessImages(m.Selected.Name()); err != nil {
			return err
		}
		return ProcessFinishedMessage{}
	}
}
