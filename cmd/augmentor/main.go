package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	initPkg "github.com/idmaksim/img-augmentor/internal/init"
)

func main() {
	model, err := initPkg.InitModel()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}

	app := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}
}
