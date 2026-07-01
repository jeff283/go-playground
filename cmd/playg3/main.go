package main

import "fmt"

func main() {
	fmt.Println("\nARRAYS")
	arraysAround()

	fmt.Println("\nSLICES")
	sliceAround()

	fmt.Println()
}

func arraysAround() {
	var intArr [3]int32 = [3]int32{0, 2, 3}
	intArr[0] = 1

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(intArr)

	infArr := [...]int32{4, 3, 2, 1}
	fmt.Println(infArr)

}

func sliceAround() {
	var basicSlice []int32 = []int32{6, 7, 8}

	fmt.Printf("The len of the slice is %v and the capacity is %v\n", len(basicSlice), cap(basicSlice))
	fmt.Println(basicSlice)

	basicSlice = append(basicSlice, 9)

	fmt.Printf("Append: The len of the slice is %v and the capacity is %v\n", len(basicSlice), cap(basicSlice))
	fmt.Println(basicSlice)

	var madeSlice []int32 = make([]int32, 3, 6)
	fmt.Printf("Made: The len of the slice is %v and the capacity is %v\n", len(madeSlice), cap(madeSlice))
	fmt.Println(madeSlice)
}
