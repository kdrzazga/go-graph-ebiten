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
    c64gameAnimator *GIFAnimator
    returnOfFuryImg *ebiten.Image
	bigPic          *ebiten.Image
	c64Pic          *ebiten.Image
	flyingKickPic   *ebiten.Image
	bigPicY             int
	shiftX4             int
	extraDelay          int
	bruceleePosition    int
	stage4Counter       float64
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
    bigPic, err = loadImage("pics/big3.png")
    if err != nil {
        log.Fatal(err)
    }
    c64Pic, err = loadImage("pics/c64.png")
    if err != nil {
        log.Fatal(err)
    }
    flyingKickPic, err = loadImage("pics/fk.png")
    if err != nil {
        log.Fatal(err)
    }
    c64gameAnimator, err = NewGIFAnimator("pics/c64game.gif", false)
    if err != nil {
        log.Fatal(err)
    }

    extraDelay = 0
    bigPicY = 2500
    shiftX4 = 0
    bruceleePosition = -4500
    stage4Counter = 0
}

func stage4(screen *ebiten.Image, counter float64){
    extraDelay += 1
    stage4Counter += 1
    //log.Println(extraDelay, (extraDelay % 3), (extraDelay %3 == 0), bigPicY)
    if (extraDelay % 4 > 0 && bigPicY > 0){
        bigPicY -= 1
    }

    if (counter > 50000+ stage3Timeout){
        outro(screen, stage4Counter)
    }

    if (counter > (30000 + stage3Timeout)){
        bruceleePosition +=2

        op := &ebiten.DrawImageOptions{}
        op.GeoM.Translate(float64(bruceleePosition), float64(350))
        if (bruceleePosition > -120){
            screen.DrawImage(flyingKickPic, op)
        }
        drawBackgroundScaled(screen, c64Pic, 0, 0, 1200, 722, float64(1))
        c64gameAnimator.Draw(screen, float64(606), float64(224))
        c64gameAnimator.Update()
    } else if (counter > (25000 + stage3Timeout)){
        bigPicY += 2
    } else if (counter > 15000 + stage3Timeout){
        shiftX4 -= 1
    }

    if (counter < (30000 + stage3Timeout)){
        drawBackground(screen, bigPic, shiftX4, bigPicY, 940,811)
    }

    if (counter < 2000 + stage3Timeout) {
        returnOfFuryAnimator.Update()
        returnOfFuryAnimator.Draw(screen, float64(shiftX4), 0)
    } else if ((counter < 3000 + stage3Timeout) || (counter > 12000 + stage3Timeout && counter < 13000 + stage3Timeout)) {
        drawBackgroundScaled(screen, returnOfFuryImg, 0, 0, 400+shiftX4, 245, float64(1))
        chuckNorrisAnimator.Update()
        chuckNorrisAnimator.Draw(screen, 400+float64(shiftX4), 245)
    } else if (counter < 8300 + stage3Timeout){
        drawBackgroundScaled(screen, returnOfFuryImg, 0, 0, 400+shiftX4, 245, float64(1))
        chuckNorrisAnimator.Draw(screen, 400+float64(shiftX4), 245)
        kickingAnimator.Update()
        kickingAnimator.Draw(screen, 10+float64(shiftX4), 245+280)
    } else if (counter < 30000 + stage3Timeout){
        drawBackgroundScaled(screen, returnOfFuryImg, shiftX4, 0, 400, 245, float64(1))
        chuckNorrisAnimator.Draw(screen, 400+float64(shiftX4), 245)

        kickingAnimator.Draw(screen, 10+float64(shiftX4), 245+280)

        kickdownAnimator.Draw(screen, 400+468+float64(shiftX4), 245+280)
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

func outro(screen *ebiten.Image, counter float64){
    msg := "Greetings to K&A+ team (including Pan Areczek)....\n"
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
    y := 8900 - stage4Counter
    //log.Printf("Outro will be displayed at y=%f", y)
    animateText(screen, msg, 30, 610, y)
}
