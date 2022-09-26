package main

import "fmt"

type Student struct {
	name string
}

//Define method
//func (t Type) method(params) return { //code}

//(t Type) => Receiver
//1.Value Receiver:Không làm thay đổi giá trị field trong struct
func (s Student) changeName() {
	s.name = "Mary"
}

//2.Poiter Receiver

func (s Student) getNameI() string {
	return s.name
}

func main() {
	student := Student{"kai"}

	name := student.getNameI()

	fmt.Println(name)

	student.changeName()
	fmt.Println(student.name)
}
