package font

import (
	"github.com/constantincuy/go-gui/ui/font/default/open_sans"
	"github.com/tinne26/etxt"
)

var Manager = initManager()

type manager struct {
	fontLib      *etxt.FontLibrary
	textRenderer *etxt.Renderer
}

func (fm *manager) LoadFontFromPath(path string) error {
	_, err := fm.fontLib.ParseFontFrom(path)
	return err
}

func (fm *manager) TextRenderer(fontName string) *etxt.Renderer {
	if fm.fontLib.HasFont(fontName) {
		fm.textRenderer.SetFont(fm.fontLib.GetFont(fontName))
	} else {
		fm.textRenderer.SetFont(open_sans.Font())
	}

	return fm.textRenderer
}

func initManager() manager {
	fontLib := etxt.NewFontLibrary()
	textRenderer := etxt.NewStdRenderer()
	glyphsCache := etxt.NewDefaultCache(10 * 1024 * 1024) // 10MB
	textRenderer.SetCacheHandler(glyphsCache.NewHandler())
	textRenderer.SetVertAlign(etxt.Top)
	return manager{
		fontLib:      fontLib,
		textRenderer: textRenderer,
	}
}
