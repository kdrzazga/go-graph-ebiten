package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "log"
    )

type Scene4 struct {
    count int
}

func (s *Scene4) Update() error {
    s.count++
    moveSprite()

     if posX < 0 {
         currentScene = &Scene3{}
         log.Println("Scene3 loaded")
         posX = float64(screenWidth - 3)
     }
    return nil
}

func (s *Scene4) Draw(screen *ebiten.Image) {
    drawBackground(screen, yieArKF)
    drawSprite(s.count, screen)
}
