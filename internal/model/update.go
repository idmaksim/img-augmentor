package model

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/idmaksim/img-augmentor/internal/keymap"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keymap.Keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, keymap.Keys.Up):
			if m.Cursor > 0 {
				m.Cursor--
			}

		case key.Matches(msg, keymap.Keys.Down):
			if m.Cursor < len(m.Files)-1 {
				m.Cursor++
			}

		case key.Matches(msg, keymap.Keys.Enter):

		}
	}

	return m, nil
}
