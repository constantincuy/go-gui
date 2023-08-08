# Getting Started
To get started lets first create a new go project by executing the following commands in your favourite terminal.
```
mkdir my-app
cd my-app

go mod init github.com/yourname/my-app
```

Add a `main.go` file and add the following code to init a new app.
```go
package main

import (
	"github.com/yourname/my-app/ui/views"
	"github.com/constantincuy/go-gui/ui/app"
	"github.com/constantincuy/go-gui/ui/window"
	"image/color"
)

func main() {
	darkGray := color.RGBA{
		R: 0x22,
		G: 0x22,
		B: 0x22,
		A: 0xff,
	}
	win := window.NewDefaultWindow(views.NewMainView)
	// Window background color
	win.SetBackground(darkGray)
	myApp := app.NewApp(win)
	myApp.Start()
}

```
Here we initialize a new window object with a new main view, which we will be creating in the next step. After creating the
window we create a new app instance using the created window.

After that create a new folder `ui/views` and a new go file inside the `views` folder called `main-view.go` and add the following code to it

```go
package views

import (
	"github.com/constantincuy/go-gui/ui/component"
)

type MainView struct {
	core common.Core
}

// Every component implementation needs to make it's core public
func (view *MainView) Core() *common.Core {
	return &view.core
}

// Mount is called when a component instance is created and added to the screen
func (view *MainView) Mount() {
	// Center the text vertically and horizontally
	view.Core().SetDisplayType(component.FlexCentered())

	//Add a new text component to the view
	headline := view.Core().AddChild(component.NewText).(*component.Text)
	headline.SetFontSize(30)
	headline.SetText("Hello World!")
}

// Update of the component is called 60 per second
func (view *MainView) Update() {}

// Destroy Gets called when a component instance is removed from the screen allowing for resource clean up
func (view *MainView) Destroy() {}

// NewMainView A convention for components is to create a constructor that accepts a core the core is automatically created and injected by AddChild
func NewMainView(core common.Core) common.Component {
	return &MainView{core: core}
}
```
In Go-Gui everything is a component from a simple button to a full-fledged view. When a component is added to the screen
it's `Mount()` method is called allowing you to add children and set default settings of your component. In our example view we 
tell it to use a centered flex layout which is a shorthand for `component.Flex().Justify(component.FlexCenter).Align(component.FlexCenter)`.
Flex layouts are similar to [CSS Flexbox](https://www.w3schools.com/css/css3_flexbox.asp).

After setting the layout type we now add a new Text component to our view and set its font size and the actual text to display.
Text is a built-in component of Go-Gui.

## What's next?
- [Components](components.md)