package oop

import (
	"testing"
)

func TestStudent(t *testing.T) {
	t.Run("Primary student", func(t *testing.T) {
		tom := PrimaryStudent{Student{"Tom", 10, 0}}
		tom.Testing()    // 等同于 tom.Student.Testing()
		tom.SetScore(80) // GO会先从调用的对象自身找调用的方法，如果不存在，就会找嵌套的结构体是否有同名方法，直到找到或者找不到抛异常
		got := tom.String()
		wanted := "Primary Student Tom, age: 10, score: 80"

		if got != wanted {
			t.Errorf("expect '%s' but got '%s'", wanted, got)
		}
	})

	t.Run("University student", func(t *testing.T) {
		tom := University{Student{"Jim", 22, 0}}
		tom.Testing()    // 等同于 tom.Student.Testing()
		tom.SetScore(70) // GO会先从调用的对象自身找调用的方法，如果不存在，就会找嵌套的结构体是否有同名方法，直到找到或者找不到抛异常
		got := tom.String()
		wanted := "University Student Jim, age: 22, score: 70"

		if got != wanted {
			t.Errorf("expect '%s' but got '%s'", wanted, got)
		}
	})

}
