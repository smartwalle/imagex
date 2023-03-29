package nimage

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

	SetHorizontalAlignment(alignment HorizontalAlignment)

	HorizontalAlignment() HorizontalAlignment

	SetVerticalAlignment(alignment VerticalAlignment)

	VerticalAlignment() VerticalAlignment
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

type Padding struct {
	Left   int
	Right  int
	Top    int
	Bottom int
}

func NewPadding(left, right, top, bottom int) Padding {
	return Padding{Left: left, Right: right, Top: top, Bottom: bottom}
}

func isNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func calcRect(pRect, sRect image.Rectangle, padding Padding, horizontalAlignment HorizontalAlignment, verticalAlignment VerticalAlignment) (rect image.Rectangle) {
	// 处理 padding
	var left = padding.Left
	var right = padding.Right
	var top = padding.Top
	var bottom = padding.Bottom
	// 除去 padding 之后的大小
	pRect = image.Rect(0, 0, pRect.Max.X-right-left, pRect.Max.Y-top-bottom)

	var pWidth = pRect.Max.X - pRect.Min.X
	var sWidth = sRect.Max.X - sRect.Min.X

	switch horizontalAlignment {
	case HorizontalAlignmentDefault:
		rect.Min.X = sRect.Min.X
		rect.Max.X = sRect.Max.X
	case HorizontalAlignmentLeft:
		rect.Min.X = pRect.Min.X
		rect.Max.X = sWidth
	case HorizontalAlignmentCenter:
		var w = pWidth - sWidth
		rect.Min.X = w / 2
		rect.Max.X = rect.Min.X + sWidth
	case HorizontalAlignmentRight:
		rect.Min.X = pWidth - sWidth
		rect.Max.X = rect.Min.X + sWidth
	default:
		rect.Min.X = sRect.Min.X
		rect.Max.X = sRect.Max.X
	}

	var pHeight = pRect.Max.Y - pRect.Min.Y
	var sHeight = sRect.Max.Y - sRect.Min.Y

	switch verticalAlignment {
	case VerticalAlignmentDefault:
		rect.Min.Y = sRect.Min.Y
		rect.Max.Y = sRect.Max.Y
	case VerticalAlignmentTop:
		rect.Min.Y = pRect.Min.Y
		rect.Max.Y = sHeight
	case VerticalAlignmentMiddle:
		var h = pHeight - sHeight
		rect.Min.Y = h / 2
		rect.Max.Y = rect.Min.Y + sHeight
	case VerticalAlignmentBottom:
		rect.Min.Y = pHeight - sHeight
		rect.Max.Y = rect.Min.Y + sHeight
	default:
		rect.Min.Y = sRect.Min.Y
		rect.Max.Y = sRect.Max.Y
	}

	// 修正 padding 区域
	rect.Min.X += left
	rect.Max.X += left
	rect.Min.Y += top
	rect.Max.Y += top

	return rect
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
