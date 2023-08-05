package component

import (
	"fmt"
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/event"
	"github.com/constantincuy/go-gui/ui/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type StyleListener func(property theme.Property)

type Core struct {
	styleName      string
	style          map[string]theme.Property
	position       image.Point
	displayType    LayoutOptions
	size           common.Size
	canvas         *ebiten.Image
	dirty          bool
	visible        bool
	z              int
	children       []*Component
	renderer       func(bounds image.Rectangle, screen *ebiten.Image)
	eventQueue     event.Queue
	styleListeners []StyleListener
}

func (core *Core) ApplyStyle(name string) {
	if name != core.styleName {
		core.ForceFrameRedraw()
		t := theme.Engine.GetTheme()
		s := t.Select(name)
		core.styleName = name
		core.style = make(map[string]theme.Property)
		for _, key := range s {
			core.ApplyProperty(key)
		}
	}
}

func (core *Core) ApplyPixelProperty(name string, px int) {
	core.ApplyPropertyValue(name, fmt.Sprintf("%dpx", px))
}

func (core *Core) ApplyColorProperty(name string, color color.RGBA) {
	core.ApplyPropertyValue(name, fmt.Sprintf("#%02x%02x%02x", color.R, color.G, color.B))
}
func (core *Core) ApplyPropertyValue(name string, value string) {
	core.ApplyProperty(theme.Property{
		Name:  name,
		Value: value,
	})
}

func (core *Core) ApplyProperty(prop theme.Property) {
	core.style[prop.Name] = prop
	for _, l := range core.styleListeners {
		l(prop)
	}
	core.ForceFrameRedraw()
}

func (core *Core) OnStyleChange(listener StyleListener) {
	core.styleListeners = append(core.styleListeners, listener)
}

func (core *Core) Style() *map[string]theme.Property {
	return &core.style
}

func (core *Core) GetPixelProperty(name string) int {
	prop, exists := core.style[name]

	if !exists {
		return 0
	}

	px, _ := prop.AsPX()
	return px
}

func (core *Core) GetColorProperty(name string, fallbackColor color.RGBA) color.RGBA {
	prop, exists := core.style[name]

	if !exists {
		return fallbackColor
	}

	px, err := prop.AsColor()

	if err != nil {
		return fallbackColor
	}

	return px
}

func (core *Core) Events() *event.Queue {
	return &core.eventQueue
}

func (core *Core) SetDisplayType(op LayoutOptions) {
	core.displayType = op
}

func (core *Core) DisplayType() LayoutOptions {
	return core.displayType
}

func (core *Core) Children() []*Component {
	return core.children
}

func (core *Core) AddChild(factory func(core Core) Component) Component {
	newCore := NewCore()
	child := factory(newCore)
	child.Mount()
	core.children = append(core.children, &child)

	return child
}

func (core *Core) SetZ(z int) {
	if core.z != z {
		core.ForceFrameRedraw()
		core.z = z
	}
}

func (core *Core) GetZ() int {
	return core.z
}

func (core *Core) IsVisible() bool {
	return core.visible
}

func (core *Core) SetVisible(visible bool) {
	if core.visible != visible {
		core.ForceFrameRedraw()
		core.visible = visible
	}
}

func (core *Core) CausesFrameRedraw() bool {
	return core.dirty
}

func (core *Core) ResolveFrameRedraw() {
	core.dirty = false
}

func (core *Core) ForceFrameRedraw() {
	core.dirty = true
}

func (core *Core) Move(point image.Point) {
	core.SetPosition(core.Position().Add(point))
}

func (core *Core) MoveXY(x int, y int) {
	core.Move(image.Point{
		X: x,
		Y: y,
	})
}

func (core *Core) Position() image.Point {
	return core.position
}

func (core *Core) SetPosition(point image.Point) {
	if !core.position.Eq(point) {
		core.ForceFrameRedraw()
		core.position = point
	}
}

func (core *Core) SetPositionXY(x int, y int) {
	core.SetPosition(image.Point{
		X: x,
		Y: y,
	})
}

func (core *Core) SetSize(size common.Size) {
	if core.size != size {
		core.ForceFrameRedraw()
		core.size = size
	}
}

func (core *Core) GetSize() common.Size {
	return core.size
}

func (core *Core) OnRender(renderer func(bounds image.Rectangle, screen *ebiten.Image)) {
	core.renderer = renderer
}
func (core *Core) Render(bounds image.Rectangle, screen *ebiten.Image) {
	if core.renderer != nil {
		core.renderer(bounds, screen)
	}
}

func (core *Core) Destroy() {
	core.canvas.Dispose()
}

func NewCore() Core {
	// Always set dirty to true on creation to trigger initial render
	return Core{
		eventQueue:  event.NewEventQueue(),
		canvas:      nil,
		dirty:       true,
		size:        common.Size{Width: 0, Height: 0},
		visible:     true,
		displayType: BlockLayout{},
	}
}
