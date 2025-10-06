package main

import (
    "bufio"
    "os"
)

func readTextFile(filename string) ([]string, error) {
    // Open the file
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)

    // Read line by line
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    // Check for scanner errors
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return lines, nil
}
