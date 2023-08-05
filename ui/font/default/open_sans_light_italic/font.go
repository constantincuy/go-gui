package open_sans_light_italic

import _ "embed"
import "golang.org/x/image/font/sfnt"

//go:embed OpenSans-LightItalic.ttf
var bytes []byte

var cachedFont *sfnt.Font

func Release()     { cachedFont = nil }
func Name() string { return "Open Sans Light" }
func Font() *sfnt.Font {
	if cachedFont != nil {
		return cachedFont
	}

	var err error
	cachedFont, err = sfnt.Parse(bytes)
	if err != nil {
		panic("Can not init default font: " + err.Error())
	}
	return cachedFont
}
