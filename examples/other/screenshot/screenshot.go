// go get github.com/kbinani/screenshot
package main

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	numScreenshots := 20
	interval := 2 * time.Second

	fmt.Printf("Taking %d screenshots every %s...\n", numScreenshots, interval)

	for i := 1; i <= numScreenshots; i++ {
		// Capture the main display (display 0)
		img, err := screenshot.CaptureDisplay(0)
		if err != nil {
			fmt.Println("Error capturing screen:", err)
			return
		}

		// Save screenshot with timestamp in filename
		filename := filepath.Join(".", fmt.Sprintf("screenshot_%02d.png", i))
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		err = png.Encode(file, img)
		if err != nil {
			fmt.Println("Error saving screenshot:", err)
			return
		}

		fmt.Printf("Saved %s\n", filename)

		if i < numScreenshots {
			time.Sleep(interval)
		}
	}

	fmt.Println("Done!")
}
