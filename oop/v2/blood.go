package oop

type A struct {
	name string
}

type B struct {
	name string
}

type AB struct {
	// 多重嵌套，挡多重嵌套的结构体有同名结构体时，就需要使用指明具体访问的类型
	A
	B
	name string
}
