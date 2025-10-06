package main

import (
	_ "image/jpeg"
	"time"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
    err       error
    startTime time.Time
    clock     []string
    )

type Game struct {
	count int
}

func (g *Game) Update() error {
    duration := time.Since(startTime)
    log.Println("Duration: ", duration)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
        for _, line := range clock {
            log.Println(line)
        }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.Monitor().DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func init() {
    clock, err = readTextFile("clock.txt")

    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    startTime = time.Now()

    ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("ASCII")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
