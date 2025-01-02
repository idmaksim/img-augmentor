package model

import (
	"os"

	"github.com/idmaksim/img-augmentor/internal/augmentor"
)

type Model struct {
	Files        []os.DirEntry
	Cursor       int
	Selected     os.DirEntry
	IsProcessing bool
	Err          error
	Augmentor    *augmentor.Augmentor
}
