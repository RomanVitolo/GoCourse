package main

import (
	"fmt"

	"github.com/RomanVitolo/GoCourse/variables"
)

func main() {
	state, text := variables.ConvertToText(1588)
	fmt.Println(state)
	fmt.Println(text)
}
