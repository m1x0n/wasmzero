package main

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/gif"
	_ "image/png"
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
	Arena  *ebiten.Image
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

	arena, _, err := ebitenutil.NewImageFromFileSystem(imageFiles, "img/arena.png")
	if err != nil {
		log.Fatal(err)
	}
	wz.Arena = arena

	return wz
}

func (wz *WasmZero) Update() error {
	wz.Count++
	return nil
}

func (wz *WasmZero) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "WasmZero")

	arenaOp := &ebiten.DrawImageOptions{}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2-32, -float64(frameHeight)/2+32)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (wz.Count / 8) % frameCount

	offsetX := 0

	screen.DrawImage(wz.Arena, arenaOp)

	if i >= 3 {
		offsetX = i*32 + 16
		// Render knees posture
		op.GeoM.Translate(0, 0)
		screen.DrawImage(wz.Images[2], op)

		op.GeoM.Translate(float64(offsetX), 0)
		screen.DrawImage(wz.Images[i], op)
	} else {
		op.GeoM.Translate(float64(offsetX), 0)
		screen.DrawImage(wz.Images[i], op)
	}
}

func (wz *WasmZero) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

//go:embed img/*.gif
//go:embed img/*.png
var imageFiles embed.FS

func main() {
	// Scale x2
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("WasmZero!")

	if err := ebiten.RunGame(NewWasmZero()); err != nil {
		log.Fatal(err)
	}
}
