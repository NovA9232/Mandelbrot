package main

import (
  "github.com/gen2brain/raylib-go/raylib"
)

const (
  SCREEN_W float64 = 1400
  SCREEN_H float64 = 800

	maxIteration float64 = 10
	offsetSpd float64 = 10
	zoomSpd float64 = 0.2
)

func main() {
  rl.InitWindow(int32(SCREEN_W), int32(SCREEN_H), "Particles")
  rl.SetTargetFPS(60)

	rendered := false

	var focusX, focusY float64 = 0, 0
	w, h := SCREEN_W, SCREEN_H
	var zoom float64 = 0.5

  for !rl.WindowShouldClose() {
		if rl.IsKeyDown(rl.KeyA) {   // go left
			focusX -= offsetSpd
			rendered = false
		}
		if rl.IsKeyDown(rl.KeyD) {
			focusX += offsetSpd
			rendered = false
		}
		if rl.IsKeyDown(rl.KeyW) {
			focusY -= offsetSpd
			rendered = false
		}
		if rl.IsKeyDown(rl.KeyS) {
			focusY += offsetSpd
			rendered = false
		}

		mouseMv := float64(rl.GetMouseWheelMove())
		if mouseMv != 0 {
			zoom += zoomSpd*mouseMv
			focusX = (w/2) + zoom*focusX
			focusY = (h/2) + zoom*focusY
			rendered = false
		}

		if rl.IsMouseButtonDown(0) {
			focusX += float64(rl.GetMouseX()) - w/2
			focusY += float64(rl.GetMouseY()) - h/2
			println(focusX, focusY)
			rendered = false
		}

    rl.BeginDrawing()

		if !rendered {
			rl.ClearBackground(rl.Black)
			mandelbrotFull(focusX, focusY, w, h, zoom, maxIteration)
			rendered = true
		}

		rl.EndDrawing()
  }

  rl.CloseWindow()
}
