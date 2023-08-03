package layout

import (
	"github.com/constantincuy/go-gui/ui/component"
)

type FlexPositions string

const (
	FlexStart  FlexPositions = "flex-start"
	FlexCenter FlexPositions = "center"
	FlexEnd    FlexPositions = "flex-end"
)

// FlexLayout This layout can position based on horizontal and vertical layout
type FlexLayout struct {
	JustifyContent FlexPositions
	AlignItems     FlexPositions
}

func (l FlexLayout) ProcessLayout(comp component.Component) []*component.Component {
	children := comp.Core().Children()
	for i, child := range children {
		y := l.yCalculation(comp, child)
		x := l.xCalculation(i, children, comp, child)
		(*child).Core().SetPositionXY(x, y)
	}
	return children
}

func (l FlexLayout) yCalculation(parent component.Component, child *component.Component) int {
	parentSize := parent.Core().GetSize()
	childPos := (*child).Core().Position()
	childSize := (*child).Core().GetSize()
	switch l.AlignItems {
	case FlexStart:
		return 0
	case FlexCenter:
		return (parentSize.Height / 2) - (childSize.Height / 2)
	case FlexEnd:
		return parentSize.Height - childSize.Height
	}

	return childPos.Y
}

func (l FlexLayout) xCalculation(index int, allInRow []*component.Component, parent component.Component, child *component.Component) int {
	parentSize := parent.Core().GetSize()
	childPos := (*child).Core().Position()
	switch l.AlignItems {
	case FlexStart:
		if index == 0 {
			return 0
		}
		prev := allInRow[index-1]
		prevSize := (*prev).Core().GetSize()
		prevPos := (*prev).Core().Position()
		return prevPos.X + prevSize.Width
	case FlexCenter:
		rowWidth := sumWidth(allInRow)
		offset := sumWidth(allInRow[:index])
		startingPoint := (parentSize.Width / 2) - (rowWidth / 2)
		return startingPoint + offset
	case FlexEnd:
		rowWidth := sumWidth(allInRow)
		offset := sumWidth(allInRow[:index])
		startingPoint := parentSize.Width - rowWidth
		return startingPoint + offset
	}

	return childPos.Y
}

func sumWidth(selection []*component.Component) int {
	width := 0
	for _, c := range selection {
		width += (*c).Core().GetSize().Width
	}

	return width
}
