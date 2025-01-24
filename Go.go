package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    fmt.Println("Hello, World!")
    rand.Seed(time.Now().UnixNano())
    fmt.Println(rand.Intn(100) + 1)
    fmt.Println(time.Now())
}
