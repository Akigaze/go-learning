package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet 如果参数是接口类型，似乎不需要使用指针类型作为形参
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// GreeterHandler 使用http Response 作为输出
func GreeterHandler(writer http.ResponseWriter, request *http.Request) {
	Greet(writer, "world")
}

func main() {
	http.ListenAndServe(":5500", http.HandlerFunc(GreeterHandler))
}
