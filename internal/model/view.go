package model

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	var s strings.Builder
	s.WriteString("\n┌──────────────────────────────┐\n")
	s.WriteString("│      Image Augmentor         │\n")
	s.WriteString("└──────────────────────────────┘\n\n")

	if m.Err != nil {
		s.WriteString(fmt.Sprintf("❌ Error: %v\n\nPress q to exit", m.Err))
		return s.String()
	}

	if m.IsProcessing {
		s.WriteString("⏳ Processing...\n\nPress q to exit")
		return s.String()
	}

	s.WriteString("📁 Select archive to process:\n\n")

	for i, file := range m.Files {
		cursor := "  "
		marker := "⛔️"
		if m.Cursor == i {
			cursor = "▶ "
		}

		if strings.HasSuffix(file.Name(), ".zip") {
			marker = "✅"
		}
		s.WriteString(fmt.Sprintf("%s - %s%s\n", marker, cursor, file.Name()))
	}

	return s.String()
}
