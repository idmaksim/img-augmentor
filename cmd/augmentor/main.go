package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	initPkg "github.com/idmaksim/img-augmentor/internal/init"
)

func main() {
	// aug := augmentor.New("data", 1)

	// if err := aug.ProcessImages("archive.zip"); err != nil {
	// 	fmt.Printf("Error processing images: %v\n", err)
	// 	os.Exit(1)
	// }

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
