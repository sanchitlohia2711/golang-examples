package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	//Both start and end
	num1 := numbers[2:4]
	fmt.Println("Both start and end")
	fmt.Printf("num1=%v\n", num1)
	fmt.Printf("length=%d\n", len(num1))
	fmt.Printf("capacity=%d\n", cap(num1))

	//Only start
	num2 := numbers[2:]
	fmt.Println("\nOnly start")
	fmt.Printf("num1=%v\n", num2)
	fmt.Printf("length=%d\n", len(num2))
	fmt.Printf("capacity=%d\n", cap(num2))

	//Only end
	num3 := numbers[:3]
	fmt.Println("\nOnly end")
	fmt.Printf("num1=%v\n", num3)
	fmt.Printf("length=%d\n", len(num3))
	fmt.Printf("capacity=%d\n", cap(num3))

	//None
	num4 := numbers[:]
	fmt.Println("\nOnly end")
	fmt.Printf("num1=%v\n", num4)
	fmt.Printf("length=%d\n", len(num4))
	fmt.Printf("capacity=%d\n", cap(num4))

	numbers[3] = 8
	fmt.Printf("num1=%v\n", num2)
	fmt.Printf("num3=%v\n", num3)
	fmt.Printf("num4=%v\n", num4)
}
