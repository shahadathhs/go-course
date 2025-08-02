package basic

import "fmt"

func DataTypes() {
	fmt.Println("Hello" + "World")
	fmt.Println("9 X 10 =", 9*10)
	fmt.Println("180.18/2.0 = ", 180.18/2.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("Hello", 1, true)

	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	map1 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(map1)

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	
}
