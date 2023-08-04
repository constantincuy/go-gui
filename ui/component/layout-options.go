package component

type LayoutOptions interface {
	ProcessLayout(comp Component) []*Component
}
