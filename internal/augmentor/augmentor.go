package augmentor

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/disintegration/imaging"
	"github.com/idmaksim/img-augmentor/internal/archiver"
	"golang.org/x/exp/rand"
)

type Augmentor struct {
	OutputDir string
	archiver  *archiver.Archiver
	copies    int
}

func New(outputDir string, copies int) *Augmentor {
	return &Augmentor{
		OutputDir: outputDir,
		archiver:  archiver.New(),
		copies:    copies,
	}
}

func (a *Augmentor) ProcessImages(archivePath string) error {
	wg := sync.WaitGroup{}

	if err := os.MkdirAll(a.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	for file := range a.archiver.ReadImageFiles(archivePath) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := a.processImage(file.Name); err != nil {
				fmt.Printf("Error processing file %s: %v\n", file.Name, err)
			}
		}()
	}
	wg.Wait()

	return nil
}

func (a *Augmentor) processImage(fileName string) error {
	src, err := a.loadSourceImage(fileName)
	if err != nil {
		return err
	}

	processed := a.applyAugmentations(src)

	if err := a.saveProcessedImages(fileName, processed); err != nil {
		return err
	}

	return nil
}

func (a *Augmentor) loadSourceImage(fileName string) (image.Image, error) {
	filePath := filepath.Join("data", fileName)
	src, err := imaging.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	return src, nil
}

func (a *Augmentor) saveProcessedImages(fileName string, processed []image.Image) error {
	outputPath := filepath.Join(a.OutputDir, fileName)
	outputDirPath := filepath.Dir(outputPath)
	if err := os.MkdirAll(outputDirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", outputDirPath, err)
	}

	for i, img := range processed {
		go func(i int, img image.Image) {
			outputFileName := a.generateOutputFileName(outputPath, i)
			if err := imaging.Save(img, outputFileName); err != nil {
				fmt.Printf("Error saving file %s: %v\n", outputFileName, err)
			}
		}(i, img)
	}
	return nil
}

func (a *Augmentor) generateOutputFileName(outputPath string, index int) string {
	return fmt.Sprintf("%s_%d%s",
		strings.TrimSuffix(outputPath, filepath.Ext(outputPath)),
		index,
		filepath.Ext(outputPath))
}

func (a *Augmentor) applyAugmentations(img image.Image) []image.Image {
	result := make([]image.Image, a.copies)
	for i := 0; i < a.copies; i++ {
		rotated := imaging.Rotate(img, float64(rand.Intn(360)), color.White)
		result[i] = rotated
	}
	return result
}
