package image4go

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"os"
)

// 水平对齐
const (
	TextAlignmentLeft   = 0 // default
	TextAlignmentCenter = 1
	TextAlignmentRight  = 2
)

// 垂直对齐
const (
	TextAlignmentTop    = 0 // default
	TextAlignmentMiddle = 1
	TextAlignmentBottom = 2
)

type TextLayer struct {
	*BaseLayer
	font                  *truetype.Font
	dpi                   float64
	fontSize              float64
	textColor             color.Color
	text                  string
	textAlignment         int
	textVerticalAlignment int
	bgImage               image.Image
}

func NewTextLayer(width, height int) *TextLayer {
	var l = &TextLayer{}
	l.BaseLayer = NewBaseLayer(width, height)
	l.dpi = 72
	l.fontSize = 12
	l.textColor = color.Black
	return l
}

func (this *TextLayer) LoadFont(font string) (err error) {
	fontBytes, err := ioutil.ReadFile(font)
	if err != nil {
		return err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}

	this.font = f
	return err
}

func (this *TextLayer) SetDPI(dpi float64) {
	this.dpi = dpi
}

func (this *TextLayer) DPI() float64 {
	return this.dpi
}

func (this *TextLayer) SetFontSize(size float64) {
	this.fontSize = size
}

func (this *TextLayer) FontSize() float64 {
	return this.fontSize
}

func (this *TextLayer) SetTextColor(c color.Color) {
	this.textColor = c
}

func (this *TextLayer) TextColor() color.Color {
	return this.textColor
}

func (this *TextLayer) SetText(t string) {
	this.text = t
}

func (this *TextLayer) Text() string {
	return this.text
}

func (this *TextLayer) SetTextAlignment(alignment int) {
	this.textAlignment = alignment
}

func (this *TextLayer) TextAlignment() int {
	return this.textAlignment
}

func (this *TextLayer) SetTextVerticalAlignment(alignment int) {
	this.textVerticalAlignment = alignment
}

func (this *TextLayer) TextVerticalAlignment() int {
	return this.textVerticalAlignment
}

func (this *TextLayer) SetBackgroundImage(nImage image.Image) {
	this.bgImage = nImage
}

func (this *TextLayer) BackgroundImage() image.Image {
	return this.bgImage
}

func (this *TextLayer) LoadBackgroundImage(file string) (err error) {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	fImage, _, err := image.Decode(f)
	if err != nil {
		return err
	}

	this.SetBackgroundImage(fImage)
	return nil
}

func (this *TextLayer) Render() image.Image {
	var opt = &truetype.Options{}
	opt.Size = this.fontSize
	opt.DPI = this.dpi

	var fontFace = truetype.NewFace(this.font, opt)

	// 文字位置信息
	var textPoint, textSize = this.textRect(fontFace, this.text)

	var mLayer = image.NewRGBA(image.Rect(0, 0, this.size.Width, this.size.Height))
	// 创建背景层
	if this.bgColor != nil {
		var bgLayer = image.NewUniform(this.bgColor)
		draw.Draw(mLayer, mLayer.Bounds(), bgLayer, image.ZP, draw.Src)
	}
	if this.bgImage != nil {
		draw.Draw(mLayer, mLayer.Bounds(), this.bgImage, image.ZP, draw.Over)
	}

	// 文字颜色
	var src = image.NewUniform(this.textColor)

	var drawer = &font.Drawer{}
	drawer.Face = fontFace
	drawer.Src = src
	drawer.Dst = mLayer

	var textX = 0
	var textY = 0

	switch this.textAlignment {
	case TextAlignmentLeft:
		textX = 0
	case TextAlignmentCenter:
		textX = (this.size.Width - textSize.Width) / 2
	case TextAlignmentRight:
		textX = this.size.Width - textSize.Width
	default:
		textX = 0
	}

	switch this.textVerticalAlignment {
	case TextAlignmentTop:
		textY = 0
	case TextAlignmentMiddle:
		textY = (this.size.Height - textSize.Height) / 2
	case TextAlignmentBottom:
		textY = this.size.Height - textSize.Height
	default:
		textY = 0
	}

	drawer.Dot = fixed.Point26_6{
		X: fixed.I(textX),
		Y: fixed.I(textY + textPoint.Y),
	}

	drawer.DrawString(this.text)

	return mLayer
}

func (this *TextLayer) textRect(face font.Face, text string) (point Point, size Size) {
	var bounds, advance = font.BoundString(face, text)
	var w = advance.Ceil()
	var h = int(bounds.Max.Y.Ceil() - bounds.Min.Y.Ceil())
	return Point{X: 0, Y: bounds.Min.Y.Ceil() * -1}, Size{Width: w, Height: h}
}
