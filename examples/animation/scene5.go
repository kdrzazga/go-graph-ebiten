package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "log"
    )

type Scene5 struct {
    count int
}

func (s *Scene5) Update() error {
    s.count++
    moveSprite()

     if posX < 0 {
         currentScene = &Scene4{
                count: 0,
                t:     0,
                tDir:  1,
            }
         log.Println("Scene4 loaded")
         posX = float64(screenWidth - 3)
     }
    return nil
}

func (s *Scene5) Draw(screen *ebiten.Image) {
    drawBackground(screen, sidney)
    drawSprite(s.count, screen)
}
