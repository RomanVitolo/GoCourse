package main

import (
	"fmt"
	"runtime"

	"github.com/RomanVitolo/GoCourse/variables"
)

func main() {
	state, text := variables.ConvertToText(1588)
	fmt.Println(state)
	fmt.Println(text)
	os := runtime.GOOS
	if os == "linux" || os == "OS X." {
		fmt.Println("This is not Windows", os)
	} else {
		fmt.Println("This is Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("This is Linux")
	case "darwin":
		fmt.Println("This is Darwin")
	default:
		fmt.Printf("%s \n", os)
	}
}
