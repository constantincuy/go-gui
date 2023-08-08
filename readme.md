# Go-Gui

Go-Gui is a user interface library based on [ebitengine](https://ebitengine.org/). It allows you to build cross-platform
applications that look and feel the same across all operating systems.
This project is not meant for production use and is solely a learning project to understand
common challenges in low level UI rendering.

![Go Gui Example](assets/go-gui-example.png)

(See [Counter Example](example/counter))


## Getting Started
For a fast introduction take a look at the [Getting Started](docs/getting-started.md) guide.

## Examples
- [Counter Example](example/counter)
- [Dynamic Content Example](example/dynamic-content)
- [Layout Example](example/layout)

## Features
- Component based architecture
- Data Binding
- Layout system (FlexLayout, GridLayout)
- Event System (Click, Hover)
- Efficient rendering
- Customizable render pipelines
- Theming support (pseudo CSS)


### Component based architecture
Go-Gui lets you build user interfaces out of individual pieces called components. 
Components can be reused across your application. Go-Gui requires components to implement the `common.Component` interface. Each component is required to store a `common.Core`.
The component core offers methods that are needed for common operations in UI apps.

Common needed methods offered by the component core:
- `Events()`  Allows access to the components event queue
- `SetDisplayType(op LayoutOptions)` Used to change the components layout (BlockLayout, FlexLayout)
- `AddChild(factory func(core Core) Component)` Add child component
- `RemoveChild(child *Component)` Remove child component
- `SetZ(z int)` Set Z layer (Render order)
- `SetVisible(visible bool)` Make component visible or invisible
- `Move(point image.Point)` Move component by X, Y offset
- `SetPosition(point image.Point)` Set the absolute position
- `SetSize(size common.Size)` Set the size of the component

Each defined component should offer a factory method that accepts a `common.Core` and returns a new instance of the component.
```go
func NewRect(core Core) Component {
	return &Rect{core: core}
}
```
(Example from built-in Rect component)

The core will be injected by the `AddChild()` method of `common.Core`. (See full definition above)

#### Native components
Native components register to the native render call of the render pipeline to 
directly draw on the screen. These are mostly low level elements like a rect, text
or any other geometric from. Native components can force frame redraws if their properties are changed (e.g. color, border, text value).

The following native components are provided as built-ins:
- Text (`component.Text`, `component.NewText(core common.Core)`)
- Rect (`component.Rect`, `component.NewRect(core common.Core)`)
- Image (`component.Image`, `component.NewImage(core common.Core)`)

#### Virtual Components
Virtual Components don't draw to the screen instead they compose different virtual or native components
to a new reusable piece. A button is an example for this its made out of a Rect and Text native component.

The following virtual components are provided as built-ins:
- Button (`component.Button`, `component.NewButton(core common.Core)`)
- Container (`component.Container`, `component.NewFlexContainer(initial FlexLayout)`, `NewGridContainer(cols int, gap int)`)

Counter button example (virtual component):
```go
package components

import (
	"fmt"
	"github.com/constantincuy/go-gui/ui/component"
	"github.com/constantincuy/go-gui/ui/event"
)


type Counter struct {
	core         common.Core
	counterState component.State[int]
	button       *component.Button
}

func (c *Counter) Core() *common.Core {
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
func NewCounter(core common.Core) common.Component {
	return &Counter{core: core}
}

```

Read more about components in the [docs](docs/components.md).

### Data Binding
With state objects you can easily keep your view up to date with your apps internal state.
```go
// Add a button to display the state
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
Every component can listen to relevant events by registering a listener in its `Mount()` method via the components 
core and update its state accordingly.

Advanced event handling including keyboard events are planed.

### Rendering
The default render pipeline uses a cached image in ram this image is not redrawn on every frame.
Go-Gui refreshes the cached image automatically when it detects that properties of components in the current
view changed. This ensures that we don't waste resources when the app is doing nothing instead Go-Gui just displays a static
image.

### Customizable render pipelines
It's possible to completely customize the render pipeline for your own use-case. Go-Gui includes 2 render pipelines
the default pipeline responsible for rendering the ui components and a additional one that is used for debugging (Draws debug information on the screen).
You can add pipelines on your app instance using the `AddPipeline()` method.

### Theming support
Go-Gui supports theming via a pseudo CSS format this is limited to mostly colors making it easier to apply
you color scheme UI wide or offer custom theme support for your end users.

## Attribution
This project would not be possible without
- [ebitengine](https://ebitengine.org/)
- [etxt](https://github.com/tinne26/etxt)
- [Open Sans](https://fonts.google.com/specimen/Open+Sans) [[License](ui/font/default/OFL.txt)]