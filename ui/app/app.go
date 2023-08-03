package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gode/ui/component"
	"gode/ui/input"
	pipeline2 "gode/ui/pipeline"
	"gode/ui/theme"
	"gode/ui/window"
	"log"
	"os"
)

type App struct {
	renderPipelines []pipeline2.Pipeline
	currentWindow   window.Window
	currentTitle    string
	themePath       string
	inputManager    input.Manager
}

func (app *App) AddPipeline(pipe pipeline2.Pipeline) {
	app.renderPipelines = append(app.renderPipelines, pipe)
}

func (app *App) SetWindow(win window.Window) {
	app.currentWindow = win
}

func (app *App) SetTheme(themePath string) {
	app.themePath = themePath
	cssData, err := os.ReadFile(themePath)
	if err != nil {
		log.Fatalln("Could not load theme file ", themePath)
		return
	}

	theme.Engine.ParseThemeFile(string(cssData))
}

func (app *App) Update() error {
	if app.currentWindow.GetTitle() != app.currentTitle {
		app.setCurrentTitle(app.currentWindow.GetTitle())
	}

	app.inputManager.Update()

	v := *app.currentWindow.GetView()
	app.inputManager.ProcessEvents(v.Core().Children())
	updateRecursive(v.Core().Children())
	return nil
}

func updateRecursive(comps []*component.Component) {
	for _, comp := range comps {
		(*comp).Update()
		compCore := (*comp).Core()
		for _, a := range compCore.Children() {
			(*a).Update()
		}

		updateRecursive(compCore.Children())
	}
}

func (app *App) Draw(screen *ebiten.Image) {
	for _, pipe := range app.renderPipelines {
		pipe.Render(screen, app.currentWindow)
	}
}

func (app *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	app.currentWindow.Layout(outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}

func (app *App) setCurrentTitle(title string) {
	app.currentTitle = title
	ebiten.SetWindowTitle(app.currentTitle)
}

func (app *App) Start() {
	app.currentWindow.Layout(640, 480)
	ebiten.SetWindowSize(640, 480)
	app.setCurrentTitle(app.currentWindow.GetTitle())
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if !theme.Engine.IsReady() {
		theme.Engine.LoadDefaultTheme()
	}

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}

func NewApp(win window.Window) App {
	app := App{}
	app.SetWindow(win)
	app.AddPipeline(pipeline2.NewDefaultPipeline())
	app.inputManager = input.NewManager()

	return app
}
