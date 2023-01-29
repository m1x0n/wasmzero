package main

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/gif"
	"log"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameWidth  = 132
	frameHeight = 132
	frameCount  = 7
)

type WasmZero struct {
	Images []*ebiten.Image
	Count  int
}

func NewWasmZero() *WasmZero {
	wz := &WasmZero{
		Images: make([]*ebiten.Image, 0),
	}

	for i := 1; i <= 7; i++ {
		gifFile, err := imageFiles.Open(fmt.Sprintf("img/f0%d.gif", i))
		if err != nil {
			log.Fatal(err)
		}
		defer gifFile.Close()

		gifImage, err := gif.Decode(gifFile)
		if err != nil {
			log.Fatal(err)
		}

		img := ebiten.NewImageFromImage(gifImage)

		wz.Images = append(wz.Images, img)
	}

	return wz
}

func (wz *WasmZero) Update() error {
	wz.Count++
	return nil
}

func (wz *WasmZero) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "WasmZero")

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (wz.Count / 8) % frameCount
	screen.DrawImage(wz.Images[i], op)
}

func (wz *WasmZero) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

//go:embed img/*.gif
var imageFiles embed.FS

func main() {
	// Scale x2
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("WasmZero!")

	if err := ebiten.RunGame(NewWasmZero()); err != nil {
		log.Fatal(err)
	}
}
