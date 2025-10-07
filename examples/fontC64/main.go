package main

import (
    "bufio"
    "io/ioutil"
    "image/color"
    "log"
    "os"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/text"
    "golang.org/x/image/font"
    "golang.org/x/image/font/opentype"
)

const (
    screenWidth  = 800
    screenHeight = 600
)

var (
    fontFace font.Face
    caption  string
)

func init() {
    fontBytes, err := ioutil.ReadFile("C64ProMono.ttf")
    if err != nil {
        log.Fatal(err)
    }
    f, err := opentype.Parse(fontBytes)
    if err != nil {
        log.Fatal(err)
    }
    fontFace, err = opentype.NewFace(f, &opentype.FaceOptions{
        Size:    24,
        DPI:     72,
        Hinting: font.HintingFull,
    })
    if err != nil {
        log.Fatal(err)
    }
    file, err := os.Open("caption.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    caption = ""
    for _, line := range lines {
        caption += line + "\n"
    }
}

type Game struct{}

func (g *Game) Update() error {
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    //screen.Fill(nil)
    whiteColor := color.RGBA{200, 20, 20, 255}
    text.Draw(screen, caption, fontFace, 10, 50, whiteColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}

func main() {
    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Caption Display")
    if err := ebiten.RunGame(&Game{}); err != nil {
        log.Fatal(err)
    }
}
