package app

import (
	"github.com/constantincuy/go-gui/ui/component"
	"github.com/constantincuy/go-gui/ui/input"
	"github.com/constantincuy/go-gui/ui/pipeline"
	"github.com/constantincuy/go-gui/ui/theme"
	"github.com/constantincuy/go-gui/ui/window"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"os"
)

type App struct {
	renderPipelines []pipeline.Pipeline
	currentWindow   window.Window
	currentTitle    string
	themePath       string
	inputManager    input.Manager
}

func (app *App) AddPipeline(pipe pipeline.Pipeline) {
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

	v := *app.currentWindow.GetView()
	app.inputManager.Update()
	app.inputManager.ProcessEvents(v)
	component.Engine.UpdateRecursive(v)
	component.Engine.RecalculatePositions(v)
	return nil
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
	app.AddPipeline(pipeline.NewDefaultPipeline())
	app.inputManager = input.NewManager()

	return app
}
