package pipeline

import (
	"github.com/constantincuy/go-gui/ui/common"
	"github.com/hajimehoshi/ebiten/v2"
)

type FrameRenderer func(sceneGraph []ComponentRef, screen *ebiten.Image)

type FrameCache struct {
	cachedImage   *ebiten.Image
	currentScene  Scene
	frameRenderer FrameRenderer
}

func (frameCache *FrameCache) SetCurrentScene(currentScene Scene) {
	frameCache.currentScene = currentScene
}

func (frameCache *FrameCache) creatNewCachedImage() {
	win := frameCache.currentScene.Window
	size := win.GetSize()
	if frameCache.cachedImage != nil {
		frameCache.cachedImage.Dispose()
	}

	frameCache.cachedImage = ebiten.NewImage(size.Width, size.Height)
	frameCache.cachedImage.Fill(win.GetBackground())
}

func (frameCache *FrameCache) InvalidateCache() {
	if frameCache.cachedImage != nil {
		frameCache.cachedImage.Dispose()
		frameCache.cachedImage = nil
	}
}

func (frameCache *FrameCache) CheckCacheInvalidation(newSize common.Size) {
	if frameCache.cachedImage != nil {
		bounds := frameCache.cachedImage.Bounds()
		if newSize.Width != bounds.Dx() || newSize.Height != bounds.Dy() {
			frameCache.InvalidateCache()
		}
	}
}

func (frameCache *FrameCache) RenderFrame(frameRenderer FrameRenderer) {
	frameCache.frameRenderer = frameRenderer
}

func (frameCache *FrameCache) Render(screen *ebiten.Image) {
	graph, forceInvalidation := frameCache.currentScene.SceneGraph()

	if forceInvalidation {
		frameCache.InvalidateCache()
	}

	if frameCache.cachedImage == nil {
		frameCache.creatNewCachedImage()
		frameCache.frameRenderer(graph, frameCache.cachedImage)
	}

	screen.DrawImage(frameCache.cachedImage, nil)
}

func (frameCache *FrameCache) CurrentFrame() *ebiten.Image {
	return frameCache.cachedImage
}

func NewFrameCache() FrameCache {
	return FrameCache{}
}
