package main

import (
    "log"
	"time"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

func stage1clock(screen *ebiten.Image){
    duration := time.Since(startTime)
    milliseconds := duration.Milliseconds()
    y := 1000 - milliseconds/10
    //log.Printf("Outro will be displayed at y=%f", y)
    animateText(screen, clock, 20, 10, float64(y))

    x2 := 850.5
    y2 := float64(y + 250)
    x1 := float64(400)
    thickness := 4
    col := color.RGBA{200, 20, 20, 255}

    for i := -int(thickness / 2); i <= int(thickness/2); i++ {
        ebitenutil.DrawLine(screen, x1+float64(i), float64(y+83)+float64(i), x2+float64(i), y2+float64(i), col)
    }

    log.Printf("", x1, y, x2, y2)
}
