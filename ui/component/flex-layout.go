package component

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
	Gap            int
}

func (l FlexLayout) ProcessLayout(comp Component) []*Component {
	children := comp.Core().Children()
	for i, child := range children {
		y := l.yCalculation(comp, child)
		x := l.xCalculation(i, children, comp, child)
		(*child).Core().SetPositionXY(x, y)
	}
	return children
}

func (l FlexLayout) yCalculation(parent Component, child *Component) int {
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

func (l FlexLayout) xCalculation(index int, allInRow []*Component, parent Component, child *Component) int {
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
		rowWidth := l.sumWidth(allInRow)
		offset := l.sumWidth(allInRow[:index])
		startingPoint := (parentSize.Width / 2) - (rowWidth / 2)
		return startingPoint + offset + l.Gap
	case FlexEnd:
		rowWidth := l.sumWidth(allInRow)
		offset := l.sumWidth(allInRow[:index])
		startingPoint := parentSize.Width - rowWidth
		return startingPoint + offset
	}

	return childPos.Y
}

func (l FlexLayout) sumWidth(selection []*Component) int {
	width := 0
	for _, c := range selection {
		width += (*c).Core().GetSize().Width
	}

	return width + (len(selection) * l.Gap)
}

func FlexLayoutCentered() LayoutOptions {
	return FlexLayout{
		JustifyContent: FlexCenter,
		AlignItems:     FlexCenter,
	}
}
