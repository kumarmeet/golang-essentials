package main

import (
	"fmt"

	"example.com/myproject/mylib"
)

func main() {
	result1 := mylib.Add(2, 3)
	result2 := mylib.Subtract(5, 3)
	fmt.Println(result1, result2)
	fullName := mylib.Concat("Meet", "Kumar")

	fmt.Println(fullName)

	ignore, _ := mylib.GetPoint(20)

	fmt.Println(ignore)

	_, _, c := mylib.Calculator(10, 0)

	fmt.Printf("%T\n", c)

	st := mylib.Car{}

	fmt.Println(mylib.GetCarValues(), st)

	//anonymous struct
	myId := struct {
		id   int
		name string
	}{
		id:   44,
		name: "Meet Kumar",
	}

	fmt.Println(myId)
}
