package main

import (
	"go3dngine/buffer"
	"go3dngine/primitives"
	"go3dngine/types"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

const (
	defaultWidth  = 800
	defaultHeight = 800
)

func main() {
	mApp := app.New()
	w := mApp.NewWindow("GO 3D Ngine")
	w.Resize(fyne.NewSize(defaultWidth, defaultHeight))

	b := buffer.New(defaultWidth, defaultHeight, &w)

	p1 := types.Vector2{X: 0, Y: 400}
	p2 := types.Vector2{X: 800, Y: 400}
	primitives.DrawLine(p2, p1, color.RGBA{255, 255, 0, 255}, &b)

	t := [3]types.Vector2{
		{X: 50, Y: 550},
		{X: 400, Y: 50},
		{X: 750, Y: 550},
	}
	primitives.DrawTriangle(t, color.White, &b)

	b.Dump()

	w.ShowAndRun()
}
