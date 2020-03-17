package main

import (
	"fmt"
	"math/rand"
)

const (
	RandomN     = 3927
	RandomSpeed = 39
)

func main() {
	rand.Seed(RandomSpeed)
	fmt.Println("Hello world!", rand.Intn(RandomN))
	fmt.Printf("Random int: %v\n", rand.Intn(RandomN))
}
