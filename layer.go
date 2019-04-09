package image4go

import (
	"bufio"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"reflect"
)

type Layer interface {
	Render() image.Image

	Rect() image.Rectangle
}

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

type Size struct {
	Width  int
	Height int
}

func NewSize(width, height int) Size {
	return Size{Width: width, Height: height}
}

func isNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func WriteToPNG(l Layer, file string) (err error) {
	nFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer nFile.Close()

	b := bufio.NewWriter(nFile)

	if err = png.Encode(nFile, l.Render()); err != nil {
		return err
	}
	return b.Flush()
}

func WriteToJPEG(l Layer, file string, quality int) (err error) {
	nFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer nFile.Close()

	b := bufio.NewWriter(nFile)

	if err = jpeg.Encode(nFile, l.Render(), &jpeg.Options{Quality: quality}); err != nil {
		return err
	}
	return b.Flush()
}
