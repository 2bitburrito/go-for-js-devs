package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"sync"
	"time"

	"golang.org/x/image/draw"
)

var (
	inDir  = "../../../shared/img/"
	outDir = "./out/"
)

func resizeImg(path os.DirEntry, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	if path.Name() == ".DS_Store" {
		return
	}
	fullPath := filepath.Join(inDir, path.Name())
	inputImg, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}
	defer inputImg.Close()

	outPath := filepath.Join(outDir, path.Name())

	outImg, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer outImg.Close()

	src, err := jpeg.Decode(inputImg)
	if err != nil {
		fmt.Println("error with: ", path.Name())
		panic(err)
	}
	dest := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/2, src.Bounds().Max.Y/2))

	draw.CatmullRom.Scale(dest, dest.Rect, src, src.Bounds(), draw.Over, nil)
	jpeg.Encode(outImg, dest, nil)
}

type results struct {
	sizes []int64
	m     sync.Mutex
}

func resizeImgWithMutex(path os.DirEntry, wg *sync.WaitGroup, result *results) {
	if wg != nil {
		defer wg.Done()
	}
	if path.Name() == ".DS_Store" {
		return
	}
	fullPath := filepath.Join(inDir, path.Name())
	inputImg, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}
	defer inputImg.Close()

	outPath := filepath.Join(outDir, path.Name())

	outImg, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer outImg.Close()

	src, err := jpeg.Decode(inputImg)
	if err != nil {
		fmt.Println("error with: ", path.Name())
		panic(err)
	}
	dest := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/2, src.Bounds().Max.Y/2))

	draw.CatmullRom.Scale(dest, dest.Rect, src, src.Bounds(), draw.Over, nil)
	jpeg.Encode(outImg, dest, nil)
	outImgFileInfo, err := outImg.Stat()
	if err != nil {
		panic(err)
	}
	result.m.Lock()
	result.sizes = append(result.sizes, outImgFileInfo.Size())
	result.m.Unlock()
}

func resizeImgWithChannel(path os.DirEntry, result chan int64) {
	if path.Name() == ".DS_Store" {
		result <- 0
		return
	}
	fullPath := filepath.Join(inDir, path.Name())
	inputImg, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}
	defer inputImg.Close()

	outPath := filepath.Join(outDir, path.Name())

	outImg, err := os.Create(outPath)
	if err != nil {
		panic(err)
	}
	defer outImg.Close()

	src, err := jpeg.Decode(inputImg)
	if err != nil {
		fmt.Println("error with: ", path.Name())
		panic(err)
	}
	dest := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/2, src.Bounds().Max.Y/2))

	draw.CatmullRom.Scale(dest, dest.Rect, src, src.Bounds(), draw.Over, nil)
	jpeg.Encode(outImg, dest, nil)
	outImgFileInfo, err := outImg.Stat()
	if err != nil {
		panic(err)
	}
	result <- outImgFileInfo.Size()
}

func main() {
	err := os.RemoveAll(outDir)
	if err != nil {
		panic(err)
	}
	os.Mkdir(outDir, 0o755)
	files, err := os.ReadDir(inDir)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	// Sequential
	for _, f := range files {
		resizeImg(f, nil)
	}
	fmt.Println("Time for sequential image resize: ", time.Since(start))

	start = time.Now()

	// Parallel
	wg := sync.WaitGroup{}
	for _, f := range files {
		wg.Add(1)
		go resizeImg(f, &wg)
	}
	wg.Wait()
	fmt.Println("Time for parallel image resize: ", time.Since(start))

	// With Mutex
	result := results{
		sizes: make([]int64, 0, len(files)),
		m:     sync.Mutex{},
	}

	for _, f := range files {
		wg.Add(1)
		go resizeImgWithMutex(f, &wg, &result)
	}
	wg.Wait()

	start = time.Now()

	// With Channel
	ch := make(chan int64)
	for _, f := range files {
		go resizeImgWithChannel(f, ch)
	}
	res := make([]int64, 0, len(files))

	for range len(files) {
		res = append(res, <-ch)
	}
	close(ch)
}
