package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	subjects := []string{"my cat", "my WiFi", "my dog"}
	actions := []string{"joined", "attacked", "got angry with"}
	ending := []string{"tibetan monks", "the communist party", "a cult"}
	allSlices := [][]string{
		subjects,
		actions,
		ending,
	}

	var words []string

	for _, e := range allSlices {
		currentSlice := e
		words = append(words, currentSlice[rand.Intn(len(currentSlice))])
	}

	fmt.Println("Sorry, ", strings.Join(words, " "), ".")
}
