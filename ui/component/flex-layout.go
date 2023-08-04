package component

import "github.com/constantincuy/go-gui/ui/common"

type FlexPosition string

const (
	FlexStart  FlexPosition = "flex-start"
	FlexCenter FlexPosition = "center"
	FlexEnd    FlexPosition = "flex-end"
)

type FlexDirection string

const (
	FlexRow    FlexDirection = "row"
	FlexColumn FlexDirection = "column"
)

// FlexLayout This layout can position based on horizontal and vertical layout
type FlexLayout struct {
	JustifyContent FlexPosition
	AlignItems     FlexPosition
	Direction      FlexDirection
	Gap            int
}

func (l FlexLayout) UseDirection(dir FlexDirection) FlexLayout {
	l.Direction = dir
	return l
}

func (l FlexLayout) UseGap(gap int) FlexLayout {
	l.Gap = gap
	return l
}

func (l FlexLayout) Justify(pos FlexPosition) FlexLayout {
	l.JustifyContent = pos
	return l
}

func (l FlexLayout) Align(pos FlexPosition) FlexLayout {
	l.AlignItems = pos
	return l
}

func (l FlexLayout) ProcessLayout(comp Component) []*Component {
	children := comp.Core().Children()
	for i, child := range children {
		var y, x int
		widthGetter := func(size common.Size) int { return size.Width }
		heightGetter := func(size common.Size) int { return size.Height }
		childPos := (*child).Core().Position()
		if l.Direction == FlexColumn {
			y = l.groupCentricCalculation(i, children, l.AlignItems, comp, heightGetter, childPos.Y)
			x = l.selfCentricCalculation(l.JustifyContent, comp, child, widthGetter, childPos.X)
		} else {
			y = l.selfCentricCalculation(l.AlignItems, comp, child, heightGetter, childPos.Y)
			x = l.groupCentricCalculation(i, children, l.JustifyContent, comp, widthGetter, childPos.X)
		}
		(*child).Core().SetPositionXY(x, y)
	}
	return children
}

func (l FlexLayout) groupCentricCalculation(index int, allInRow []*Component, flexPos FlexPosition, parent Component, getter func(size common.Size) int, defaultValue int) int {
	parentValue := getter(parent.Core().GetSize())
	switch flexPos {
	case FlexStart:
		if index == 0 {
			return 0
		}
		prev := allInRow[index-1]
		prevSize := (*prev).Core().GetSize()
		prevPos := (*prev).Core().Position()
		return prevPos.X + prevSize.Width
	case FlexCenter:
		rowWidth := l.sumBy(allInRow, getter)
		offset := l.sumBy(allInRow[:index], getter)
		startingPoint := (parentValue / 2) - (rowWidth / 2)
		return startingPoint + offset + l.Gap
	case FlexEnd:
		rowWidth := l.sumBy(allInRow, getter)
		offset := l.sumBy(allInRow[:index], getter)
		startingPoint := parentValue - rowWidth
		return startingPoint + offset
	}

	return defaultValue
}

func (l FlexLayout) selfCentricCalculation(flexPos FlexPosition, parent Component, child *Component, getter func(size common.Size) int, defaultValue int) int {
	parentValue := getter(parent.Core().GetSize())
	childValue := getter((*child).Core().GetSize())
	switch flexPos {
	case FlexStart:
		return 0
	case FlexCenter:
		return (parentValue / 2) - (childValue / 2)
	case FlexEnd:
		return parentValue - childValue
	}

	return defaultValue
}

func (l FlexLayout) sumBy(selection []*Component, sumValueGetter func(size common.Size) int) int {
	value := 0
	for _, c := range selection {
		value += sumValueGetter((*c).Core().GetSize())
	}

	return value + (len(selection) * l.Gap)
}

func Flex() FlexLayout {
	return FlexLayout{}
}

func FlexCentered() FlexLayout {
	return Flex().Justify(FlexCenter).Align(FlexCenter)
}
