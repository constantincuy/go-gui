package open_sans_semi_bold_italic

import _ "embed"
import "golang.org/x/image/font/sfnt"

//go:embed OpenSans-SemiBoldItalic.ttf
var bytes []byte

var cachedFont *sfnt.Font

func Release()     { cachedFont = nil }
func Name() string { return "Open Sans SemiBold" }
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
