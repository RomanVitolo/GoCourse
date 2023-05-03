package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var legend string
var err error

func GetNumbers() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Set a number: ")

	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Invalid data Type " + err.Error())
		}
	}

	fmt.Println("Set another number: ")

	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Invalid data Type " + err.Error())
		}
	}

	fmt.Println("Set a Legend: ")

	if scanner.Scan() {
		legend = scanner.Text()
	}

	fmt.Println(legend, number1*number2)
}
