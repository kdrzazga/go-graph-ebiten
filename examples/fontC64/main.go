package main

import (
    "io/ioutil"
    "bufio"
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

func createFont(fontPath string, size float64) (font.Face, error) {
    fontBytes, err := ioutil.ReadFile(fontPath)
    if err != nil {
        return nil, err
    }
    f, err := opentype.Parse(fontBytes)
    if err != nil {
        return nil, err
    }
    face, err := opentype.NewFace(f, &opentype.FaceOptions{
        Size:    size,
        DPI:     72,
        Hinting: font.HintingFull,
    })
    if err != nil {
        return nil, err
    }
    return face, nil
}

func readCaption(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return "", err
    }

    caption := ""
    for _, line := range lines {
        caption += line + "\n"
    }
    return caption, nil
}

func init() {
    var err error
    fontFace, err = createFont("C64ProMono.ttf", 24)
    if err != nil {
        log.Fatal(err)
    }

    caption, err = readCaption("caption.txt")
    if err != nil {
        log.Fatal(err)
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
