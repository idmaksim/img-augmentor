package model

import (
	"fmt"
	"os"
	"strings"
)

func (m Model) View() string {
	var s strings.Builder
	s.WriteString(m.renderHeader())

	if m.Err != nil {
		return m.renderError()
	}

	if m.IsProcessing {
		return m.renderProcessing()
	}

	return m.renderFileList()
}

func (m Model) renderHeader() string {
	return "\n┌──────────────────────────────┐\n│      Image Augmentor         │\n└──────────────────────────────┘\n\n"
}

func (m Model) renderError() string {
	return fmt.Sprintf("❌ Error: %v\n\nPress q to exit", m.Err)
}

func (m Model) renderProcessing() string {
	return "⏳ Processing...\n\nPress q to exit"
}

func (m Model) renderFileList() string {
	var s strings.Builder
	s.WriteString("📁 Select archive to process:\n\n")

	for i, file := range m.Files {
		cursor := m.getCursor(i)
		marker := m.getMarker(file)
		s.WriteString(fmt.Sprintf("%s - %s%s\n", marker, cursor, file.Name()))
	}

	return s.String()
}

func (m Model) getCursor(index int) string {
	if m.Cursor == index {
		return "▶ "
	}
	return "  "
}

func (m Model) getMarker(file os.DirEntry) string {
	if strings.HasSuffix(file.Name(), ".zip") {
		return "✅"
	}
	return "⛔️"
}
