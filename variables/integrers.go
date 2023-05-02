package variables

import "fmt"

func ShowIntegrers() {
	var commondInt int
	intOf32 := int32(10)
	intof64 := int64(100)

	fmt.Println("commondInt = ", commondInt)
	fmt.Println("intOf32 = ", intOf32)
	fmt.Println("intof64 = ", intof64)
}
