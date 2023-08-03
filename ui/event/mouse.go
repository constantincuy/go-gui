package event

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type HoverState bool

const (
	StartHover HoverState = true
	EndHover              = false
)

type MouseClickEvent struct {
	Position image.Point
	Button   []ebiten.MouseButton
}

type MouseHoverEvent struct {
	Position image.Point
	State    HoverState
}

func (e MouseClickEvent) EventImpl() {}
func (e MouseHoverEvent) EventImpl() {}
