package main

import (
    "log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
    returnOfFuryAnimator *GIFAnimator
    chuckNorrisAnimator *GIFAnimator
    kickdownAnimator *GIFAnimator
    kickingAnimator *GIFAnimator
    returnOfFuryImg *ebiten.Image
)

func initStage4(){
    var err error

    returnOfFuryAnimator, err = NewGIFAnimator("pics/return-of-fury.gif", false)
    if err != nil {
        log.Fatal(err)
    }
    chuckNorrisAnimator, err = NewGIFAnimator("pics/brucelee-chucknorris.gif", false)
    if err != nil {
        log.Fatal(err)
    }
    kickdownAnimator, err = NewGIFAnimator("pics/kickdown.gif", false)
    if err != nil {
        log.Fatal(err)
    }
    returnOfFuryImg, err = loadImage("pics/return-of-fury.jpg")
    if err != nil {
        log.Fatal(err)
    }
    kickingAnimator, err = NewGIFAnimator("pics/kicking.gif", false)
    if err != nil {
        log.Fatal(err)
    }
}

func stage4(screen *ebiten.Image, counter float64){

    if (counter < 2000 + stage3Timeout) {
        returnOfFuryAnimator.Update()
        returnOfFuryAnimator.Draw(screen, 0, 0)
    } else if ((counter < 3000 + stage3Timeout) || (counter > 12000 + stage3Timeout && counter < 13000 + stage3Timeout)) {
        drawBackgroundScaled(screen, returnOfFuryImg, 0, 0, 400, 245, float64(1))
        chuckNorrisAnimator.Update()
        chuckNorrisAnimator.Draw(screen, 400, 245)
    } else if (counter < 8300 + stage3Timeout){
        drawBackgroundScaled(screen, returnOfFuryImg, 0, 0, 400, 245, float64(1))
        chuckNorrisAnimator.Draw(screen, 400, 245)
        kickingAnimator.Update()
        kickingAnimator.Draw(screen, 10, 245+280)
    } else {
        drawBackgroundScaled(screen, returnOfFuryImg, 0, 0, 400, 245, float64(1))
        chuckNorrisAnimator.Draw(screen, 400, 245)

        kickingAnimator.Draw(screen, 10, 245+280)

        kickdownAnimator.Draw(screen, 400+468, 245+280)
        kickdownAnimator.Update()
        chuckNorrisAnimator.Reset()
    }

    if themePlayer == nil{
        themePlayer, err = initAudio(stage4MusicPath)
        themePlayer.Play()

        if err != nil {
        	log.Fatal(err)
        }
    }
}
