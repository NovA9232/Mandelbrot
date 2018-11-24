package main

import (
  "github.com/fogleman/gg"
)

func mandelbrot() {
  const W = 1920
  const H = 1080
  const f = 30
  dc := gg.NewContext(W, H)
  maxIteration := 255 * f
  for i:=float64(0); i < W; i++ {
    for j:=float64(0); j < H; j++ {
      var zx float64 = ((i/W)*3.5)-2.5
      var zy float64 = ((j/H)*2)-1
      var x float64
      var y float64
      iteration := 0
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
      dc.SetRGB(float64(iteration/f), float64(iteration/f), float64(iteration/f))
      dc.SetPixel(int(i), int(j))
    }
  }

  dc.SavePNG("out.png")
}

func main() {
  mandelbrot()
}
