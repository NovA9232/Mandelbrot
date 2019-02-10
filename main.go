package main

import (
  "github.com/gen2brain/raylib-go/raylib"
)

const (
  SCREEN_W float64 = 1400
  SCREEN_H float64 = 800

	maxIteration float64 = 100
	offsetSpd float64 = 10
	zoomSpd float64 = 0.2
)

func main() {
  rl.InitWindow(int32(SCREEN_W), int32(SCREEN_H), "Particles")
  rl.SetTargetFPS(60)

	rendered := false

	var focusX, focusY float64 = -SCREEN_W/2, -SCREEN_H/2
	var w, h float64 = SCREEN_W, SCREEN_H
	var zoom float64 = 1

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
			dz := zoomSpd*mouseMv*zoom
			oldZoom := zoom
			zoom += dz
			w, h = w/zoom, h/zoom
			focusX -= (w * (1 - 1/(zoom/oldZoom)))
			focusY -= (h * (1 - 1/(zoom/oldZoom)))
			println(focusX, focusY, dz, w, h, "focX, focY, dz, w, h")
			rendered = false
		}

		if rl.IsMouseButtonDown(0) {
			focusX += (float64(rl.GetMouseX()) - SCREEN_W/2)
			focusY += (float64(rl.GetMouseY()) - SCREEN_H/2)
			println(focusX, focusY)
			rendered = false
		}

    rl.BeginDrawing()

		if !rendered {
			rl.ClearBackground(rl.Black)
			mandelbrotFull(focusX, focusY, SCREEN_W, SCREEN_H, zoom, maxIteration)
			rl.DrawPixel(-int32(focusX), -int32(focusY), rl.Red)
			rendered = true
		}

		rl.EndDrawing()
  }

  rl.CloseWindow()
}
