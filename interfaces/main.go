package main

import "fmt"

func main() {
	student := Student{firstName: "gilbert", lastName: "ndenzi", regNo: "123", studentId: "R2323"}
	professor := Professor{firstName: "barya", lastName: "mureeba", professorId: "P123", classNo: "C2323"}

	fmt.Println(getFullName(student))
	fmt.Println(getFullName(professor))

	var numbers = make([]int,3,5)
	fmt.Println(numbers)
}

type Person interface {
	createFullName() string
}

type Student struct {
	studentId, regNo, firstName, lastName string
}

type Professor struct {
	classNo, professorId, firstName, lastName string
}

func (s Student) createFullName() string {
	return s.firstName + " " + s.lastName
}

func (p Professor) createFullName() string {
	return "Prof. " + p.firstName + " " + p.lastName
}

func getFullName(p Person) string {
	return p.createFullName()
}
