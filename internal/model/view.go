package model

import (
	"fmt"
)

func (m Model) View() string {
	if m.Err != nil {
		return fmt.Sprintf("Error: %v\n\nPress q to exit", m.Err)
	}

	if m.IsProcessing {
		return "Processing...\n\nPress q to exit"
	}

	s := "Select archive to process\n\n"

	for i, file := range m.Files {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, file.Name())
	}

	return s
}
