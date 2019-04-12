package main

import (
	"github.com/smartwalle/image4go"
)

func main() {
	var l1 = image4go.NewImageLayer(100, 100)
	l1.LoadImage("a.jpg")
	l1.SizeToFit()

	var l2 = image4go.NewTextLayer(100, 100)
	l1.AddLayer(l2)
	l2.LoadFont("ZCOOLKuaiLe-Regular.ttf")
	l2.SetFontSize(24)
	l2.SetText("SmartWalle@Copyright")
	l2.SizeToFit()

	var l3 = image4go.NewImageLayer(0, 0)
	l1.AddLayer(l3)
	l3.LoadImage("walle.jpg")
	l3.SizeToFit()
	l3.SetPoint(200, 100)

	//var l1 = image4go.NewBaseLayer(200, 200)
	//l1.SetBackgroundColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
	//
	//var l2 = image4go.NewBaseLayer(20, 20)
	//l2.SetBackgroundColor(color.RGBA{0x00, 0xff, 0x00, 0xff})
	//l1.AddLayer(l2)
	//
	//var l3 = image4go.NewImageLayer(30, 30)
	//l3.SetBackgroundColor(color.RGBA{0x00, 0x00, 0xff, 0xff})
	//l3.SetPoint(20, 0)
	//l1.AddLayer(l3)
	//
	//var l4 = image4go.NewImageLayer(30, 30)
	////l4.SetBackgroundColor(color.RGBA{0xff, 0xff, 0x00, 0xff})
	//l4.SetPoint(0, 30)
	//l4.LoadImage("walle.jpg")
	//l1.AddLayer(l4)
	//
	//var l5 = image4go.NewTextLayer(50, 50)
	//l5.SetBackgroundColor(color.RGBA{0xff, 0x00, 0xff, 0xff})
	//l5.SetPoint(10, 70)
	//l5.SetText("我")
	//l5.LoadFont("ZCOOLKuaiLe-Regular.ttf")
	//l5.SetFontSize(34)
	//l5.SetTextAlignment(image4go.TextAlignmentCenter)
	//l5.SetTextVerticalAlignment(image4go.TextAlignmentMiddle)
	//l5.LoadBackgroundImage("circle.png")
	//l1.AddLayer(l5)
	//
	//l5 = image4go.NewTextLayer(50, 50)
	//l5.SetBackgroundColor(color.RGBA{0xff, 0x00, 0xff, 0xff})
	//l5.SetPoint(60, 70)
	//l5.SetText("的")
	//l5.LoadFont("ZCOOLKuaiLe-Regular.ttf")
	//l5.SetFontSize(34)
	//l5.SetTextAlignment(image4go.TextAlignmentCenter)
	//l5.SetTextVerticalAlignment(image4go.TextAlignmentMiddle)
	//l5.LoadBackgroundImage("circle.png")
	//l5.SetBackgroundColor(nil)
	//l1.AddLayer(l5)
	//
	image4go.WriteToPNG(l1, "out.png")
	//fmt.Println(l5.Rect())
}
