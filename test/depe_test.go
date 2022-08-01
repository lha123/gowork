package test

import (
	"fmt"
	"testing"
)

func TestDept(t *testing.T) {

	show[string]("sdf")

}

type UserId interface {
	int | string
}

type student[T UserId] struct {
	Name string
	str  int
}

type student1[T any] struct {
	Name string
	str  int
}

func (b *student[T]) Show(a T) {
	fmt.Println(a)
}
func (b *student1[T]) Show(a T) {
	fmt.Println("a", a)
}

type studentService[T any] interface {
	Show(a T)
}

func show[T any](a T) {
	fmt.Println(a)
}
