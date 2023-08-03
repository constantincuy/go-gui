package theme

import (
	"bytes"
	_ "embed"
	"github.com/tdewolff/parse/v2"
	"github.com/tdewolff/parse/v2/css"
	"log"
)

//go:embed default.css
var defaultCSS string

var Engine engine

type engine struct {
	theme       Sheet
	themeLoaded bool
}

func (engine *engine) GetTheme() Sheet {
	if !engine.IsReady() {
		engine.LoadDefaultTheme()
	}

	return engine.theme
}

func (engine *engine) IsReady() bool {
	return engine.themeLoaded
}

func (engine *engine) LoadDefaultTheme() {
	log.Println("[Theme Engine] Using default theme")
	engine.ParseThemeFile(defaultCSS)
}

func (engine *engine) ParseThemeFile(cssString string) {
	engine.theme = NewSheet()
	parser := css.NewParser(parse.NewInput(bytes.NewBufferString(cssString)), false)
	currentKey := ""
	for {
		gt, _, data := parser.Next()
		if gt == css.ErrorGrammar {
			break
		} else if gt == css.BeginRulesetGrammar {
			var value string
			for _, token := range parser.Values() {
				value += string(token.Data)
			}
			currentKey = value
		} else if gt == css.DeclarationGrammar {
			key := string(data)
			d := ""
			for _, val := range parser.Values() {
				d += string(val.Data)
			}
			engine.theme.AddToElement(currentKey, Property{
				Name:  key,
				Value: d,
			})
		}
	}
	engine.themeLoaded = true
}
