package component

import (
	"github.com/constantincuy/go-gui/ui/testutils"
	"testing"
)

type testableComponent struct {
	core            Core
	updateListener  func()
	destroyListener func()
	mountListener   func()
}

func (t *testableComponent) Core() *Core {
	return &t.core
}

func (t *testableComponent) OnMount(l func()) {
	t.mountListener = l
}

func (t *testableComponent) Mount() {
	if t.mountListener != nil {
		t.mountListener()
	}
}

func (t *testableComponent) OnUpdate(l func()) {
	t.updateListener = l
}

func (t *testableComponent) Update() {
	if t.updateListener != nil {
		t.updateListener()
	}
}

func (t *testableComponent) OnDestroy(l func()) {
	t.destroyListener = l
}

func (t *testableComponent) Destroy() {
	if t.destroyListener != nil {
		t.destroyListener()
	}
}

func newTestableComponent(core Core) Component {
	return &testableComponent{core: core}
}

func TestEngine_UpdateRecursive(t *testing.T) {
	con1 := NewGridContainer(1, 0)(NewCore()).(*Container)
	con2 := con1.Add(NewGridContainer(1, 0)).(*Container)
	testComp := con2.Add(newTestableComponent).(*testableComponent)

	childUpdateCount := 0
	testComp.OnUpdate(func() {
		childUpdateCount++
	})
	Engine.UpdateRecursive(con1)

	if childUpdateCount != 1 {
		t.Error(testutils.IntError("Child component should have updated once", 1, childUpdateCount))
	}
}
