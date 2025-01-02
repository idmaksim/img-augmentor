package model

import (
	"fmt"
)

func (m Model) View() string {
	if m.Err != nil {
		return fmt.Sprintf("Ошибка: %v\n\nНажмите q для выхода", m.Err)
	}

	s := "What should you buy?\n\n"

	for i, file := range m.Files {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, file)
	}

	s += m.Help.View(m.Keys)

	return s
}
