package main

import (
	"github.com/smartwalle/imagex"
)

func main() {
	var l1 = imagex.NewImageLayer(100, 100)
	l1.LoadImage("bg.jpg")
	l1.SizeToFit()
	l1.SetPadding(imagex.NewPadding(10, 20, 10, 20))

	var l2 = imagex.NewTextLayer(100, 100)
	l1.AddLayer(l2)
	l2.LoadFont("ZCOOLKuaiLe-Regular.ttf")
	l2.SetFontSize(24)
	l2.SetText("SmartWalle@Copyright")
	l2.SizeToFit()
	l2.SetHorizontalAlignment(imagex.HorizontalAlignmentLeft)
	l2.SetVerticalAlignment(imagex.VerticalAlignmentBottom)

	var l3 = imagex.NewImageLayer(0, 0)
	l1.AddLayer(l3)
	l3.LoadImage("walle.jpg")
	l3.SizeToFit()
	l3.SetPoint(0, 0)
	l3.SetHorizontalAlignment(imagex.HorizontalAlignmentLeft)
	l3.SetVerticalAlignment(imagex.VerticalAlignmentDefault)

	//var l1 = imagex.NewBaseLayer(200, 200)
	//l1.SetBackgroundColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	//
	//var l2 = imagex.NewBaseLayer(20, 20)
	//l2.SetBackgroundColor(color.RGBA{0x00, 0xff, 0x00, 0xff})
	//l1.AddLayer(l2)
	//
	//var l3 = imagex.NewImageLayer(30, 30)
	//l3.SetBackgroundColor(color.RGBA{0x00, 0x00, 0xff, 0xff})
	//l3.SetPoint(20, 0)
	//l1.AddLayer(l3)
	//
	//var l4 = imagex.NewImageLayer(30, 30)
	////l4.SetBackgroundColor(color.RGBA{0xff, 0xff, 0x00, 0xff})
	//l4.SetPoint(0, 30)
	//l4.LoadImage("walle.jpg")
	//l1.AddLayer(l4)
	//
	//var l5 = imagex.NewTextLayer(50, 50)
	//l5.SetBackgroundColor(color.RGBA{0xff, 0x00, 0xff, 0xff})
	//l5.SetPoint(10, 70)
	//l5.SetText("我")
	//l5.LoadFont("ZCOOLKuaiLe-Regular.ttf")
	//l5.SetFontSize(34)
	//l5.SetHorizontalAlignment(imagex.HorizontalAlignmentCenter)
	//l5.SetVerticalAlignment(imagex.VerticalAlignmentMiddle)
	//l5.LoadBackgroundImage("circle.png")
	//l1.AddLayer(l5)
	//
	//l5 = imagex.NewTextLayer(50, 50)
	//l5.SetBackgroundColor(color.RGBA{0xff, 0x00, 0xff, 0xff})
	//l5.SetPoint(60, 70)
	//l5.SetText("的")
	//l5.LoadFont("ZCOOLKuaiLe-Regular.ttf")
	//l5.SetFontSize(34)
	//l5.SetHorizontalAlignment(imagex.HorizontalAlignmentCenter)
	//l5.SetVerticalAlignment(imagex.VerticalAlignmentMiddle)
	//l5.LoadBackgroundImage("circle.png")
	//l5.SetBackgroundColor(nil)
	//l1.AddLayer(l5)

	imagex.WritePNG(l1, "image_out.png")
	//fmt.Println(l5.Rect())
}
