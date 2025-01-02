package model

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/idmaksim/img-augmentor/internal/keymap"
)

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
