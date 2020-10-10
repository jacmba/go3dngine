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

	yellow := color.RGBA{255, 255, 0, 255}

	p1 := types.Vector2{X: 0, Y: 400}
	p2 := types.Vector2{X: 800, Y: 400}
	primitives.DrawLine(p2, p1, yellow, &b)

	p1 = types.Vector2{X: 400, Y: 0}
	p2 = types.Vector2{X: 400, Y: 800}
	primitives.DrawLine(p1, p2, yellow, &b)

	t := [3]types.Vector2{
		{X: 50, Y: 550},
		{X: 400, Y: 50},
		{X: 750, Y: 550},
	}
	primitives.DrawTriangle(t, color.White, &b)

	b.Dump()

	w.ShowAndRun()
}
