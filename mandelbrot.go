package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func mandelbrotPixel(i, j, W, H, zoom, maxIteration float64) float64 {
	var (
		zx float64 = ((i/W)*2.5)/zoom
		zy float64 = ((j/H)*1.5)/zoom
		x float64
		y float64
		iteration float64
	)
	for ((x*x + y*y < 4) && (iteration < maxIteration)) {
		xtemp := x*x - y*y + zx
		ytemp := 2*x*y + zy
		if (x == xtemp && y == ytemp) {
			iteration = maxIteration
			break
		}
		x = xtemp
		y = ytemp
		iteration++
	}
	return iteration
}

func mandelbrotFull(focusX, focusY, W, H, zoom, maxIteration float64) {
	println(focusX, focusY)
  for i:=float64(0); i < W; i++ {
		for j:=float64(0); j < H; j++ {
			iteration := mandelbrotPixel(i+focusX, j+focusY, W, H, zoom, maxIteration)
			clr := rl.NewColor(uint8((iteration/maxIteration)*255), uint8((iteration/maxIteration)*255), uint8((iteration/maxIteration)*255), 255)
			rl.DrawPixel(int32(i), int32(j), clr)
		}
	}
}
