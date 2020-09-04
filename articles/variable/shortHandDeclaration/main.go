package main

import "fmt"

func main() {
	m2 := 123                   //Type Inferred will be int
	n2 := "circle"              //Type Inferred will be string
	o2 := 5.6                   //Type Inferred will be float64
	p2 := true                  //Type Inferred will be bool
	q2 := 'a'                   //Type Inferred will be rune
	r2 := 3 + 5i                //Type Inferred will be complex128
	s2 := &sample{name: "test"} //Type Inferred will be pointer to main.sample
	t2 := sample{name: "test"}  //Type Inferred will be main.sample
	u2 := get()                 //Type inferred will []string as that is the return value of function get()

	fmt.Printf("m2: Type: %T Value: %v\n", m2, m2)
	fmt.Printf("n2: Type: %T Value: %v\n", n2, n2)
	fmt.Printf("o2: Type: %T Value: %v\n", o2, o2)
	fmt.Printf("p2: Type: %T Value: %v\n", p2, p2)
	fmt.Printf("q2: Type: %T Value: %v\n", q2, q2)
	fmt.Printf("r2: Type: %T Value: %v\n", r2, r2)
	fmt.Printf("s2: Type: %T Value: %v\n", s2, s2)
	fmt.Printf("t2: Type: %T Value: %v\n", t2, t2)
	fmt.Printf("u2: Type: %T Value: %v\n", u2, u2)

}

type sample struct {
	name string
}

func get() []string {
	return []string{"a", "b"}
}
