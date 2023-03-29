package nimage

// HorizontalAlignment 水平对齐
type HorizontalAlignment int

// VerticalAlignment 垂直对齐
type VerticalAlignment int

const (
	HorizontalAlignmentDefault HorizontalAlignment = iota
	HorizontalAlignmentLeft
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
)

const (
	VerticalAlignmentDefault VerticalAlignment = iota
	VerticalAlignmentTop
	VerticalAlignmentMiddle
	VerticalAlignmentBottom
)
