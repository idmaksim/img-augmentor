package model

import (
	"fmt"
	"os"
	"strings"
)

func (m Model) View() string {
	if m.Err != nil {
		return m.renderError()
	}

	if m.IsProcessing {
		return m.renderProcessing()
	}

	var s strings.Builder
	s.WriteString(m.renderHeader())
	s.WriteString(m.renderFileList())
	return s.String()
}

func (m Model) renderHeader() string {
	var s strings.Builder
	s.Grow(70)
	s.WriteString("\n┌──────────────────────────────┐\n")
	s.WriteString("│      Image Augmentor         │\n")
	s.WriteString("└──────────────────────────────┘\n\n")
	return s.String()
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

	if m.PageSize <= 0 {
		m.PageSize = 10
	}

	startIdx := m.CurrentPage * m.PageSize
	endIdx := min(startIdx+m.PageSize, len(m.Files))

	for i := startIdx; i < endIdx; i++ {
		file := m.Files[i]
		cursor := m.getCursor(i)
		marker := m.getMarker(file)
		s.WriteString(fmt.Sprintf("%s - %s%s\n", marker, cursor, file.Name()))
	}

	totalPages := (len(m.Files) + m.PageSize - 1) / m.PageSize
	s.WriteString(fmt.Sprintf("\nPage %d/%d", m.CurrentPage+1, totalPages))
	s.WriteString("\n← → to navigate pages")

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
