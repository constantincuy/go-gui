package font

import (
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
		return fm.textRenderer
	}
	return nil
}

func initManager() manager {
	fontLib := etxt.NewFontLibrary()
	textRenderer := etxt.NewStdRenderer()
	glyphsCache := etxt.NewDefaultCache(10 * 1024 * 1024) // 10MB
	textRenderer.SetCacheHandler(glyphsCache.NewHandler())
	textRenderer.SetAlign(etxt.YCenter, etxt.XCenter)
	textRenderer.SetSizePx(64)
	return manager{
		fontLib:      fontLib,
		textRenderer: textRenderer,
	}
}
