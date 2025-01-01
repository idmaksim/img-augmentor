package main

import (
	"fmt"

	"github.com/idmaksim/img-augmentor/internal/archiver"
)

func main() {
	a := archiver.Archiver{}

	for file := range a.ReadImageFiles("archive.zip") {
		fmt.Println(file.Name)
	}
	fmt.Println(a.CountImages)
	// src, err := imaging.Open("input.jpg")
	// if err != nil {
	// 	panic(err)
	// }

	// dst := imaging.Resize(src, 10000, 10000, imaging.Lanczos)
	// err = imaging.Save(dst, "output.jpg")
	// if err != nil {
	// 	panic(err)
	// }

	// Распаковка архива
	// archive, err := zip.OpenReader("archive.zip")
	// if err != nil {
	// 	panic(err)
	// }
	// defer archive.Close()

	// for _, f := range archive.File {
	// 	filePath := filepath.Join("output_directory", f.Name)

	// 	// Выводим информацию о файле
	// 	fmt.Printf("Файл: %s\n", f.Name)
	// 	fmt.Printf("Размер: %d байт\n", f.UncompressedSize64)
	// 	fmt.Printf("Сжатый размер: %d байт\n", f.CompressedSize64)
	// 	fmt.Printf("Метод сжатия: %d\n", f.Method)
	// 	fmt.Printf("Время модификации: %v\n\n", f.Modified)

	// 	if f.FileInfo().IsDir() {
	// 		os.MkdirAll(filePath, os.ModePerm)
	// 		continue
	// 	}

	// 	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
	// 		panic(err)
	// 	}

	// 	dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fileInArchive, err := f.Open()
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	if _, err := io.Copy(dstFile, fileInArchive); err != nil {
	// 		panic(err)
	// 	}

	// 	dstFile.Close()
	// 	fileInArchive.Close()
	// }
}
