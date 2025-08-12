package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"os"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/audio"
    "github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
    err           error
	background    *ebiten.Image
	mplusFaceSource *text.GoTextFaceSource
    context *audio.Context
    player  *audio.Player
)

func init() {
    background, err = loadImage("broken-land.png")
    if err != nil {
        log.Fatal(err)
    }
}

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s
}

func initAudio() error {
    context = audio.NewContext(44100)
    f, err := os.Open("introduction.wav")
    if err != nil {
        return err
    }
    stream, err := wav.Decode(context, f)
    if err != nil {
        return err
    }
    player, err = audio.NewPlayer(context, stream)
    if err != nil {
        return err
    }
    return nil
}

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++

	if runtime.GOOS == "js" {
		if ebiten.IsKeyPressed(ebiten.KeyF) || len(inpututil.AppendJustPressedTouchIDs(nil)) > 0 {
			ebiten.SetFullscreen(true)
		}
	}
	if runtime.GOOS != "js" && ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scale := ebiten.Monitor().DeviceScaleFactor()

	drawBackground(screen, background)
/*
	w, h := background.Bounds().Dx(), background.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Rotate(float64(g.count%360) * 2 * math.Pi / 360)*/
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
/*	op.GeoM.Translate(float64(sw)/2, float64(sh)/2)
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(background, op)
*/
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
	textOp.GeoM.Translate(50*scale, 50*scale)
	textOp.ColorScale.ScaleWithColor(color.White)
	textOp.LineSpacing = 12 * ebiten.Monitor().DeviceScaleFactor() * 1.5
	text.Draw(screen, msg, &text.GoTextFace{
		Source: mplusFaceSource,
		Size:   12 * ebiten.Monitor().DeviceScaleFactor(),
	}, textOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.Monitor().DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func main() {
    initAudio()
    player.Play()

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Fullscreen (Ebitengine Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func drawBackground(screen *ebiten.Image, bg *ebiten.Image) {
    subImg := bg.SubImage(image.Rect(0, 0, 512, 512)).(*ebiten.Image)
    op := &ebiten.DrawImageOptions{}
    //op.GeoM.Scale(2, 2)
    screen.DrawImage(subImg, op)
}

func loadImage(path string) (*ebiten.Image, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()
    img, _, err := image.Decode(file)
    if err != nil {
        return nil, err
    }
    return ebiten.NewImageFromImage(img), nil
}
