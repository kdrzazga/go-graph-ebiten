package main

import (
    "os"
    "log"
    "bufio"
    "bytes"
	"image/color"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
)

var(
	mplusFaceSource     *text.GoTextFaceSource
)

func initLib(){
    s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
    	if err != nil {
    		log.Fatal(err)
    	}
	mplusFaceSource = s
}

func readTextFile(filename string) (string, error) {

    file, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer file.Close()

    var lines string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        lines += scanner.Text()
        lines += "\n"
    }

    if err := scanner.Err(); err != nil {
        return "", err
    }

    return lines, nil
}

func animateText(screen *ebiten.Image, msg string, size float64, x float64, y float64){
	scale := ebiten.Monitor().DeviceScaleFactor()
    textOp := &text.DrawOptions{}

    textOp.GeoM.Translate(x, y)
    textOp.ColorScale.ScaleWithColor(color.White)
    textOp.LineSpacing = size * scale * 1
    text.Draw(screen, msg, &text.GoTextFace{
    	Source: mplusFaceSource,
    	Size:   size * scale,
    }, textOp)
}
