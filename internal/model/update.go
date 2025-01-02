package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/idmaksim/img-augmentor/internal/augmentor"
)

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
