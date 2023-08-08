package component

import "math"

type GridLayout struct {
	Columns int
	Gap     int
}

func (g GridLayout) UseColumns(cols int) GridLayout {
	g.Columns = cols
	return g
}

func (g GridLayout) UseGap(gap int) GridLayout {
	g.Gap = gap
	return g
}

func (g GridLayout) ProcessLayout(comp Component) []*Component {
	children := comp.Core().Children()
	size := comp.Core().GetSize()
	columnWidth := size.Width / g.Columns
	maxRows := math.Ceil(float64(len(children) / g.Columns))
	rowHeights := make([]int, int(maxRows))

	for i, child := range children {
		currentRow := i / g.Columns
		childSize := (*child).Core().GetSize()
		if childSize.Height > rowHeights[currentRow] {
			rowHeights[currentRow] = childSize.Height * currentRow
		}
	}

	for i, child := range children {
		currentRow := int(math.Floor(float64(i / g.Columns)))
		currentCol := i % g.Columns
		var y, x int
		x = (i%g.Columns)*columnWidth + (currentCol * g.Gap)
		y = rowHeights[currentRow] + (currentRow * g.Gap)

		(*child).Core().SetPositionXY(x, y)
	}

	size.Height = sumArray(rowHeights) + (g.Gap * int(maxRows-1))
	comp.Core().SetSize(size)

	return children
}

func sumArray(rowHeights []int) int {
	sum := 0
	for _, height := range rowHeights {
		sum += height
	}

	return sum
}

func NewGridLayout() LayoutOptions {
	return GridLayout{}
}
