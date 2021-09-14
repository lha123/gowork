package test

import (
	"fmt"
	"testing"
)

type Person struct {
	name string
	age  int
}

func TestUser(t *testing.T) {
	var p Person
	fmt.Println(p)
}
