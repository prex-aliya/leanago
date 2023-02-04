package main

import (
        "fmt"
        "math/rand"
)

func main() {
        rand.Seed(10)
        fmt.Println("My Favorite Number Is",
                rand.Intn(10))
}
