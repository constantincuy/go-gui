package common

import (
	"testing"
)

func TestEngine_UpdateRecursive(t *testing.T) {
	engine := ComponentEngine{}
	con1 := newTestableComponent(NewCore())
	con2 := con1.Core().AddChild(newTestableComponent)
	testComp := con2.Core().AddChild(newTestableComponent).(*testableComponent)

	childUpdateCount := 0
	testComp.OnUpdate(func() {
		childUpdateCount++
	})
	engine.UpdateComponentTree(con1)

	if childUpdateCount != 1 {
		t.Error(IntError("Child component should have updated once", 1, childUpdateCount))
	}
}
