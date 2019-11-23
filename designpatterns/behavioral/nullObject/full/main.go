package main

import "fmt"

type college struct {
	departments []department
}

func (c *college) addDepartment(departmentName string, numOfProfessors int) {
	if departmentName == "computerscience" {
		computerScienceDepartment := &computerscience{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, computerScienceDepartment)
	}

	if departmentName == "mechanical" {
		mechanicalDepartment := &mechanical{numberOfProfessors: numOfProfessors}
		c.departments = append(c.departments, mechanicalDepartment)
	}
	return
}
func (c *college) getDepartment(departmentName string) department {
	for _, department := range c.departments {
		if department.getName() == departmentName {
			return department
		}
	}
	//Return a null department if the department doesn't exits
	return &nullDepartment{}

}

type department interface {
	getNumberOfProfessors() int
	getName() string
}

type computerscience struct {
	numberOfProfessors int
}

func (c *computerscience) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *computerscience) getName() string {
	return "computerscience"
}

type mechanical struct {
	numberOfProfessors int
}

func (c *mechanical) getNumberOfProfessors() int {
	return c.numberOfProfessors
}

func (c *mechanical) getName() string {
	return "mechanical"
}

type nullDepartment struct {
	numberOfProfessors int
}

func (c *nullDepartment) getNumberOfProfessors() int {
	return 0
}

func (c *nullDepartment) getName() string {
	return "nullDepartment"
}

func main() {

	college1 := createCollege1()
	college2 := createCollege2()

	totalProfessors := 0
	departmentArray := []string{"computerscience", "mechanical", "civil", "electronics"}
	for _, deparmentName := range departmentArray {
		d := college1.getDepartment(deparmentName)
		totalProfessors += d.getNumberOfProfessors()
	}

	fmt.Printf("Total number of professors in college1 is %d\n", totalProfessors)

	totalProfessors = 0
	for _, deparmentName := range departmentArray {
		d := college2.getDepartment(deparmentName)
		totalProfessors += d.getNumberOfProfessors()
	}

	fmt.Printf("Total number of professors in college2 is %d\n", totalProfessors)

}

func createCollege1() *college {
	college := &college{}
	college.addDepartment("computerscience", 4)
	college.addDepartment("mechanical", 5)
	return college
}

func createCollege2() *college {
	college := &college{}
	college.addDepartment("computerscience", 2)
	return college
}
