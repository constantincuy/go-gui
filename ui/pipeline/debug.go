package pipeline

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gode/ui/component"
	"gode/ui/window"
)

type DebugPipeline struct {
	ancestor Pipeline
}

func (pipe *DebugPipeline) Render(screen *ebiten.Image, window window.Window) {
	size := window.GetSize()
	comps := (*window.GetView()).Core().Children()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("[DEBUG]\nFPS: %.0f\nTPS: %.0f\nComponents:%d", ebiten.ActualFPS(), ebiten.ActualTPS(), countComponents(comps)), size.Width-100, 0)
}

func countComponents(comps []*component.Component) int {
	count := len(comps)

	for _, c := range comps {
		children := (*c).Core().Children()
		count += countComponents(children)
	}

	return count
}

func NewDebugPipeline() Pipeline {
	return &DebugPipeline{}
}
