package main

import (
    "io/ioutil"
    "bufio"
    "image/color"
    "strings"
    "log"
    "os"

    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/text"
    "golang.org/x/image/font"
    "golang.org/x/image/font/opentype"
)

const (
    screenWidth  = 1200
    screenHeight = 600
)

var (
    fontFace font.Face
    fontFace2 font.Face
    caption  string
    caption2  string
    captionShadow  string
    captionEnjoy2  string
    X int
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

    caption = strings.ReplaceAll(caption, ".", " ")

    return caption, nil
}

func init() {
    var err error
    fontFace, err = createFont("C64ProMono.ttf", 7)
    if err != nil {
        log.Fatal(err)
    }
    fontFace2, err = createFont("C64ProMono.ttf", 72)
    if err != nil {
        log.Fatal(err)
    }

    caption, err = readCaption("ascii-art.txt")
    if err != nil {
        log.Fatal(err)
    }

    caption2, err = readCaption("caption.txt")
    if err != nil {
        log.Fatal(err)
    }
    captionShadow, err = readCaption("enjoy.txt")
    if err != nil {
        log.Fatal(err)
    }
    captionEnjoy2, err = readCaption("enjoy2.txt")
    if err != nil {
        log.Fatal(err)
    }

}

type Game struct{
    X int
}

func (g *Game) Update() error {
    g.X--
    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    //screen.Fill(nil)
    whiteColor := color.RGBA{200, 200, 200, 255}
    cyanColor := color.RGBA{5, 255, 255, 255}
    purpleColor := color.RGBA{255, 5, 255, 255}
    greenColor := color.RGBA{5, 255, 5, 255}
    text.Draw(screen, caption2, fontFace2, 10, 100, whiteColor)
    text.Draw(screen, caption, fontFace, -500 - g.X, 150, cyanColor)
    text.Draw(screen, captionShadow, fontFace, 5, 210, purpleColor)
    text.Draw(screen, captionEnjoy2, fontFace, g.X, 250, greenColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}

func main() {
    game := &Game{}
    game.X = screenWidth
    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Caption Display")
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
