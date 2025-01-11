package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/idmaksim/img-augmentor/internal/augmentor"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			m.Cursor = max(m.Cursor-1, 0)
		case "down":
			m.Cursor = min(m.Cursor+1, len(m.Files)-1)
		case "left":
			m.CurrentPage = max(m.CurrentPage-1, 0)
			m.Cursor = m.CurrentPage * m.PageSize
		case "right":
			maxPage := (len(m.Files)+m.PageSize-1)/m.PageSize - 1
			m.CurrentPage = min(m.CurrentPage+1, maxPage)
			m.Cursor = m.CurrentPage * m.PageSize
		}
	}
	return m, nil
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
