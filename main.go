// Code example taken from https://github.com/fyne-io/fyne/issues/3052

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"image"
	"image/color"
	"time"
)

var width = 1920
var height = 1080

var lineWidth = 100
var fps = 60

func main() {
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{width, height}

	rgbaImage := image.NewRGBA(image.Rectangle{topLeft, bottomRight})

	red := color.RGBA{255, 0, 255, 0xff}

	mainApp := app.New()
	imageWindow := mainApp.NewWindow("Images")

	// This is very slow
	//canvasToWrite := canvas.NewImageFromImage(rgbaImage)

	// This was the fix in the issue (https://github.com/fyne-io/fyne/issues/3052)
	canvasToWrite := canvas.NewRasterFromImage(rgbaImage)

	imageWindow.SetContent(canvasToWrite)
	imageWindow.Resize(fyne.NewSize(float32(width), float32(height)))

	drawImageRed := func(x int) {
		for i := 0; i < lineWidth; i++ {
			rgbaImage.Set(x+100, 100+i, red)
		}
	}

	go func() {
		pos := 0
		for true {
			drawImageRed(pos)
			pos++
			canvas.Refresh(canvasToWrite)
			time.Sleep(time.Duration(1000/fps) * time.Millisecond)
		}
	}()

	imageWindow.ShowAndRun()
}
