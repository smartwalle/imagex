package imagex

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"os"
)

type TextLayer struct {
	*BaseLayer
	font      *truetype.Font
	dpi       float64
	fontSize  float64
	textColor color.Color
	text      string
	bgImage   image.Image
}

func NewTextLayer(width, height int) *TextLayer {
	var l = &TextLayer{}
	l.BaseLayer = NewBaseLayer(width, height)
	l.dpi = 72
	l.fontSize = 12
	l.textColor = color.Black
	return l
}

func (l *TextLayer) LoadFont(file string) (err error) {
	fontBytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	nFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}

	l.font = nFont
	return err
}

func (l *TextLayer) SetDPI(dpi float64) {
	l.dpi = dpi
}

func (l *TextLayer) DPI() float64 {
	return l.dpi
}

func (l *TextLayer) SetFontSize(size float64) {
	l.fontSize = size
}

func (l *TextLayer) FontSize() float64 {
	return l.fontSize
}

func (l *TextLayer) SetTextColor(c color.Color) {
	l.textColor = c
}

func (l *TextLayer) TextColor() color.Color {
	return l.textColor
}

func (l *TextLayer) SetText(text string) {
	l.text = text
}

func (l *TextLayer) Text() string {
	return l.text
}

func (l *TextLayer) SetBackgroundImage(img image.Image) {
	l.bgImage = img
}

func (l *TextLayer) BackgroundImage() image.Image {
	return l.bgImage
}

func (l *TextLayer) LoadBackgroundImage(file string) (err error) {
	imgFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return err
	}

	l.SetBackgroundImage(img)
	return nil
}

func (l *TextLayer) Render() image.Image {
	var opt = &truetype.Options{}
	opt.Size = l.fontSize
	opt.DPI = l.dpi

	var fontFace = truetype.NewFace(l.font, opt)

	// 文字位置及尺寸信息
	var textPoint, textSize = l.textRect(fontFace, l.text)

	var mRect = image.Rect(0, 0, l.size.Width, l.size.Height)
	var mLayer = image.NewRGBA(mRect)
	// 创建背景层
	if l.bgColor != nil {
		var bgLayer = image.NewUniform(l.bgColor)
		draw.Draw(mLayer, mLayer.Bounds(), bgLayer, image.Point{}, draw.Src)
	}
	if l.bgImage != nil {
		draw.Draw(mLayer, mLayer.Bounds(), l.bgImage, image.Point{}, draw.Over)
	}

	// 文字颜色
	var uniform = image.NewUniform(l.textColor)

	var drawer = &font.Drawer{}
	drawer.Face = fontFace
	drawer.Src = uniform
	drawer.Dst = mLayer

	var textX = 0
	var textY = 0

	switch l.horizontalAlignment {
	case HorizontalAlignmentDefault:
		textX = 0
	case HorizontalAlignmentLeft:
		textX = 0
	case HorizontalAlignmentCenter:
		textX = (l.size.Width - textSize.Width) / 2
	case HorizontalAlignmentRight:
		textX = l.size.Width - textSize.Width
	default:
		textX = 0
	}

	switch l.verticalAlignment {
	case VerticalAlignmentDefault:
		textY = 0
	case VerticalAlignmentTop:
		textY = 0
	case VerticalAlignmentMiddle:
		textY = (l.size.Height - textSize.Height) / 2
	case VerticalAlignmentBottom:
		textY = l.size.Height - textSize.Height
	default:
		textY = 0
	}

	drawer.Dot = fixed.Point26_6{
		X: fixed.I(textX),
		Y: fixed.I(textY + textPoint.Y),
	}

	drawer.DrawString(l.text)

	return mLayer
}

func (l *TextLayer) textRect(face font.Face, text string) (point Point, size Size) {
	var bounds, advance = font.BoundString(face, text)
	var w = advance.Ceil()
	var h = bounds.Max.Y.Ceil() - bounds.Min.Y.Ceil()
	return Point{X: 0, Y: bounds.Min.Y.Ceil() * -1}, Size{Width: w, Height: h}
}

func (l *TextLayer) SizeToFit() Size {
	var opts = &truetype.Options{}
	opts.Size = l.fontSize
	opts.DPI = l.dpi

	var fontFace = truetype.NewFace(l.font, opts)
	var _, textSize = l.textRect(fontFace, l.text)
	l.SetSize(textSize.Width, textSize.Height)
	return l.size
}
