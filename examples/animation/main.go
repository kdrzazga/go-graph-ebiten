package main

import (
    "log"
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/audio"
)

var currentScene Scene

type Scene interface {
    Update() error
    Draw(screen *ebiten.Image)
}

const (
    screenWidth  = 474
    screenHeight = 299

    frameOX     = 0
    frameOY     = 0
    frameWidth  = 55
    frameHeight = 71
    frameCount  = 3
)

var (
    runnerImage      *ebiten.Image
    backgroundImage1 *ebiten.Image
    backgroundImage2 *ebiten.Image
    backgroundImage3 *ebiten.Image
    newYork          *ebiten.Image
    yieArKF          *ebiten.Image
    sidney           *ebiten.Image

    posX = float64(474) * 0.4
    posY = float64(220 - 71/2)

    circleX = float64(474) * 0.5
    circleY = float64(299) * 0.8

    movement = float64(3)

    context *audio.Context
    player  *audio.Player
)

func main() {
    currentScene = &Scene1{
        count: 0,
        t:     0,
        tDir:  1,
    }

    ebiten.SetWindowSize(screenWidth * 2, screenHeight * 2)
    ebiten.SetWindowTitle("Local Karate Minus")
    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}

type Game struct{}

func (g *Game) Update() error {
    return currentScene.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
    currentScene.Draw(screen)
}

func (g *Game) Layout(outsideW, outsideH int) (int, int) {
    return screenWidth * 2, screenHeight * 2
}
