package main

import (
    "log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
    returnOfFuryAnimator *GIFAnimator
)

func initStage4(){
    var err error

    returnOfFuryAnimator, err = NewGIFAnimator("pics/return-of-fury.gif")
    if err != nil {
        log.Fatal(err)
    }
}

func stage4(screen *ebiten.Image, counter float64){
    returnOfFuryAnimator.Update()
    returnOfFuryAnimator.Draw(screen, 0, 0)
}
