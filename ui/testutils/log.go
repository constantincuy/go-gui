package testutils

import (
	"fmt"
	"github.com/constantincuy/go-gui/ui/common"
	"image"
)

func SizeError(message string, want, got common.Size) string {
	return fmt.Sprintf("%s; want (Width: %d, Height: %d), got (Width: %d, Height: %d)", message, want.Width, want.Height, got.Width, got.Height)
}

func PositionError(message string, want, got image.Point) string {
	return fmt.Sprintf("%s; want (X: %d, Y: %d), got (X: %d, Y: %d)", message, want.X, want.Y, got.X, got.Y)
}

func BoolError(message string, want, got bool) string {
	return fmt.Sprintf("%s; want `%t`, got `%t`", message, want, got)
}

func StringError(message, want, got string) string {
	return fmt.Sprintf("%s; want `%s`, got `%s`", message, want, got)
}

func IntError(message string, want, got int) string {
	return fmt.Sprintf("%s; want `%d`, got `%d`", message, want, got)
}
