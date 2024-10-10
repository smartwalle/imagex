package imagex

import (
	"image"
	"image/draw"
	"os"
)

type ImageLayer struct {
	*BaseLayer
	image image.Image
}

func NewImageLayer(width, height int) *ImageLayer {
	var l = &ImageLayer{}
	l.BaseLayer = NewBaseLayer(width, height)
	return l
}

func (l *ImageLayer) SetImage(img image.Image) {
	l.image = img
}

func (l *ImageLayer) Image() image.Image {
	return l.image
}

func (l *ImageLayer) LoadImage(file string) (err error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}

	l.SetImage(img)
	return nil
}

func (l *ImageLayer) Render() image.Image {
	var mRect = image.Rect(0, 0, l.size.Width, l.size.Height)
	var mLayer = image.NewRGBA(mRect)

	// 创建背景层
	if l.bgColor != nil {
		var bgLayer = image.NewUniform(l.bgColor)
		draw.Draw(mLayer, mLayer.Bounds(), bgLayer, image.Point{}, draw.Src)
	}

	// 绘制图片
	if l.image != nil {
		draw.Draw(mLayer, mRect, l.image, image.Point{}, draw.Over)
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

func (l *ImageLayer) SizeToFit() Size {
	var s = l.image.Bounds().Size()
	l.SetSize(s.X, s.Y)
	return l.size
}
