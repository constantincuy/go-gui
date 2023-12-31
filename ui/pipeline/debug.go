package pipeline

import (
	"fmt"
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/constantincuy/go-gui/ui/window"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type DebugPipeline struct{}

func (pipe *DebugPipeline) Render(screen *ebiten.Image, win window.Window) {
	size := win.GetSize()
	comps := (*win.GetView()).Core().Children()
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("[DEBUG]\nFPS: %.0f\nTPS: %.0f\nComponents:%d", ebiten.ActualFPS(), ebiten.ActualTPS(), countComponents(comps)), size.Width-100, 0)
}

func countComponents(comps []*common.Component) int {
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
