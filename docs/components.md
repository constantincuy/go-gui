# Components
Components are the main building block which are used to build up a Go-Gui UI. Components allow you to 
efficiently split up and reuse common UI elements across your application. Go-Gui components are heavily
inspired by modern web frontend frameworks like react or vue to just name a few.


**Example Counter Component**
```go
package components

import (
	"fmt"
	"github.com/constantincuy/go-gui/ui/component"
	"github.com/constantincuy/go-gui/ui/event"
)

type Counter struct {
	core    component.Core
	counter int
	button  *component.Button
}

func (c *Counter) Core() *component.Core {
	return &c.core
}

func (c *Counter) Mount() {
	// Adding a button to our counter component
	c.button = c.Core().AddChild(component.NewButton).(*component.Button)
	c.counter = 0
	c.setCurrentCount()

	// Register click listener to increment the count state
	c.button.OnClick(func(e event.Event) {
		c.counter++
		c.setCurrentCount()
	})
}

func (c *Counter) setCurrentCount() {
	c.button.SetText(fmt.Sprintf("Clicked %d times", c.counter))
}

func (c *Counter) Update() {
	//Handle update logic target is 60 update calls per second
}

func (c *Counter) Destroy() {
	//Clean up of resources if needed
}

// NewCounter Factory for `AddChild` method `c.Core().AddChild(NewCounter)`
func NewCounter(core component.Core) component.Component {
	return &Counter{core: core}
}


```

## Best Practices

### Avoid tight coupling of components
There are a few simple rules to follow when designing components which allow
you to maximize the reusability.

#### Don't manipulate the children of children
You should only add to your current components children and not to their sub children
this causes problems if you want to reuse the components as a child of a new component
now you may have to copy code over from the first component to get the same result on the screen.

If you need a component to add new components as child expose a public method for parent components
to use. An example for this could be a simple grid component you may want to create a method `AddToGrid(func factory(c component.Core) component.Component)`

````go
func (view *MainView) Mount() {
	// Adding a grid (allowed operation)
	grid := view.Core().AddChild(component.NewGrid).(*component.Grid)
	
	// This should be avoided and causes tight coupling and hinders reuse.
	// What if the grid needs to do some processing after a child was added?
	// It won't notice the child being added
	grid.Core().AddChild(component.NewText).(*component.Text)
	
	// Instead grid should expose a method to add child elements
	// In this case the consumer does not have access to the child after adding it to the grid
	// But what if we want to change the text, font or size? See the next example
	grid.Add(component.NewText)

	// If the grid wants to expose the child to a consumer
	// it can return a pointer to it as a result 
	child := grid.Add(component.NewText).(*component.Text)
	child.SetText("Hello World!")
}
````

The same apply to other properties like sizes, colors or text values.

#### Don't listen to children events
Similar to the first point you should not hook to the events of children instead expose a public
method from the child again e.g. `OnClick()`.

### Mental Model
To better understand the mental model of components you should see them as modules or libraries they should expose methods to change configurable parts or logic
to their consumers but should keep their internals private and free from side effects. The core of each component being publicly exposed invites to
manipulate children but should be avoided at all cost due to the things listed above.

The core needs to be exposed to give the render pipeline read access to the components properties
like size, position and draw calls of native components.

Except for the components manipulating their own core it is only manipulated by the layout process of the engine, which is executed after each component was updated it calculates new sizes
and positions based on layout options (e.g. FlexLayout) and auto sizing options.

