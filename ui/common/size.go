package common

import "image"

type Size struct {
	Width  int
	Height int
}

func (s Size) ToPoint() image.Point {
	return image.Point{
		X: s.Width,
		Y: s.Height,
	}
}

func SizeFromBounds(rect image.Rectangle) Size {
	return Size{
		Width:  rect.Dx(),
		Height: rect.Dy(),
	}
}
