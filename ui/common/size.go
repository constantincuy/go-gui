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
