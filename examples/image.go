package main

import (
	"github.com/smartwalle/nimage"
)

func main() {
	var l1 = nimage.NewImageLayer(100, 100)
	l1.LoadImage("bg.jpg")
	l1.SizeToFit()
	l1.SetPadding(nimage.NewPadding(10, 20, 10, 20))

	var l2 = nimage.NewTextLayer(100, 100)
	l1.AddLayer(l2)
	l2.LoadFont("ZCOOLKuaiLe-Regular.ttf")
	l2.SetFontSize(24)
	l2.SetText("SmartWalle@Copyright")
	l2.SizeToFit()
	l2.SetHorizontalAlignment(nimage.HorizontalAlignmentLeft)
	l2.SetVerticalAlignment(nimage.VerticalAlignmentBottom)

	var l3 = nimage.NewImageLayer(0, 0)
	l1.AddLayer(l3)
	l3.LoadImage("walle.jpg")
	l3.SizeToFit()
	l3.SetPoint(0, 0)
	l3.SetHorizontalAlignment(nimage.HorizontalAlignmentLeft)
	l3.SetVerticalAlignment(nimage.VerticalAlignmentDefault)

	//var l1 = nimage.NewBaseLayer(200, 200)
	//l1.SetBackgroundColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	//
	//var l2 = nimage.NewBaseLayer(20, 20)
	//l2.SetBackgroundColor(color.RGBA{0x00, 0xff, 0x00, 0xff})
	//l1.AddLayer(l2)
	//
	//var l3 = nimage.NewImageLayer(30, 30)
	//l3.SetBackgroundColor(color.RGBA{0x00, 0x00, 0xff, 0xff})
	//l3.SetPoint(20, 0)
	//l1.AddLayer(l3)
	//
	//var l4 = nimage.NewImageLayer(30, 30)
	////l4.SetBackgroundColor(color.RGBA{0xff, 0xff, 0x00, 0xff})
	//l4.SetPoint(0, 30)
	//l4.LoadImage("walle.jpg")
	//l1.AddLayer(l4)
	//
	//var l5 = nimage.NewTextLayer(50, 50)
	//l5.SetBackgroundColor(color.RGBA{0xff, 0x00, 0xff, 0xff})
	//l5.SetPoint(10, 70)
	//l5.SetText("我")
	//l5.LoadFont("ZCOOLKuaiLe-Regular.ttf")
	//l5.SetFontSize(34)
	//l5.SetHorizontalAlignment(nimage.HorizontalAlignmentCenter)
	//l5.SetVerticalAlignment(nimage.VerticalAlignmentMiddle)
	//l5.LoadBackgroundImage("circle.png")
	//l1.AddLayer(l5)
	//
	//l5 = nimage.NewTextLayer(50, 50)
	//l5.SetBackgroundColor(color.RGBA{0xff, 0x00, 0xff, 0xff})
	//l5.SetPoint(60, 70)
	//l5.SetText("的")
	//l5.LoadFont("ZCOOLKuaiLe-Regular.ttf")
	//l5.SetFontSize(34)
	//l5.SetHorizontalAlignment(nimage.HorizontalAlignmentCenter)
	//l5.SetVerticalAlignment(nimage.VerticalAlignmentMiddle)
	//l5.LoadBackgroundImage("circle.png")
	//l5.SetBackgroundColor(nil)
	//l1.AddLayer(l5)

	nimage.WriteToPNG(l1, "image_out.png")
	//fmt.Println(l5.Rect())
}
