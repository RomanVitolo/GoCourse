package main

import (
	"fmt"
	"runtime"

	"github.com/RomanVitolo/GoCourse/exercises"
	"github.com/RomanVitolo/GoCourse/keyboard"
	"github.com/RomanVitolo/GoCourse/variables"
)

func main() {

	//GetExercises()
	//GetOS()
	keyboard.GetNumbers()
}

func GetInputs() {

}

func GetExercises() {
	numberText01, textNumber01 := exercises.TestFunction("500")
	fmt.Println(numberText01)
	fmt.Println(textNumber01 + "\n")

	numberText02, textNumber02 := exercises.TestFunction("90")
	fmt.Println(numberText02)
	fmt.Println(textNumber02 + "\n")
}

func GetOS() {
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
