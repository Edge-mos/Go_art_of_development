package main

import "fmt"

func main() {
	fmt.Println("Range and Map")

	arr := []string{"a", "b", "c"}
	fmt.Printf("Initial arr: %v \n", arr)

	for i, val := range arr {
		fmt.Println(i, val)
	}

	fmt.Println("\nMap----------------------")
	pointsMap := map[string]Point{"one": {X: 11, Y: 22}}
	pointsMapV2 := make(map[int]Point)
	var pointsMapV3 = map[string]Point{"zero": {0, 0}} // !!!

	pointsMap["two"] = Point{X: 1, Y: 2}
	pointsMapV2[1] = Point{1, 2}
	fmt.Println(pointsMap, pointsMapV2, pointsMapV3)

	fmt.Println("\nMap values------------------")
	fmt.Println(pointsMap["one"], pointsMapV2[1])

	fmt.Println("\nMap search by key-----------")

	var keyToSearch = "one"
	val, exs := pointsMap[keyToSearch]
	if exs {
		fmt.Printf("Key: %v exists in map with value:%v \n", keyToSearch, val)
	}

	fmt.Println("Map iterate--------------------")

	for key, val := range pointsMap {
		fmt.Printf("Key: %v Val: %v \n", key, val)
	}

}

type Point struct {
	X int
	Y int
}
