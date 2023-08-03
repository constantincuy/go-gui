package font

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"os"
)

var Manager = manager{
	fonts:         make(map[string]*opentype.Font, 0),
	fontFaceCache: make(map[string]font.Face, 0),
}

type manager struct {
	fonts         map[string]*opentype.Font
	fontFaceCache map[string]font.Face
}

func (fm *manager) LoadFontFromPath(name string, path string) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	tt, err := opentype.Parse(raw)
	if err != nil {
		return err
	}

	fm.LoadFont(name, tt)

	return nil
}

func (fm *manager) LoadFont(name string, font *opentype.Font) {
	fm.fonts[name] = font
}

// GetFontFace Generates a font face based on the requested size and line height and caches it.
// If the requested font face was generated before it is loaded from cache instead.
func (fm *manager) GetFontFace(name string, size float64, lineHeight float64) (font.Face, error) {
	key := fmt.Sprintf("%s%.2f%.2f", name, size, lineHeight)

	cached, cachedExists := fm.fontFaceCache[key]

	if cachedExists {
		return cached, nil
	}

	fontType, exists := fm.fonts[name]

	if !exists {
		return nil, errors.New("font could not be found! Check font manager registration")
	}

	face, err := opentype.NewFace(fontType, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull, // Use quantization to save glyph cache images.
	})
	if err != nil {
		return nil, err
	}

	// Adjust the line height.
	ff := text.FaceWithLineHeight(face, lineHeight)
	fm.fontFaceCache[key] = ff
	return ff, nil
}
