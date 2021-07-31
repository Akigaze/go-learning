package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Buffer 是一个结构体
	buffer := bytes.Buffer{}

	Greet(&buffer, "Akigaze")

	got := buffer.String()
	wanted := "Hello, Akigaze"

	if got != wanted {
		t.Errorf("expect '%s' but not '%s'", wanted, got)
	}

}
