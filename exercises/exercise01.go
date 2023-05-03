package exercises

import (
	//"fmt"
	"strconv"
)

func TestFunction(convertText string) (int, string) {
	number, err := strconv.Atoi(convertText)
	if err != nil {
		return 0, "Get an Error " + err.Error()
	}
	if number > 100 {
		return number, "Its higher than 100"
	} else {
		return number, "Its lower than 100"
	}
}
