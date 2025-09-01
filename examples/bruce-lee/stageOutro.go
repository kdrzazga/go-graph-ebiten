package main

import (
    "log"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
    gameOverPic *ebiten.Image
	stage4Counter       float64
    )

func initStageOutro(){
    var err error

    gameOverPic, err = loadImage("pics/gameover.png")
    if err != nil {
        log.Fatal(err)
    }

    stage4Counter = 0
}

func stageOutro(screen *ebiten.Image){
    screen.Fill(color.RGBA{R: 72, G: 58, B: 170, A: 255})
    outro(screen, stage4Counter)
	drawBackground(screen, gameOverPic, 0, 0, 1583, 138)
    stage4Counter +=1
}

func outro(screen *ebiten.Image, counter float64){
    msg := "Demo written in GO language with EBITEN framework.\n"
    msg += "\n"
    msg += "Greetings to K&A+ team (including Pan Areczek)....\n"
    msg += "\nKudoz for publishing a great magazine!\n"
    msg += "Your dedication and passion shine through in every issue,\n"
    msg += "making K&A+ a true treasure for enthusiasts of retro computers,\n"
    msg += "Commodore, Amiga, and beyond.\n"
    msg += "The team’s deep knowledge, meticulous research, and love for\n"
    msg += "vintage technology are evident in the high-quality content you produce.\n"
    msg += "Arek, Leon, Maciek, Tomxx, and Beszcza — each of you bring\n"
    msg += "unique expertise and enthusiasm that contribute to the magazine’s success.\n"
    msg += "Your collective effort preserves the history and culture of\n"
    msg += "classic computing, inspiring both seasoned fans and newcomers alike.\n"
    msg += "Keep up the fantastic work—your passion keeps the spirit of\n"
    msg += "retro computing alive and thriving!"
    y := 1000 - stage4Counter
    //log.Printf("Outro will be displayed at y=%f", y)
    animateText(screen, msg, 30, 10, y)
}
