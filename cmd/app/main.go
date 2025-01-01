package main

import (
	"github.com/disintegration/imaging"
)

func main() {
	src, err := imaging.Open("input.jpg")
	if err != nil {
		panic(err)
	}

	dst := imaging.Resize(src, 10000, 10000, imaging.Lanczos)
	err = imaging.Save(dst, "output.jpg")
	if err != nil {
		panic(err)
	}
}
