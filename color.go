package main

import (
  "golang.org/x/tour/pic"
  "image"
  "image/color"
  "math"
)

type Image struct{}

func (i Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, 300, 300)
}

func (i Image) At(x, y int) color.Color {
  var x1, y1 = float64(x), float64(y)

  dr := ((x1-150)*(x1-150) + (y1-150)*(y1-150) - 100*100)
//  var r, g, a uint8
//  if math.Abs(dr) < 100 {
//    r, g, a = 210, 200, 255
//  } else {
//    r, g, a = 0, 100, 255
//  }

  var a = uint8(math.Abs(dr)/200)
  return color.RGBA{100, 155, 255, a}
}

func main() {
  m := Image{}
  pic.ShowImage(m)
}
