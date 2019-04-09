package image4go

//
//import (
//	"errors"
//	"github.com/golang/freetype"
//	"image"
//	"image/draw"
//	"io/ioutil"
//)
//
//func AddText(dst image.Image, text, fontFile string, dpi, fontSize float64, offsetX, offsetY int) (nImage image.Image, err error) {
//	if dst == nil {
//		return nil, errors.New("dst 参数不能为空")
//	}
//
//	// 处理字体
//	fontBytes, err := ioutil.ReadFile(fontFile)
//	if err != nil {
//		return nil, err
//	}
//	font, err := freetype.ParseFont(fontBytes)
//	if err != nil {
//		return nil, err
//	}
//
//	var nCanvas = image.NewRGBA(dst.Bounds())
//	draw.Draw(nCanvas, dst.Bounds(), dst, image.ZP, draw.Src)
//
//	fontCtx := freetype.NewContext()
//	fontCtx.SetDPI(dpi)
//	fontCtx.SetFont(font)
//	fontCtx.SetFontSize(fontSize)
//	fontCtx.SetDst(nCanvas)
//	fontCtx.SetClip(nCanvas.Bounds())
//	fontCtx.SetSrc(image.Black)
//
//	//字体大小
//	var pt = freetype.Pt(offsetX, offsetY+int(fontCtx.PointToFixed(fontSize)>>6))
//
//	fontCtx.DrawString(text, pt)
//
//	nImage = nCanvas
//	return nImage, err
//}
