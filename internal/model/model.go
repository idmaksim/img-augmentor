package model

import (
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/idmaksim/img-augmentor/internal/augmentor"
	"github.com/idmaksim/img-augmentor/internal/keymap"

)

type Model struct {
	Files     []os.DirEntry
	Cursor    int
	Selected  os.DirEntry
	Help      help.Model
	Keys      keymap.Keymap
	Err       error
	augmentor *augmentor.Augmentor
}
