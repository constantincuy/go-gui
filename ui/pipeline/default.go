package pipeline

import (
	"github.com/constantincuy/go-gui/ui/component"
	"github.com/constantincuy/go-gui/ui/window"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"sort"
)

type DefaultPipeline struct {
	offscreen *ebiten.Image
}

func (pipe *DefaultPipeline) creatNewCachedImage(window window.Window) {
	size := window.GetSize()
	if pipe.offscreen != nil {
		pipe.offscreen.Dispose()
	}

	pipe.offscreen = ebiten.NewImage(size.Width, size.Height)
	pipe.offscreen.Fill(window.GetBackground())
}

func (pipe *DefaultPipeline) getCached(win window.Window, forceEvict bool) (*ebiten.Image, bool) {
	size := win.GetSize()
	forcedRedraw := false

	if forceEvict {
		pipe.creatNewCachedImage(win)
		return pipe.offscreen, forceEvict
	}

	if pipe.offscreen != nil {
		bounds := pipe.offscreen.Bounds()
		if size.Width != bounds.Dx() || size.Height != bounds.Dy() {
			pipe.creatNewCachedImage(win)
			forcedRedraw = true
		}
	}

	if pipe.offscreen == nil {
		pipe.creatNewCachedImage(win)
		forcedRedraw = true
	}

	return pipe.offscreen, forcedRedraw
}

func (pipe *DefaultPipeline) buildSceneGraph(root component.Component, win window.Window, currentZLevel int, relativePosition image.Point) ([]ComponentRef, bool) {
	rootComponents := root.Core().Children()
	sceneGraph := make([]ComponentRef, 0)
	forceRedraw := false

	for _, child := range rootComponents {
		childCore := (*child).Core()
		newZLevel := currentZLevel + childCore.GetZ()
		pos := childCore.Position()
		newRelativePos := relativePosition.Add(pos)
		if childCore.IsVisible() && pipe.inScreen(win, *child) {
			sceneGraph = append(sceneGraph, ComponentRef{
				Component:        child,
				Parent:           &root,
				ComputedZ:        newZLevel,
				ComputedPosition: newRelativePos,
			})
		}

		if childCore.CausesFrameRedraw() {
			forceRedraw = true
			childCore.ResolveFrameRedraw()
		}
		childGraph, childForcesRedraw := pipe.buildSceneGraph(*child, win, newZLevel, newRelativePos)
		sceneGraph = append(sceneGraph, childGraph...)
		forceRedraw = forceRedraw || childForcesRedraw
	}

	return sceneGraph, forceRedraw
}

func (pipe *DefaultPipeline) Render(screen *ebiten.Image, win window.Window) {
	view := *win.GetView()
	pos := image.Point{X: 0, Y: 0}
	sceneGraph, forceRedraw := pipe.buildSceneGraph(view, win, 0, pos)

	cached, forcedRedraw := pipe.getCached(win, forceRedraw)
	sort.Sort(byZ(sceneGraph))

	if forcedRedraw {
		for _, ref := range sceneGraph {
			c := ref.Component
			drawComponent(cached, c, ref.ComputedPosition)
		}
	}
	screen.DrawImage(cached, nil)
}

func (pipe *DefaultPipeline) inScreen(win window.Window, comp component.Component) bool {
	winSize := win.GetSize()
	pos := comp.Core().Position()
	return pos.X <= winSize.Width && pos.Y <= winSize.Height
}

func drawComponent(screen *ebiten.Image, component *component.Component, computedPosition image.Point) {
	core := (*component).Core()
	size := core.GetSize()
	bounds := image.Rectangle{
		Min: computedPosition,
		Max: computedPosition.Add(size.ToPoint()),
	}
	core.Render(bounds, screen)
}

func NewDefaultPipeline() Pipeline {
	return &DefaultPipeline{}
}
