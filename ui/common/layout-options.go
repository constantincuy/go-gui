package common

type LayoutOptions interface {
	ProcessLayout(comp Component) []*Component
}
