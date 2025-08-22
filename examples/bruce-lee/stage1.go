package main

import (
	"fmt"
	"log"
	"image/color"
	_ "image/jpeg"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func stage1(screen *ebiten.Image) {

    if player == nil{
        player, err = initAudio(stage2MusicPath)
        player.Play()

        if err != nil {
        	log.Fatal(err)
        }
    }

	scale := ebiten.Monitor().DeviceScaleFactor()

	drawBackground(screen, logo, 20, 20, 410, 371)
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	fw, fh := ebiten.Monitor().Size()
	msg := ""
	if runtime.GOOS == "js" {
		msg += "Press F or touch the screen to enter fullscreen (again).\n"
	} else {
		msg += "Press Q to quit.\n"
	}
	msg += fmt.Sprintf("Screen size in fullscreen: %d, %d\n", fw, fh)
	msg += fmt.Sprintf("Game's screen size: %d, %d\n", sw, sh)
	msg += fmt.Sprintf("Device scale factor: %0.2f\n", scale)

	textOp := &text.DrawOptions{}
	textOp.GeoM.Translate(50*scale, 650*scale)
	textOp.ColorScale.ScaleWithColor(color.White)
	textOp.LineSpacing = 12 * ebiten.Monitor().DeviceScaleFactor() * 1.5
	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   12 * ebiten.Monitor().DeviceScaleFactor(),
	}, textOp)

	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   12 * ebiten.Monitor().DeviceScaleFactor(),
	}, textOp)

    msg = story()


	textOp.GeoM.Translate(610*scale, (400-counter)*scale)
	textOp.LineSpacing = 30 * ebiten.Monitor().DeviceScaleFactor() * 2
	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   30 * ebiten.Monitor().DeviceScaleFactor(),
	}, textOp)
}
