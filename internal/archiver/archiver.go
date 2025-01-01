package archiver

import (
	"archive/zip"
	"strings"
)

type Archiver struct {
	CountImages int
}

func (a *Archiver) ReadImageFiles(filename string) chan *zip.File {
	imagesChan := make(chan *zip.File)

	go func() {
		archive, err := zip.OpenReader(filename)
		if err != nil {
			close(imagesChan)
			return
		}
		defer archive.Close()

		for _, file := range archive.File {
			if isImage(file) {
				a.CountImages++
				imagesChan <- file
			}
		}
		close(imagesChan)
	}()

	return imagesChan
}

func isImage(file *zip.File) bool {
	extensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".webp": true,
		".tiff": true,
		".svg":  true,
	}

	for ext := range extensions {
		if strings.HasSuffix(strings.ToLower(file.Name), ext) {
			return true
		}
	}
	return false
}
