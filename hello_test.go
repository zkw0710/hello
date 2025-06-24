package hello_test

import (
	"fmt"
	"testing"

	"github.com/zkw0710/hello"
)

func TestHello(t *testing.T) {
	data := "jack"
	expected := fmt.Sprintf("hello %s!", data)
	result := hello.Hello(data)

	if result != expected {
		t.Fatalf("expected result %s, but got %s ", expected, result)
	}
}
