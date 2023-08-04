package component

var Engine engine

type engine struct{}

func (e engine) RecalculatePositions(rootComponent Component) {
	children := rootComponent.Core().DisplayType().ProcessLayout(rootComponent)
	for _, child := range children {
		e.RecalculatePositions(*child)
	}
}

func (e engine) UpdateRecursive(rootComponent Component) {
	rootComponent.Update()
	for _, comp := range rootComponent.Core().Children() {
		e.UpdateRecursive(*comp)
	}
}
