package main

import (
    "embed"
    "fmt"
    "log"
)

//go:embed file.txt
var fileContent embed.FS

func main() {
    contentBytes, err := fileContent.ReadFile("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    content := string(contentBytes)
    fmt.Println("Content of embedded file:")
    fmt.Println(content)
}
