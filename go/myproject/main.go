package main

import (
	"fmt"
	"strconv"

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

	fmt.Println(mylib.UserDetails())

	i, err := strconv.Atoi("-8459")

	if err != nil {
		fmt.Println("Error")
	} else {
		fmt.Println(i)
	}

	mylib.NewReader()

	mylib.NewScanner()

	rect := mylib.Rectangle{
		Height: 12.1,
		Width:  44.54,
	}

	fmt.Println(rect.Area())

	againReact := mylib.AgainRectangle{
		Width:  14.22,
		Height: 84.97,
	}

	circle := mylib.Circle{
		Radius: 3.4,
	}

	mylib.PrintShapeDetails(againReact)
	mylib.PrintShapeDetails(circle)

}
