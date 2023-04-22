package main

import "fmt"

type Student struct {
	name   string
	rollno int
	marks  int
	grade  string
}

func (s Student) CalculateGrade() {
	if s.marks >= 90 {
		s.grade = "A+"
	} else if s.marks >= 80 {
		s.grade = "B+"
	} else if s.marks >= 70 {
		s.grade = "C+"
	} else if s.marks >= 60 {
		s.grade = "D+"
	} else if s.marks > 50 {
		s.grade = "E+"
	} else {
		s.grade = "fail"
	}
}
func main() {
	s := Student{"chinni", 11, 90, "A+"}
	s.CalculateGrade()
	fmt.Println("Grade is :", s.grade)
}
