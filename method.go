package main

import "fmt"

type Student struct {
	name string
}

//Define method
//func (t Type) method(params) return { //code}
func (s Student) getNameI() string {
	return s.name
}

//(t Type) => Receiver
//1.Value Receiver:Không làm thay đổi giá trị field trong struct
func (s Student) changeName() {
	s.name = "Mary"
}

//2.Pointer Receiver:Thay đổi field gốc trong struct
func (s *Student) changeName2() {
	fmt.Printf("p1 = %p", s)

	s.name = "Mary"
}

func main() {
	student := &Student{"kai"}

	// name := student.getNameI()

	// fmt.Println(name)

	// student.changeName()
	// fmt.Println(student.name)

	fmt.Printf("p1 = %p", student)
	fmt.Println()
	student.changeName2()
	fmt.Println(student.name)
}
