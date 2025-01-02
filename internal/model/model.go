package model

import (
	"os"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/idmaksim/img-augmentor/internal/augmentor"
	"github.com/idmaksim/img-augmentor/internal/keymap"
)

type Model struct {
	Files        []os.DirEntry
	Cursor       int
	Selected     os.DirEntry
	Help         help.Model
	Keys         *keymap.Keymap
	Err          error
	Augmentor    *augmentor.Augmentor
	IsProcessing bool
}

func (m Model) MoveCursorUp() Model {
	if m.Cursor > 0 {
		m.Cursor--
	}
	return m
}

func (m Model) MoveCursorDown() Model {
	if m.Cursor < len(m.Files)-1 {
		m.Cursor++
	}
	return m
}

func (m Model) StartProcessing() (Model, tea.Cmd) {
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
