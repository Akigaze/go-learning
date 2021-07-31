package di

import (
	"fmt"
	"io"
)

// Greet 如果参数是接口类型，似乎不需要使用指针类型作为形参
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
