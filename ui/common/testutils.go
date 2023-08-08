package common

import (
	"fmt"
	"image"
)

type testableComponent struct {
	core            Core
	updateListener  func()
	destroyListener func()
	mountListener   func()
}

func (t *testableComponent) Core() *Core {
	return &t.core
}

func (t *testableComponent) OnMount(l func()) {
	t.mountListener = l
}

func (t *testableComponent) Mount() {
	if t.mountListener != nil {
		t.mountListener()
	}
}

func (t *testableComponent) OnUpdate(l func()) {
	t.updateListener = l
}

func (t *testableComponent) Update() {
	if t.updateListener != nil {
		t.updateListener()
	}
}

func (t *testableComponent) OnDestroy(l func()) {
	t.destroyListener = l
}

func (t *testableComponent) Destroy() {
	if t.destroyListener != nil {
		t.destroyListener()
	}
}

func newTestableComponent(core Core) Component {
	return &testableComponent{core: core}
}

func SizeError(message string, want, got Size) string {
	return fmt.Sprintf("%s; want (Width: %d, Height: %d), got (Width: %d, Height: %d)", message, want.Width, want.Height, got.Width, got.Height)
}

func PositionError(message string, want, got image.Point) string {
	return fmt.Sprintf("%s; want (X: %d, Y: %d), got (X: %d, Y: %d)", message, want.X, want.Y, got.X, got.Y)
}

func BoolError(message string, want, got bool) string {
	return fmt.Sprintf("%s; want `%t`, got `%t`", message, want, got)
}

func StringError(message, want, got string) string {
	return fmt.Sprintf("%s; want `%s`, got `%s`", message, want, got)
}

func IntError(message string, want, got int) string {
	return fmt.Sprintf("%s; want `%d`, got `%d`", message, want, got)
}
