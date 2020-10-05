package buffer

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"github.com/fogleman/gg"
)

// Buffer data type to draw stuff
type Buffer struct {
	width   int
	height  int
	screen  *fyne.Window
	context *gg.Context
}

// New constructor for Buffer class
func New(w, h int, s *fyne.Window) Buffer {
	context := gg.NewContext(w, h)
	return Buffer{w, h, s, context}
}

// DrawPixel draws a pixel on screen at pos x,y with color c
func (b *Buffer) DrawPixel(x, y int, c color.Color) {
	b.context.SetColor(c)
	b.context.SetPixel(x, y)
}

// Dump draws the content of the buffer into the screen
func (b *Buffer) Dump() {
	s := *b.screen
	raster := canvas.NewRasterFromImage(b.context.Image())
	s.SetContent(raster)
}
