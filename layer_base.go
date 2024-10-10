package imagex

import (
	"image"
	"image/color"
	"image/draw"
)

type BaseLayer struct {
	point               Point
	size                Size
	padding             Padding
	layers              []Layer
	bgColor             color.Color
	horizontalAlignment HorizontalAlignment
	verticalAlignment   VerticalAlignment
}

func NewBaseLayer(width, height int) *BaseLayer {
	var l = &BaseLayer{}
	l.point = NewPoint(0, 0)
	l.size = NewSize(width, height)
	l.bgColor = color.Transparent
	return l
}

func (l *BaseLayer) AddLayer(layer Layer) {
	if isNil(layer) {
		return
	}
	l.layers = append(l.layers, layer)
}

func (l *BaseLayer) RemoveLayer(layer Layer) {
	if isNil(layer) {
		return
	}

	var index = -1
	for i, ele := range l.layers {
		if ele == layer {
			index = i
		}
	}

	if index > -1 {
		l.layers = append(l.layers[:index], l.layers[index+1:]...)
	}
}

func (l *BaseLayer) SetBackgroundColor(bgColor color.Color) {
	l.bgColor = bgColor
	if l.bgColor == nil {
		l.bgColor = color.Transparent
	}
}

func (l *BaseLayer) BackgroundColor() color.Color {
	return l.bgColor
}

func (l *BaseLayer) Render() image.Image {
	var mRect = image.Rect(0, 0, l.size.Width, l.size.Height)
	var mLayer = image.NewRGBA(mRect)

	// 创建背景层
	if l.bgColor != nil {
		var bgLayer = image.NewUniform(l.bgColor)
		draw.Draw(mLayer, mLayer.Bounds(), bgLayer, image.Point{}, draw.Src)
	}

	// 处理子 layer
	for _, layer := range l.layers {
		var img = layer.Render()
		if img != nil {
			var imgRect = calcRect(mRect, layer.Rect(), l.padding, layer.HorizontalAlignment(), layer.VerticalAlignment())
			draw.Draw(mLayer, imgRect, img, image.Point{}, draw.Over)
		}
	}
	return mLayer
}

func (l *BaseLayer) SetPoint(x, y int) {
	l.point = Point{X: x, Y: y}
}

func (l *BaseLayer) Point() Point {
	return l.point
}

func (l *BaseLayer) SetSize(width, height int) {
	l.size = Size{Width: width, Height: height}
}

func (l *BaseLayer) Size() Size {
	return l.size
}

func (l *BaseLayer) Rect() image.Rectangle {
	var r = image.Rectangle{}
	r.Min.X = l.point.X
	r.Min.Y = l.point.Y
	r.Max.X = l.point.X + l.size.Width
	r.Max.Y = l.point.Y + l.size.Height
	return r
}

func (l *BaseLayer) SetHorizontalAlignment(alignment HorizontalAlignment) {
	l.horizontalAlignment = alignment
}

func (l *BaseLayer) HorizontalAlignment() HorizontalAlignment {
	return l.horizontalAlignment
}

func (l *BaseLayer) SetVerticalAlignment(alignment VerticalAlignment) {
	l.verticalAlignment = alignment
}

func (l *BaseLayer) VerticalAlignment() VerticalAlignment {
	return l.verticalAlignment
}

func (l *BaseLayer) SetPadding(p Padding) {
	l.padding = p
}

func (l *BaseLayer) Padding() Padding {
	return l.padding
}
