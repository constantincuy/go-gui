package main

import (
	"github.com/constantincuy/go-gui/example/dynamic-content/ui/views"
	"github.com/constantincuy/go-gui/ui/app"
	"github.com/constantincuy/go-gui/ui/window"
	"image/color"
)

func main() {
	win := window.NewDefaultWindow(views.NewMainView)
	win.SetBackground(color.RGBA{
		R: 0x22,
		G: 0x22,
		B: 0x22,
		A: 0xff,
	})
	myApp := app.NewApp(win)
	//myApp.AddPipeline(pipeline.NewDebugPipeline())
	myApp.Start()
}
