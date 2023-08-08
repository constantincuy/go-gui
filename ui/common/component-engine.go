package common

type ComponentEngine struct{}

func (e ComponentEngine) RecalculatePositions(rootComponent Component) {
	children := rootComponent.Core().DisplayType().ProcessLayout(rootComponent)
	for _, child := range children {
		e.RecalculatePositions(*child)
	}
}

func (e ComponentEngine) UpdateComponentTree(rootComponent Component) {
	rootComponent.Update()
	for _, comp := range rootComponent.Core().Children() {
		e.UpdateComponentTree(*comp)
	}
}
