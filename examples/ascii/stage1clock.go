package main

import (
	"time"
	"github.com/hajimehoshi/ebiten/v2"
)

func stage1clock(screen *ebiten.Image){
    duration := time.Since(startTime)
    milliseconds := duration.Milliseconds()
    y := 1000 - milliseconds/10
    //log.Printf("Outro will be displayed at y=%f", y)
    animateText(screen, clock, 20, 10, float64(y))
}