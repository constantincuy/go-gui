package common

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"reflect"
	"testing"
)

func createRenderedCore() Core {
	// Simulate initial render
	core := NewCore()
	core.ResolveFrameRedraw()

	return core
}

// TestCore_CreatingNewCoreCreatesNewFrame testing if new components will be rendered immediately
func TestCore_CreatingNewCoreCreatesNewFrame(t *testing.T) {
	core := NewCore()
	if !core.CausesFrameRedraw() {
		t.Error(BoolError("Core should cause frame redraw on initialization", true, core.CausesFrameRedraw()))
	}
}

// TestCore_CoreCausesFrameRedraw test if core can force a frame redraw
func TestCore_CoreCausesFrameRedraw(t *testing.T) {
	core := createRenderedCore()

	core.ForceFrameRedraw()
	if !core.CausesFrameRedraw() {
		t.Error(BoolError("Core should cause frame redraw", true, core.CausesFrameRedraw()))
	}
}

// TestCore_VisibleDefault Test if new components will be shown by default
func TestCore_VisibleDefault(t *testing.T) {
	core := NewCore()
	if !core.visible {
		t.Error(BoolError("Core should be visible", true, core.visible))
	}
}

// TestCore_SetVisible testing if setting the visibility works
func TestCore_SetVisible(t *testing.T) {
	core := NewCore()
	core.SetVisible(false)
	if core.visible {
		t.Error(BoolError("Core should be invisible", false, core.visible))
	}
}

// TestCore_VisibleDoesNotCauseAccidentalRedraw tests if setting the visibility does not cause a redraw when the same value as the current value is set
func TestCore_VisibleDoesNotCauseAccidentalRedraw(t *testing.T) {
	core := createRenderedCore()

	core.SetVisible(true)
	if core.CausesFrameRedraw() {
		t.Error(BoolError("Core should not cause frame redraw on same `visible` value", false, core.CausesFrameRedraw()))
	}
}

// TestCore_DefaultPosition testing if default position is 0,0
func TestCore_DefaultPosition(t *testing.T) {
	core := NewCore()
	expectedPos := image.Point{
		X: 0,
		Y: 0,
	}
	if core.Position() != expectedPos {
		t.Error(PositionError("Core is at the wrong default position", expectedPos, core.Position()))
	}
}

// TestCore_SetPosition testing if setting the position works
func TestCore_SetPosition(t *testing.T) {
	core := NewCore()
	newPos := image.Point{
		X: 10,
		Y: 15,
	}
	core.SetPosition(newPos)
	if core.Position() != newPos {
		t.Error(PositionError("Core is at the wrong position", newPos, core.Position()))
	}
}

// TestCore_PositionDoesNotCauseAccidentalRedraw tests if setting the position does not cause a redraw when the same value as the current value is set
func TestCore_PositionDoesNotCauseAccidentalRedraw(t *testing.T) {
	core := createRenderedCore()

	newPos := image.Point{
		X: 0,
		Y: 0,
	}
	core.SetPosition(newPos)
	if core.CausesFrameRedraw() {
		t.Error(BoolError("Core should not cause frame redraw on same position value", false, core.CausesFrameRedraw()))
	}
}

// TestCore_DefaultDisplayType testing if default display type is block layout
func TestCore_DefaultDisplayType(t *testing.T) {
	core := NewCore()
	expectedType := NewBlockLayout()
	if core.DisplayType() != expectedType {
		t.Error(StringError("Core has the wrong default display type", reflect.TypeOf(expectedType).String(), reflect.TypeOf(core.DisplayType()).String()))
	}
}

// TestCore_SetDisplayType testing if setting the display type (layout) works
func TestCore_SetDisplayType(t *testing.T) {
	core := NewCore()
	newType := Flex().Justify(FlexCenter)
	core.SetDisplayType(newType)
	if core.DisplayType() != newType {
		t.Error(StringError("Core has the wrong display type", reflect.TypeOf(newType).String(), reflect.TypeOf(core.DisplayType()).String()))
	}
}

// TestCore_DisplayTypeDoesNotCauseAccidentalRedraw tests if setting the display type does not cause a redraw when the same value as the current value is set
func TestCore_DisplayTypeDoesNotCauseAccidentalRedraw(t *testing.T) {
	core := createRenderedCore()

	newType := NewBlockLayout()
	core.SetDisplayType(newType)
	if core.CausesFrameRedraw() {
		t.Error(BoolError("Core should not cause frame redraw on same display type value", false, core.CausesFrameRedraw()))
	}
}

// TestCore_DefaultZ testing if default Z i 0
func TestCore_DefaultZ(t *testing.T) {
	core := NewCore()
	expectedZ := 0
	if core.GetZ() != expectedZ {
		t.Error(IntError("Core has the wrong default Z layer", expectedZ, core.GetZ()))
	}
}

// TestCore_SetZ testing if setting the z layer works
func TestCore_SetZ(t *testing.T) {
	core := NewCore()
	newZ := 10
	core.SetZ(newZ)
	if core.GetZ() != newZ {
		t.Error(IntError("Core has the wrong Z layer", newZ, core.GetZ()))
	}
}

// TestCore_ZLayerDoesNotCauseAccidentalRedraw tests if setting the Z layer does not cause a redraw when the same value as the current value is set
func TestCore_ZLayerDoesNotCauseAccidentalRedraw(t *testing.T) {
	core := createRenderedCore()

	newZ := 0
	core.SetZ(newZ)
	if core.CausesFrameRedraw() {
		t.Error(BoolError("Core should not cause frame redraw on same Z layer value", false, core.CausesFrameRedraw()))
	}
}

// TestCore_DefaultSize testing if default size is 0, 0
func TestCore_DefaultSize(t *testing.T) {
	core := NewCore()
	expectedSize := Size{
		Width:  0,
		Height: 0,
	}
	if core.GetSize() != expectedSize {
		t.Error(SizeError("Core has the wrong default size", expectedSize, core.GetSize()))
	}
}

// TestCore_SetSize testing if setting the size works
func TestCore_SetSize(t *testing.T) {
	core := NewCore()
	newSize := Size{
		Width:  100,
		Height: 100,
	}
	core.SetSize(newSize)
	if core.GetSize() != newSize {
		t.Error(SizeError("Core has the wrong size", newSize, core.GetSize()))
	}
}

// TestCore_SizeDoesNotCauseAccidentalRedraw tests if setting the size does not cause a redraw when the same value as the current value is set
func TestCore_SizeDoesNotCauseAccidentalRedraw(t *testing.T) {
	core := createRenderedCore()

	newSize := Size{
		Width:  0,
		Height: 0,
	}
	core.SetSize(newSize)
	if core.CausesFrameRedraw() {
		t.Error(BoolError("Core should not cause frame redraw on same size value", false, core.CausesFrameRedraw()))
	}
}

// TestCore_OnRender
func TestCore_OnRender(t *testing.T) {
	renderCount := 0
	core := NewCore()
	core.OnRender(func(bounds image.Rectangle, screen *ebiten.Image) {
		renderCount++
	})
	bounds := image.Rectangle{
		Min: image.Point{},
		Max: image.Point{},
	}
	core.Render(bounds, nil)

	if renderCount != 1 {
		t.Error(IntError("Core did not render", 1, renderCount))
	}
}
