package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gode/ui/component"
	"gode/ui/event"
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

func (manager *Manager) ProcessEvents(comps []*component.Component) {
	for _, comp := range comps {
		if manager.detectMouseCollision(*comp) {
			if manager.justPressed {
				var e event.Event
				if manager.pressing() {
					e = event.MouseClickEvent{
						Position: manager.mousePosition,
						Button:   manager.buttons(),
					}
				}

				(*comp).Core().Events().Fire(e)
			}
		}
	}
}

func (manager *Manager) buttons() []ebiten.MouseButton {
	res := make([]ebiten.MouseButton, 2)
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

func (manager *Manager) detectMouseCollision(comp component.Component) bool {
	mouseX := manager.mousePosition.X
	mouseY := manager.mousePosition.Y
	targetX := comp.Core().Position().X
	targetY := comp.Core().Position().Y
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
