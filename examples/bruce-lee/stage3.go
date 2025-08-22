package main

import (
	"fmt"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func stage3(screen *ebiten.Image){
    drawBackground(screen, background, shiftX, shiftY, 2555, 705)

    if player2 == nil{
        player2, err = initAudio(stage3MusicPath)
        player2.Play()

        if err != nil {
        	log.Fatal(err)
        }
    }

    move()
}

func move() {
    if (shiftX > moveSpeed) && (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft)) {
        shiftX -= moveSpeed
    } else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
        shiftX += moveSpeed
    }
    if (shiftY > moveSpeed) && (ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp)) {
        shiftY -= moveSpeed
    } else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
        shiftY += moveSpeed
    }
    fmt.Println(" [", shiftX, shiftY, "] ")
}
