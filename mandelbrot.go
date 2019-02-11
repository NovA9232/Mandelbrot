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
	halfW, halfH := W/2, H/2

	mandelbrotArea(focusX, focusY, 0, 0, halfW, halfH, W, H, zoom, maxIteration)
	mandelbrotArea(focusX, focusY, halfW, 0, halfW, halfH, W, H, zoom, maxIteration)
	mandelbrotArea(focusX, focusY, 0, halfH, halfW, halfH, W, H, zoom, maxIteration)
	mandelbrotArea(focusX, focusY, halfW, halfH, halfW, halfH, W, H, zoom, maxIteration)
}

func mandelbrotArea(offsetX, offsetY, scrnX, scrnY, maxW, maxH, W, H, zoom, maxIteration float64) { // W and H are total width and height
  for i:=float64(scrnX); i < maxW+scrnX; i++ {
		for j:=float64(scrnY); j < maxH+scrnY; j++ {
			iteration := mandelbrotPixel(i+offsetX, j+offsetY, W, H, zoom, maxIteration)
			rl.DrawPixel(int32(i), int32(j), rl.NewColor(uint8((iteration/maxIteration)*255), uint8((iteration/maxIteration)*255), uint8((iteration/maxIteration)*255), 255))
		}
	}
	rl.DrawRectangleLines(int32(scrnX), int32(scrnY), int32(maxW), int32(maxH), rl.Red)
}
