package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	subjects := []string{"my cat", "my WiFi", "my dog"}
	actions := []string{"joined", "attacked", "got angry with"}
	endings := []string{"tibetan monks", "the communist party", "a cult"}
	allSlices := [][]string{
		subjects,
		actions,
		endings,
	}

	var words []string

	for _, e := range allSlices {
		words = append(words, pick(e))
	}

	fmt.Printf("Sorry, %s.\n", strings.Join(words, " "))
}

func pick(slice []string) string {
	return slice[rand.Intn(len(slice))]
}
