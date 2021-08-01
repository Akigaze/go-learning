package oop

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) SetScore(score int) {
	s.Score = score
}

func (s *Student) Testing() {
	fmt.Printf("%s is testing\n", s.Name)
}

// PrimaryStudent GO使用嵌套的方式实现结构体的继承，只需要写类型不用指定属性名，同时这个类型名称也可以相当于属性名使用
type PrimaryStudent struct {
	Student
}

func (s *PrimaryStudent) String() string {
	return fmt.Sprintf("Primary Student %s, age: %d, score: %d", s.Name, s.Age, s.Score)
}

type University struct {
	Student
}

func (u *University) String() string {
	return fmt.Sprintf("University Student %s, age: %d, score: %d", u.Name, u.Age, u.Score)
}
