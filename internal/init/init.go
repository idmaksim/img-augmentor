package init

import (
	"os"

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
		Err:      nil,
		Cursor:   0,
	}, nil
}
