package open_sans_italic

import _ "embed"
import "golang.org/x/image/font/sfnt"

//go:embed OpenSans-Italic.ttf
var bytes []byte

var cachedFont *sfnt.Font

func Release()     { cachedFont = nil }
func Name() string { return "Open Sans" }
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
