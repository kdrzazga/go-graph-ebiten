package main

import (
    "log"
	"time"
	"math"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "image/color"
	"github.com/hajimehoshi/ebiten/v2"
)

func stage1clock(screen *ebiten.Image){
    duration := time.Since(startTime)
    milliseconds := duration.Milliseconds()
    y := float64(800 - milliseconds/60)
    //log.Printf("Outro will be displayed at y=%f", y)
    animateText(screen, clock, 20, 10, float64(y))

    x2 := float64(1020.1)
    y2 := float64(y + 310)
    x1 := float64(690)
    thickness := 14
    color := color.RGBA{200, 20, 20, 255}

    r := 211.63*2

    alpha := (float64(y) * 6.28) / 800

    xCir := r*math.Cos(-alpha)
    yCir := r*math.Sin(-alpha)

    for i := -int(thickness / 2); i <= int(thickness/2); i++ {
        ebitenutil.DrawLine(screen, x2+xCir+float64(i), float64(y)+float64(i)+yCir, x2+float64(i), y2+float64(i), color)
    }

    log.Printf("", x1, y, x2, y2)
}
