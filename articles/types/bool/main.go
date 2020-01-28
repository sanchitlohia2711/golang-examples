package main

import "fmt"

func main() {
	//Default value will be false it not initialized
	var a bool
	fmt.Printf("a's value is %t\n", a)

	//And operation on one true and other false
	andOperation := 1 < 2 && 1 > 3
	fmt.Printf("Ouput of AND operation on one true and other false %t\n", andOperation)

	//OR operation on one true and other false
	orOperation := 1 < 2 || 1 > 3
	fmt.Printf("Ouput of OR operation on one true and other false: %t\n", orOperation)

	//Negation Operation on a false value
	negationOperation := !(1 > 2)
	fmt.Printf("Ouput of NEGATION operation on false value: %t\n", negationOperation)
}
