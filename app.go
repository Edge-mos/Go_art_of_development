package main

import "fmt"

func main() {
	fmt.Println("Slices")

	letters := []string{"a", "b", "c"} // without length of array = slice
	letters[0] = "1"

	letters = append(letters, "d") // can call func append on slice
	letters = append(letters, "e", "f")

	fmt.Println(letters)

	fmt.Println("Slice by make func-----------")
	createSlice := make([]string, 3) // better way

	fmt.Printf("Capasity before: %v \n", cap(createSlice))
	fmt.Printf("Len before: %v \n", len(createSlice))

	createSlice[0] = "1"
	createSlice[1] = "2"
	createSlice[2] = "3"
	//createSlice[3] = "4" panic!
	createSlice = append(createSlice, "4")

	fmt.Printf("Capasity after: %v \n", cap(createSlice))
	fmt.Printf("Len after: %v \n", len(createSlice))

	fmt.Println(createSlice)

	animalsArr := [4]string{
		"dog",
		"cat",
		"giraffe",
		"elephant"}

	var slc []string = animalsArr[0:2] //not include last value
	var slc2 []string = animalsArr[1:3]

	fmt.Println(animalsArr, slc, slc2)

	slc2[0] = "CAT!"
	fmt.Println(animalsArr, slc, slc2)

	fmt.Println("---------------------------")
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(digits, digits[:5], digits[5:], digits[0:], digits[:])

}
