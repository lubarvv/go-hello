package main

import (
    "fmt"
    "github.com/lubarvv/hello/test"
)

func main() {
    fmt.Printf("Hello, World!\n")
    fmt.Printf("Здарова, Вован!\n")

    x, y := test.GetNumbers()
    fmt.Printf("First number: %d\nSecond number %d\n", x, y)
}