package main

import "fmt"

func main() {
	fmt.Println("Functions and methods")

	//callbackPractice()	//1
	//closurePractice() //2
	structMethodsPractice() //3
	//calculatorApp()	//4
}

// ----------------------------------------------------------------
func closurePractice() {
	orderPrice := totalPrice(1)
	fmt.Println(orderPrice(1))
	fmt.Println(orderPrice(1))
	print("end")
}

func totalPrice(initPrice int) func(int) int {
	sum := initPrice
	return func(x int) int {
		sum += x
		return sum
	}
}

// ----------------------------------------------------------------
func callbackPractice() {

	str := "plus"
	var callback func(int, int) int

	if str == "plus" {
		callback = func(first int, second int) int {
			return first + second
		}
	}

	if str == "minus" {
		callback = func(first int, second int) int {
			return first - second
		}
	}

	fmt.Println(doSomething(callback, str, 5, 1))
}

//----------------------------------------------------------------

func doSomething(callback func(int, int) int, str string, first int, second int) int {
	result := callback(first, second)
	fmt.Println(str)
	return result
}

// -----------------------myApp------------------------------------
func calculatorApp() {
	fmt.Println("--------Calculator App!--------")
	firstDigit := 5
	secondDigit := 2
	operation := "+"

	//fmt.Println(operationsFactory())
	result := calculatorEngine(operation, firstDigit, secondDigit)
	fmt.Printf("Result of: %d %s %d = %d", firstDigit, operation, secondDigit, result)
}

func calculatorEngine(operation string, firstDigit int, secondDigit int) int {
	operationFunc, exists := operationsFactory()[operation]
	if exists {
		return operationFunc(firstDigit, secondDigit)
	}
	fmt.Printf("UNKNOWN operation!")
	return 0
}

func operationsFactory() map[string]func(int, int) int {
	operations := map[string]func(left, right int) int{}

	operations["+"] = func(left, right int) int { return left + right }
	operations["-"] = func(left, right int) int { return left - right }
	operations["*"] = func(left, right int) int { return left * right }
	operations["/"] = func(left, right int) int {
		if right != 0 {
			return left / right
		}
		return 0
	}
	operations["%"] = func(left, right int) int { return left % right }
	return operations
}

// Point -------------------------------struct_area-----------------------------------------
type Point struct {
	X int
	Y int
}

func movePoint(p Point, x int, y int) Point {
	p.X += x
	p.Y += y
	return p
}

func (p *Point) movePointByPointer(x, y int) {
	p.X += x
	p.Y += y
}

func structMethodsPractice() {
	point := Point{X: 1, Y: 2}
	fmt.Printf("Initial point: %v \n", point)

	point = movePoint(point, 1, 1)
	fmt.Printf("Point after move: %v \n", point)

	point.movePointByPointer(1, 1)
	fmt.Printf("Point after move by pointer: %v \n", point)

}
