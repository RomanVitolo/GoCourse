package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var State bool
var Salary float32
var Date time.Time

func RemainderVariables() {
	Name = "Roman"
	State = true
	Salary = 1400
	Date = time.Now()

	fmt.Println(Name)
	fmt.Println(State)
	fmt.Println(Salary)
	fmt.Println(Date)
}

func ConvertToText(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}
