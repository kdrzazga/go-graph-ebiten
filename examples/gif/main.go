package main

import (
    //"image"
    "image/gif"
    "image/color"
    "log"
    "os"
    "time"

    "github.com/hajimehoshi/ebiten/v2"
    //"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type GIFAnimator struct {
    Frames       []*ebiten.Image
    FrameDelays  []int // in milliseconds
    currentFrame int
    lastTime     time.Time
}

func NewGIFAnimator(gifPath string) (*GIFAnimator, error) {
    f, err := os.Open(gifPath)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    g, err := gif.DecodeAll(f)
    if err != nil {
        return nil, err
    }

    frames := make([]*ebiten.Image, len(g.Image))
    for i, img := range g.Image {
        // Convert each image to *ebiten.Image
        ebitenImg := ebiten.NewImageFromImage(img)
        frames[i] = ebitenImg
    }

    return &GIFAnimator{
        Frames:       frames,
        FrameDelays:  g.Delay,
        currentFrame: 0,
        lastTime:     time.Now(),
    }, nil
}

func (a *GIFAnimator) Update() {
    now := time.Now()
    elapsed := now.Sub(a.lastTime)
    delay := a.FrameDelays[a.currentFrame] * 10 // gif.Delay units are in 10ms increments

    if elapsed >= time.Duration(delay)*time.Millisecond {
        a.currentFrame = (a.currentFrame + 1) % len(a.Frames)
        a.lastTime = now
    }
}

func (a *GIFAnimator) Draw(screen *ebiten.Image, x, y float64) {
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(x, y)
    screen.DrawImage(a.Frames[a.currentFrame], op)
}

type Game struct {
    animator *GIFAnimator
}

func (g *Game) Update() error {
    g.animator.Update()
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(color.Black)
    g.animator.Draw(screen, 0, 0)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return 200, 200
}

func main() {
    animator, err := NewGIFAnimator("dragon.gif")
    if err != nil {
        log.Fatal(err)
    }

    game := &Game{animator: animator}
    ebiten.SetWindowSize(640, 480)
    ebiten.SetWindowTitle("GIF in Ebiten")
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
