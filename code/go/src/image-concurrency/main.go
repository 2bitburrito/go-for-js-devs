package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/image/draw"
)

var (
	inDir  = "../../../shared/img/"
	outDir = "./out/"
)

func resizeImg(path string) {
	fullPath := filepath.Join(inDir, path)
	inputImg, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}
	defer inputImg.Close()

	outPath := filepath.Join(outDir, path)

	outImg, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer outImg.Close()

	src, err := jpeg.Decode(inputImg)
	if err != nil {
		fmt.Println("error with: ", path)
		panic(err)
	}
	dest := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/2, src.Bounds().Max.Y/2))

	draw.CatmullRom.Scale(dest, dest.Rect, src, src.Bounds(), draw.Over, nil)
	jpeg.Encode(outImg, dest, nil)
}

func main() {
	err := os.RemoveAll(outDir)
	if err != nil {
		panic(err)
	}
	os.Mkdir(outDir, 0755)
	start := time.Now()
	files, err := os.ReadDir(inDir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		if f.Name() == ".DS_Store" {
			continue
		}
		resizeImg(f.Name())
	}
	fmt.Println("Time for sequential: ", time.Since(start))
}
