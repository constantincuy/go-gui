package pipeline

import (
	"github.com/constantincuy/go-gui/ui/common"
	"image"
)

type ComponentRef struct {
	Parent           *common.Component
	Component        *common.Component
	ComputedZ        int
	ComputedPosition image.Point
}

func (ref ComponentRef) ComputeZ() int {
	ac := (*ref.Parent).Core()
	cc := (*ref.Component).Core()
	return ac.GetZ() + cc.GetZ()
}

type byZ []ComponentRef

func (s byZ) Len() int {
	return len(s)
}
func (s byZ) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byZ) Less(i, j int) bool {
	return s[i].ComputeZ() < s[j].ComputeZ()
}
