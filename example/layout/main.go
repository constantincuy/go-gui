package main

import (
	"github.com/constantincuy/go-gui/example/layout/ui/views"
	"github.com/constantincuy/go-gui/ui/app"
	"github.com/constantincuy/go-gui/ui/pipeline"
	"github.com/constantincuy/go-gui/ui/window"
	"image/color"
	"os"
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
	if len(os.Args) > 1 && os.Args[1] == "--debug" {
		myApp.AddPipeline(pipeline.NewDebugPipeline())
	}
	myApp.Start()
}
