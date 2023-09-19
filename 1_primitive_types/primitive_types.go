package main

import (
	"fmt"
	"os"
)

const GLOBAL = "Global_const"

var GlobalInt = 42

func main() {
	fmt.Println("Hello World")

	//go_types
	var a int8 = 127             // -128 <> 127 									1 byte
	var b int16 = 32_767         // -32768 <> 32767 								2 bytes
	var c int32 = 2_147_483_647  // -2_147_483_648 <> 2_147_483_647				4 bytes
	var d int64 = 99999999999999 //												8 bytes

	var e uint8 = 255            // 0 <> 255										1 byte
	var f uint16 = 65_535        // 0 <> 65_535									2 bytes
	var g uint32 = 4_294_967_295 // 0 <> 4_294_967_295            				4 bytes
	var h uint64 = 9999999999999 // 												8 bytes

	var i byte = 255           // 0 <> 255 --> uint8							1 byte
	var j rune = 2_147_483_647 // -2_147_483_648 <> 2_147_483_647 --> int32   	4 bytes

	var k int = -1            // в зависимости от платформы int32 или int64
	var m uint = 1            // в зависимости от платформы uint32 или uint64
	var n float32 = 1.8       // положительные и отрицательные					4 bytes
	var o float64 = 1000.12   // положительные и отрицательные					8 bytes
	var p complex64 = 1 + 2i  //
	var q complex64 = 4 + 90i //

	var isTrue bool = true
	var name string = "Edge"

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, m, n, o, p, q, isTrue, name, GlobalInt)
	//-----------------------------------------------------------------------
	fmt.Println("PART two ------")
	var digit = 7
	var unknownDigit int
	var emptyName string
	fmt.Println(digit, unknownDigit, emptyName)
	fmt.Printf("This for emptyName: %s \n", emptyName)
	//конкатенация-----------------------------------------------------------

	var user = "Edge_mos"
	var age uint8 = 38

	var res = fmt.Sprintf("My name is: %s. And i'm %d ears old", user, age)

	fmt.Println(res)

	fmt.Println("-------------with incoming args------------")

	args := os.Args
	fmt.Printf("Type of args: %T, Values: %v \n", args, args)
	for index, arg := range args {
		fmt.Println(index, "-->", arg)
	}

	fmt.Printf("User custom args is: %v \n", args[1:])

}
