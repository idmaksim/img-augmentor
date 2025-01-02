package model

import (
	"fmt"
)

func (m Model) View() string {
	if m.Err != nil {
		return fmt.Sprintf("Ошибка: %v\n\nНажмите q для выхода", m.Err)
	}

	if m.IsProcessing {
		return "В обработке...\n\nНажмите q для выхода"
	}

	s := "Select archive to process\n\n"

	for i, file := range m.Files {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, file)
	}

	return s
}
