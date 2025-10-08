package main

import (
	_ "image/png"
    "image/color"
    "io/ioutil"
    "strings"
    "embed"
    "math"
    "log"

    "github.com/hajimehoshi/ebiten/v2/text"
    "github.com/hajimehoshi/ebiten/v2"
    "golang.org/x/image/font/opentype"
    "golang.org/x/image/font"
)

const (
    screenWidth  = 1200
    screenHeight = 600
)

var (
    fontFace font.Face
    fontFace2 font.Face
    handwrite  string
    caption2  string
    captionShadow  string
    captionEnjoy2  string
    captionKnight  string
    //go:embed knight.txt enjoy2.txt enjoy.txt caption.txt handwrite.txt
    fileContent embed.FS
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
    contentBytes, err := fileContent.ReadFile(filePath)
        if err != nil {
            log.Fatal(err)
        }
    content := string(contentBytes)
    caption := strings.ReplaceAll(content, ".", " ")

    return caption, nil
}

func init() {
    var (
        err1, err2, err3, err4, err5, err6, err7 error
    )

    fontFace, err1 = createFont("C64ProMono.ttf", 7)
    fontFace2, err2 = createFont("C64ProMono.ttf", 72)
    handwrite, err3 = readCaption("handwrite.txt")
    caption2, err4 = readCaption("caption.txt")
    captionShadow, err5 = readCaption("enjoy.txt")
    captionEnjoy2, err6 = readCaption("enjoy2.txt")
    captionKnight, err7 = readCaption("knight.txt")

    if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil {
        log.Fatalf("Error loading resources: %v %v %v %v %v %v", err1, err2, err3, err4, err5, err6, err7)
    }
}

type Game struct{
    X int
    knightImage *ebiten.Image
    counter int
}

func (g *Game) Update() error {
    g.X--
    g.counter++

    if (g.counter > 2222){
        reset(g)
    }
    return nil
}

func reset(g *Game){
    g.X = 0
    g.counter = 0
    log.Println("Animation reset")
}

func (g *Game) Draw(screen *ebiten.Image) {
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(900, 0)
    screen.DrawImage(g.knightImage, op)
    //screen.Fill(nil)
    redColor := color.RGBA{255, 20, 20, 255}
    cyanColor := color.RGBA{5, 255, 255, 255}
    purpleColor := color.RGBA{255, 5, 255, 255}
    greenColor := color.RGBA{5, 255, 5, 255}
    text.Draw(screen, caption2, fontFace2, 10, 100, redColor)
    text.Draw(screen, handwrite, fontFace, -500 - g.X, 150, cyanColor)
    text.Draw(screen, captionShadow, fontFace, 5, 210, purpleColor)
    text.Draw(screen, captionEnjoy2, fontFace, g.X, 250, greenColor)
    x := 200 + 200*math.Sin(float64(g.X)*float64(math.Pi)/float64(200))
    text.Draw(screen, captionKnight, fontFace, int(x), 333, cyanColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}

func main() {
    game := &Game{X : screenWidth}
    //ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetFullscreen(true)
    ebiten.SetWindowTitle("Caption Display")

    knightImage, err := loadImage("knight.png")
     if err  != nil {
        log.Fatal(err)
     }

    game.knightImage = knightImage

    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
