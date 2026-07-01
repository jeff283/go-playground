package main

import "fmt"

func main() {
	fmt.Println("\nARRAYS")
	arraysAround()

	fmt.Println("\nSLICES")
	sliceAround()

	fmt.Println("\nMAPS")
	mapAround()

	fmt.Println("\nLOOPS")
	goBrrr()

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

	for i, v := range infArr {
		fmt.Printf("Index %v and Value %v \n", i, v)
	}
	fmt.Println()

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

func mapAround() {

	var basicMap map[string]uint8 = make(map[string]uint8, 10) //or map[string]uint8{}
	fmt.Println(basicMap)

	var initMap map[string]uint8 = map[string]uint8{"Jan": 1, "Feb": 2, "Mar": 3}
	fmt.Println(initMap)
	fmt.Println(initMap["Mar"])

	// Maps always return even if the key does not exist
	// It will return the defualt value of the data type for example in this case `unit8` default is 0
	fmt.Println(initMap["Apr"])
	fmt.Println()

	// Handling emptys
	month, ok := initMap["Jan"]
	if ok {
		fmt.Printf("Month index is %v and OK\n", month)
	} else {
		fmt.Printf("Apr is %v and is okay? %v\n", month, ok)
	}
	fmt.Println()

	// Deleting
	delete(initMap, "Jan")
	fmt.Println(initMap)
	fmt.Println()

	// Looping
	for month, index := range initMap {
		fmt.Printf("%v is month number %v \n", month, index)
	}
	fmt.Println()

}

func goBrrr() {
	i := 0
	fmt.Println("For Plain: ")
	for i < 3 {
		fmt.Printf("%v, ", i)
		i = i + 1
	}
	fmt.Println()

	fmt.Printf("i is global scoped: %v \n", i)

	fmt.Println("For Trad: ")
	for x := 0; x < 3; x++ {
		fmt.Printf("%v, ", x)
	}
	fmt.Println()

	fmt.Println("For Range: ")
	for x := range 3 {
		fmt.Printf("%v, ", x)
	}
	fmt.Println()

	fmt.Println("For Break: ")
	for {
		if i >= 6 {
			break
		}
		fmt.Printf("%v, ", i)
		i++
		// i--, i += 5 or -, 1 *= 2 or /

	}

	fmt.Println()
}
