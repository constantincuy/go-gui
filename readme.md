# Go-Gui

Go-Gui is a simple GUI lib based on [ebitengine](https://ebitengine.org/).
This project is not meant for production use and is solely a learning project to understand
common challenges in low level UI rendering.

![Go Gui Example](assets/go-gui-example.png)

(See [Counter Example](example/counter))


## Getting Started
For a fast introduction take a look at the [Getting Started](docs/getting-started.md) guide.

## Examples
- [Counter Example](example/counter)
- [Dynamic Content Example](example/dynamic-content)

## Features
- Efficient rendering
- Component based architecture
- Data Binding
- Event System (Click, Hover)
- Customizable render pipelines
- Theming support (pseudo CSS)

### Rendering
The default render pipeline uses a cached image in ram this image is not rerendered on every framed
instead it is only rerendered if a component of the current view proposes a FrameRedraw by marking itself as dirty.
As soon as something on the screen starts moving, resizing etc. frames will be rerendered as needed.

### Component based architecture
Go-Gui uses components to build up the rendered view. 
There are two major component types "native components" and "managed components".

#### Native components
Native components register to the native render call of the render pipeline to 
directly draw on the screen. These are mostly low level components like a box, text
or any other geometric from. Native components mark interactions that dirty the current state 
to let the render pipeline know when a rerender should occur.

#### Virtual Components
Virtual Components don't draw to the screen instead they compose different virtual or native components
to a new reusable block. A button is an example for this its made out of a Rect and Text native component.
The library manages when a component should be rerendered based on their children native components dirty flag or if the layout of the
virtual component changes.

Counter button example (virtual component):
```go
package components

import (
	"fmt"
	"github.com/constantincuy/go-gui/ui/component"
	"github.com/constantincuy/go-gui/ui/event"
)


type Counter struct {
	core         component.Core
	counterState component.State[int]
	button       *component.Button
}

func (c *Counter) Core() *component.Core {
	return &c.core
}

func (c *Counter) Mount() {
	c.button = c.Core().AddChild(component.NewButton).(*component.Button)
	c.counterState = component.NewState(0)
	c.counterState.OnChange(c.setCurrentCount)

	c.button.OnClick(func() {
		c.counterState.SetState(c.counterState.Get() + 1)
	})
}

func (c *Counter) setCurrentCount(count int) {
	c.button.SetText(fmt.Sprintf("Clicked %d times", count))
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

### Data Binding
With state objects you can easily keep your view up to date with your apps internal state.
```go
    // Add a button which should represent the state
	button := c.Core().AddChild(component.NewButton).(*component.Button)
	// Defining a new state with an initial state of 0 
	counterState := component.NewState(0)
	
	// Whenever the state is updated set the text of the button to the current count
	counterState.OnChange(func(count int) {
		button.SetText(fmt.Sprintf("Clicked %d times", count))
	})

	// When the button is clicked increase the counterState
	button.OnClick(func() {
		c.counterState.SetState(c.counterState.Get() + 1)
	})
```


### Event System
Go-Gui supports a rudimentary event system for now (Only mouse events click/hover).
Every component can listen to relevant events by registering a listener in its `Mount()` function via the components 
core and update its state accordingly.

### Customizable render pipelines
It's possible to completely customize the render pipeline for your own use-case. Go-Gui includes 2 render pipelines
the default pipeline responsible for rendering the ui components and a additional one that is used for debugging (Draws debug information on screen).
You can add pipelines on your app instance using the `AddPipeline()` method.

### Theming support
Go-Gui supports theming via a pseudo css format this is limited to mostly colors making it easier to apply
you color scheme ui wide or offer custom theme support for your end users.

## Attribution
This project would not be possible without
- [ebitengine](https://ebitengine.org/)
- [etxt](https://github.com/tinne26/etxt)
- [Open Sans](https://fonts.google.com/specimen/Open+Sans) [[License](ui/font/default/OFL.txt)]