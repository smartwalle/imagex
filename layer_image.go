package image4go

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

func (this *ImageLayer) SetImage(nImage image.Image) {
	this.image = nImage
}

func (this *ImageLayer) Image() image.Image {
	return this.image
}

func (this *ImageLayer) LoadImage(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	fImage, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	this.SetImage(fImage)
	return nil
}

func (this *ImageLayer) Render() image.Image {
	var mLayer = image.NewRGBA(image.Rect(0, 0, this.size.Width, this.size.Height))

	// 创建背景层
	if this.bgColor != nil {
		var bgLayer = image.NewUniform(this.bgColor)
		draw.Draw(mLayer, mLayer.Bounds(), bgLayer, image.ZP, draw.Src)
	}

	// 绘制图片
	if this.image != nil {
		draw.Draw(mLayer, image.Rect(0, 0, this.size.Width, this.size.Height), this.image, image.ZP, draw.Over)
	}

	// 处理子 layer
	for _, layer := range this.layers {
		var img = layer.Render()
		if img != nil {
			draw.Draw(mLayer, layer.Rect(), img, image.ZP, draw.Over)
		}
	}
	return mLayer
}

func (this *ImageLayer) SizeToFit() Size {
	var s = this.image.Bounds().Size()
	this.SetSize(s.X, s.Y)
	return this.size
}
