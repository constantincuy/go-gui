package component

import (
	"github.com/constantincuy/go-gui/ui/anchor"
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/event"
	"github.com/constantincuy/go-gui/ui/theme"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Core struct {
	name        string
	style       map[string]theme.Property
	position    image.Point
	size        common.Size
	anchor      anchor.Anchor
	canvas      *ebiten.Image
	dataDirty   bool
	layoutDirty bool
	visible     bool
	z           int
	children    []*Component
	renderer    func(bounds image.Rectangle, screen *ebiten.Image)
	eventQueue  event.Queue
}

func (core *Core) ApplyStyle(name string) {
	core.dataDirty = true
	t := theme.Engine.GetTheme()
	s := t.Select(name)
	core.style = s
}

func (core *Core) Style() *map[string]theme.Property {
	return &core.style
}

func (core *Core) Events() *event.Queue {
	return &core.eventQueue
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
	core.layoutDirty = true
	core.z = z
}

func (core *Core) GetZ() int {
	return core.z
}

func (core *Core) IsVisible() bool {
	return core.visible
}

func (core *Core) SetVisible(visible bool) {
	core.layoutDirty = true
	core.visible = visible
}

func (core *Core) ForcesFrameRedraw() bool {
	return core.layoutDirty
}

func (core *Core) ResolveFrameRedraw() {
	core.layoutDirty = false
}

func (core *Core) IsDirty() bool {
	return core.dataDirty
}

func (core *Core) SetDirty(dirty bool) {
	core.dataDirty = dirty
}

func (core *Core) ForceFrameRedraw() {
	core.layoutDirty = true
}

func (core *Core) Move(point image.Point) {
	core.layoutDirty = true
	cur := core.Position()
	core.SetPosition(image.Point{
		X: cur.X + point.X,
		Y: cur.Y + point.Y,
	})
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
	core.layoutDirty = true
	core.position = point
}

func (core *Core) SetPositionXY(x int, y int) {
	core.SetPosition(image.Point{
		X: x,
		Y: y,
	})
}

func (core *Core) SetSize(size common.Size) {
	core.layoutDirty = true
	core.size = size
}

func (core *Core) GetSize() common.Size {
	return core.size
}

func (core *Core) SetAnchor(anchor anchor.Anchor) {
	core.layoutDirty = true
	core.anchor = anchor
}

func (core *Core) GetAnchor() anchor.Anchor {
	return core.anchor
}

func (core *Core) CenterIn(c *Core) {
	core.SetPositionXY((c.size.Width/2)-(core.size.Width/2), (c.size.Height/2)-(core.size.Height/2))
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
	size := common.Size{Width: 0, Height: 0}
	return Core{canvas: nil, dataDirty: true, size: size, visible: true, layoutDirty: true}
}
