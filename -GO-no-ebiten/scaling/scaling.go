package main

import (
    "fmt"
    "image"
    "image/png"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"

    "golang.org/x/image/draw"
)

func main() {
    sourceDir := "source-dir"
    destDir := "destination-dir"

    if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
        log.Fatalf("Failed to create destination directory: %v", err)
    }

    files, err := ioutil.ReadDir(sourceDir)
    if err != nil {
        log.Fatalf("Failed to read source directory: %v", err)
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }
        if strings.ToLower(filepath.Ext(file.Name())) != ".png" {
            continue
        }

        inputPath := filepath.Join(sourceDir, file.Name())
        outputPath := filepath.Join(destDir, file.Name())

        if err := processImage(inputPath, outputPath); err != nil {
            log.Printf("Failed to process %s: %v", file.Name(), err)
        } else {
            fmt.Printf("Processed %s\n", file.Name())
        }
    }
}

func processImage(inputPath, outputPath string) error {
    infile, err := os.Open(inputPath)
    if err != nil {
        return fmt.Errorf("failed to open input image: %w", err)
    }
    defer infile.Close()

    img, err := png.Decode(infile)
    if err != nil {
        return fmt.Errorf("failed to decode PNG: %w", err)
    }

    newWidth := img.Bounds().Dx() * 2
    newHeight := img.Bounds().Dy() * 2

    scaledImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
    draw.CatmullRom.Scale(scaledImg, scaledImg.Bounds(), img, img.Bounds(), draw.Over, nil)

    outFile, err := os.Create(outputPath)
    if err != nil {
        return fmt.Errorf("failed to create output file: %w", err)
    }
    defer outFile.Close()

    if err := png.Encode(outFile, scaledImg); err != nil {
        return fmt.Errorf("failed to encode PNG: %w", err)
    }

    return nil
}
