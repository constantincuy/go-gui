package pipeline

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/window"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type DefaultPipeline struct {
	frameCache FrameCache
}

func (pipe *DefaultPipeline) Render(screen *ebiten.Image, win window.Window) {
	pipe.frameCache.SetCurrentScene(NewScene(win))
	pipe.frameCache.Render(screen)
}

func drawComponent(screen *ebiten.Image, component *common.Component, computedPosition image.Point) {
	core := (*component).Core()
	size := core.GetSize()
	bounds := image.Rectangle{
		Min: computedPosition,
		Max: computedPosition.Add(size.ToPoint()),
	}
	core.Render(bounds, screen)
}

func NewDefaultPipeline() Pipeline {
	frameCache := NewFrameCache()
	frameCache.RenderFrame(func(sceneGraph []ComponentRef, screen *ebiten.Image) {
		for _, ref := range sceneGraph {
			c := ref.Component
			drawComponent(screen, c, ref.ComputedPosition)
		}
	})
	return &DefaultPipeline{frameCache: frameCache}
}
