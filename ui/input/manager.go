package input

import (
	"github.com/constantincuy/go-gui/ui/component"
	"github.com/constantincuy/go-gui/ui/event"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
)

var mouseOffset = image.Point{X: 8, Y: 8}

type Manager struct {
	mousePosition image.Point
	justPressed   bool
	leftClick     bool
	rightClick    bool
}

func (manager *Manager) Update() {
	manager.justPressed = false
	x, y := ebiten.CursorPosition()
	manager.mousePosition = image.Point{
		X: x,
		Y: y,
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		manager.leftClick = true
		manager.justPressed = true
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		manager.rightClick = true
		manager.justPressed = true
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		manager.leftClick = false
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		manager.rightClick = false
	}
}

func (manager *Manager) ProcessEvents(rootComponent component.Component, offset image.Point) {
	if manager.detectMouseCollision(rootComponent, offset) {
		if manager.justPressed {
			var e event.Event
			if manager.pressing() {
				e = event.MouseClickEvent{
					Position: manager.mousePosition,
					Button:   manager.buttons(),
				}
			}

			rootComponent.Core().Events().Fire(e)
			for _, comp := range rootComponent.Core().Children() {
				manager.ProcessEvents(*comp, rootComponent.Core().Position().Add(offset))
			}
		}
	}
}

func (manager *Manager) buttons() []ebiten.MouseButton {
	res := make([]ebiten.MouseButton, 0)
	if manager.leftClick {
		res = append(res, ebiten.MouseButtonLeft)
	}
	if manager.rightClick {
		res = append(res, ebiten.MouseButtonRight)
	}

	return res
}

func (manager *Manager) pressing() bool {
	return manager.rightClick || manager.leftClick
}

func (manager *Manager) detectMouseCollision(comp component.Component, offset image.Point) bool {
	mouseX := manager.mousePosition.X
	mouseY := manager.mousePosition.Y
	relativePos := comp.Core().Position().Add(offset)
	targetX := relativePos.X
	targetY := relativePos.Y
	targetWidth := comp.Core().GetSize().Width
	targetHeight := comp.Core().GetSize().Height
	return (mouseX+mouseOffset.X > targetX && mouseX < targetX+targetWidth) && (mouseY+mouseOffset.Y > targetY && mouseY < targetY+targetHeight)
}

func NewManager() Manager {
	return Manager{
		mousePosition: image.Point{
			X: 0,
			Y: 0,
		},
	}
}
