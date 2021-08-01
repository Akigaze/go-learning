package main

import "fmt"

func main() {
	fmt.Println("---------- pointer of int ----------")
	pointerOfInt()
	fmt.Println("---------- pointer of array ----------")
	pointerOfArray()
	fmt.Println("---------- pointer of slice ----------")
	pointerOfSlice()
	fmt.Println("---------- pointer of struct ----------")
	pointerOfStruct()
}

type Student struct {
	Name string
	Id   int
}

func pointerOfStruct() {
	i := Student{"tom", 1}
	var p *Student = &i

	fmt.Println(p.Id)
}

func pointerOfSlice() {
	var i []int = []int{1, 3, 5}
	var p *[]int = &i

	fmt.Printf("address: &i = %p, p = %p, i = %p, &i[0] = %p \n", &i, p, i, &i[0])
	fmt.Printf("value: i = %d, *p = %d \n", i, *p)

	fmt.Println("modify i by p")
	*p = []int{2, 4}

	fmt.Printf("address: &i = %p, p = %p, i = %p \n", &i, p, i)
	fmt.Printf("value: i = %d, *p = %d \n", i, *p)

	fmt.Println("modify i by i")
	i = []int{5}

	fmt.Printf("address: &i = %p, p = %p, i = %p \n", &i, p, i)
	fmt.Printf("value: i = %d, *p = %d \n", i, *p)
}

func pointerOfArray() {
	var i [3]int = [3]int{1, 3, 5}
	var p *[3]int = &i

	fmt.Printf("address: &i = %p, p = %p \n", &i, p)
	fmt.Printf("value: i = %d, *p = %d \n", i, *p)

	fmt.Println("modify i by p")
	*p = [3]int{2, 4, 6}

	fmt.Printf("address: &i = %p, *p = %p \n", &i, p)
	fmt.Printf("value: i = %d, *p = %d \n", i, *p)

	fmt.Println("modify i by i")
	i = [3]int{0, 2, 4}

	fmt.Printf("address: &i = %p, *p = %p \n", &i, p)
	fmt.Printf("value: i = %d, *p = %d \n", i, *p)
}

func pointerOfInt() {
	var i int = 10
	var p *int = &i

	fmt.Printf("address: &i = %p, p = %p \n", &i, p)
	fmt.Printf("value: i = %d, *p = %d \n", i, *p)

	fmt.Println("modify i by p")
	*p = 20

	fmt.Printf("address: &i = %p, *p = %p \n", &i, p)
	fmt.Printf("value: i = %d, *p = %d \n", i, *p)

	fmt.Println("modify i by i")
	i = 30

	fmt.Printf("address: &i = %p, *p = %p \n", &i, p)
	fmt.Printf("value: i = %d, *p = %d \n", i, *p)
}
