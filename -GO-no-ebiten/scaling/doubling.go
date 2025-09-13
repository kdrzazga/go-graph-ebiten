package main

import (
    "fmt"
    "image"
    "os"

    "golang.org/x/image/bmp"
)

func main() {
    inputPath := "source-dir/chuck.bmp"
    outputPath := "destination-dir/chuck2.bmp"

    infile, err := os.Open(inputPath)
    if err != nil {
        fmt.Printf("Error opening input file: %v\n", err)
        return
    }
    defer infile.Close()

    // Decode the BMP image
    img, err := bmp.Decode(infile)
    if err != nil {
        fmt.Printf("Error decoding BMP: %v\n", err)
        return
    }

    // Double the size by replacing each pixel with four pixels
    newWidth := img.Bounds().Dx() * 2
    newHeight := img.Bounds().Dy() * 2
    scaledImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

    for y := 0; y < img.Bounds().Dy(); y++ {
        for x := 0; x < img.Bounds().Dx(); x++ {
            pixel := img.At(x+img.Bounds().Min.X, y+img.Bounds().Min.Y)
            // Map the original pixel to 2x2 block in new image
            startX := x * 2
            startY := y * 2

            // Fill 2x2 block with the pixel color
            scaledImg.Set(startX, startY, pixel)
            scaledImg.Set(startX+1, startY, pixel)
            scaledImg.Set(startX, startY+1, pixel)
            scaledImg.Set(startX+1, startY+1, pixel)
        }
    }

    outfile, err := os.Create(outputPath)
    if err != nil {
        fmt.Printf("Error creating output file: %v\n", err)
        return
    }
    defer outfile.Close()

    if err := bmp.Encode(outfile, scaledImg); err != nil {
        fmt.Printf("Error encoding BMP: %v\n", err)
        return
    }

    fmt.Println("Image doubled successfully: " + outputPath)
}
