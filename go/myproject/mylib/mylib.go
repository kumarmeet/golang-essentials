package mylib

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
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

// EMBEDDED STRUCTS
type User struct {
	Name string
	Age  int
	Address
}

type Address struct {
	CountyCode  string
	PhoneNumber string
	County      string
	State       string
	City        string
	Postal      string
	HouseNo     string
}

func UserDetails() (u User) {
	// address := Address{
	// 	CountyCode:  "+91",
	// 	PhoneNumber: "8888888888",
	// 	County:      "India",
	// 	State:       "MP",
	// 	City:        "BPL",
	// 	Postal:      "462323",
	// 	HouseNo:     "55",
	// }

	u = User{
		Name: "Ronie",
		Age:  64,
		// Address: address,
		Address: Address{
			CountyCode:  "+91",
			PhoneNumber: "8888888888",
			County:      "India",
			State:       "MP",
			City:        "BPL",
			Postal:      "462323",
			HouseNo:     "55",
		},
	}

	return u
}

func Welcome() (name string) {

	fmt.Println("Enter your name")

	t, _ := fmt.Scan(&name)

	fmt.Println(t)

	return name
}

func NewScanner() {

	fmt.Println("Enter lines of text (Ctrl+D to exit):")

	scanner := bufio.NewScanner(os.Stdin)

	// Read lines until Ctrl+D (EOF) is encountered
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("You entered:", line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}

func NewReader() {

	fmt.Print("Enter a sentence: ")

	reader := bufio.NewReader(os.Stdin)

	// Read a line (up to and including '\n')
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("You entered:", line)
}

// While Go is not object-oriented, it does support methods that can be defined on structs. Methods are just functions that have a receiver.
// A receiver is a special parameter that syntactically goes before the name of the function.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// interface
// Interfaces are collections of method signatures. A type "implements" an interface if it has all of the methods of the given interface defined on it.
// When a type implements an interface, it can then be used as the interface type.

// Interfaces are implemented implicitly.
// A type never declares that it implements a given interface. If an interface exists and a type has the proper methods defined, then the type automatically fulfills that interface.

/*
A type implements an interface by implementing its methods. Unlike in many other languages, there is no explicit declaration of intent, there is no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation. You may add methods to a type and in the process be unknowingly implementing various interfaces, and that's okay.
*/

// Remember, interfaces are collections of method signatures. A type "implements" an interface if it has all of the methods of the given interface defined on it.
type Shape interface {
	AgainArea() float64
	Perimeter() float64
}

type AgainRectangle struct {
	Width, Height float64
}

func (r AgainRectangle) AgainArea() float64 {
	return r.Width * r.Height
}

func (r AgainRectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

/*
If a type in your code implements an Perimeter method, with the same signature (e.g. accepts nothing and returns a float64), then that object is said to implement the shape interface.
*/
type Circle struct {
	Radius float64
}

func (c Circle) AgainArea() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeDetails(s Shape) {
	fmt.Printf("Area: %f, Perimeter: %f\n", s.AgainArea(), s.Perimeter())
}

/*
This is different from most other languages, where you have to explicitly assign an interface type to an object, like with Java:

class Circle implements Shape
*/
