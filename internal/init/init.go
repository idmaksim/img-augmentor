package init

import (
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/idmaksim/img-augmentor/internal/keymap"
	"github.com/idmaksim/img-augmentor/internal/model"
)

func InitModel() (model.Model, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return model.Model{}, err
	}

	return model.Model{
		Files:    files,
		Selected: nil,
		Keys:     &keymap.Keys,
		Help:     help.New(),
		Err:      nil,
		Cursor:   0,
	}, nil
}
