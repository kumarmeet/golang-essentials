package mylib

import (
	"errors"
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Concat(s1 string, s2 string) string {
	return fmt.Sprintf("%s %s", s1, s2)
}

/*
func getCoords() (int, int){
  var x int
  var y int
  return x, y
}

same as

func GetPoint() (x int, y int) {
	return x, y
}

*/

func GetPoint(a int) (x int, y int) {
	x = 10
	y = 20

	// temp := x + y

	if a > x {
		return x + 2, y + 4
	} else if a > y {
		return x + 4, y + 44
	} else {
		return x + 40, y + 48
	}

}

func Calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero")
	}

	mul = a * b
	div = a / b
	return mul, div, nil
}

type Car struct {
	Make   string
	Model  string
	Height int
	Width  int
	//nest anonymous structs as fields within other structs
	Price struct {
		Showroom float64
		OnRoad   float64
	}
	BackWheel  Wheel
	FrontWheel Wheel
}

type Wheel struct {
	Radius   float64
	Material string
}

func GetCarValues() (c Car) {
	c = Car{
		Make:   "Z-10",
		Model:  "Swift",
		Height: 41,
		Width:  66,
		FrontWheel: Wheel{
			Radius:   33.6,
			Material: "Rubber",
		},
		BackWheel: Wheel{
			Radius:   33.6,
			Material: "Rubber",
		},
		Price: struct {
			Showroom float64
			OnRoad   float64
		}{
			Showroom: 8599.44,
			OnRoad:   9999.69,
		},
	}

	return c
}
