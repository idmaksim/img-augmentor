package archiver

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Archiver struct {
	CountImages int
}

func New() *Archiver {
	return &Archiver{}
}

func (a *Archiver) ReadImageFiles(filename string) chan *zip.File {
	imagesChan := make(chan *zip.File)
	go a.processArchive(filename, imagesChan)
	return imagesChan
}

func (a *Archiver) processArchive(filename string, imagesChan chan *zip.File) {
	archive, err := a.openArchive(filename)
	if err != nil {
		close(imagesChan)
		return
	}
	defer archive.Close()

	a.processFiles(archive.File, imagesChan)
	close(imagesChan)
}

func (a *Archiver) processFiles(files []*zip.File, imagesChan chan *zip.File) {
	for _, file := range files {
		if isImage(file) {
			a.processImageFile(file, imagesChan)
		}
	}
}

func (a *Archiver) processImageFile(file *zip.File, imagesChan chan *zip.File) {
	if err := saveFile(file); err != nil {
		return
	}
	a.CountImages++
	imagesChan <- file
}

func (a *Archiver) openArchive(filename string) (*zip.ReadCloser, error) {
	return zip.OpenReader(filename)
}

func isImage(file *zip.File) bool {
	extensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	for ext := range extensions {
		if strings.HasSuffix(strings.ToLower(file.Name), ext) {
			return true
		}
	}
	return false
}

func saveFile(file *zip.File) error {
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	path := filepath.Join("data", file.Name)
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, rc)
	return err
}
