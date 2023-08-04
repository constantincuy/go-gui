package pipeline

import (
	"github.com/constantincuy/go-gui/ui/component"
	"github.com/constantincuy/go-gui/ui/window"
	"image"
)

type Scene struct {
	Window window.Window
}

func (scene *Scene) SceneGraph() ([]ComponentRef, bool) {
	initialPos := image.Point{
		X: 0,
		Y: 0,
	}
	return scene.buildSceneGraph(*scene.Window.GetView(), 0, initialPos)
}

func (scene *Scene) buildSceneGraph(root component.Component, currentZLevel int, relativePosition image.Point) ([]ComponentRef, bool) {
	rootComponents := root.Core().Children()
	sceneGraph := make([]ComponentRef, 0)
	forceRedraw := false

	for _, child := range rootComponents {
		childCore := (*child).Core()
		newZLevel := currentZLevel + childCore.GetZ()
		pos := childCore.Position()
		newRelativePos := relativePosition.Add(pos)
		if childCore.IsVisible() && scene.inScreen(*child) {
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
		childGraph, childForcesRedraw := scene.buildSceneGraph(*child, newZLevel, newRelativePos)
		sceneGraph = append(sceneGraph, childGraph...)
		forceRedraw = forceRedraw || childForcesRedraw
	}

	return sceneGraph, forceRedraw
}

func (scene *Scene) inScreen(comp component.Component) bool {
	winSize := scene.Window.GetSize()
	pos := comp.Core().Position()
	return pos.X <= winSize.Width && pos.Y <= winSize.Height
}

func NewScene(win window.Window) Scene {
	return Scene{Window: win}
}
