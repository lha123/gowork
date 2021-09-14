package public

import (
	"testing"
)

type Optional struct {
	value interface{}
}

func OfNullable(data interface{}) (n *Optional) {
	if data == nil {
		n = empty()
	}
	return &Optional{data}
}

func empty() *Optional {
	t := new(Optional)
	return t
}

func (p *Optional) Get() interface{} {
	return p.value
}

func (p *Optional) IsPresent() bool {
	return p.value == nil
}

func (p *Optional) IfPresent(fn func(e interface{})) {
	if p.value != nil {
		fn(p.value)
	}
}

func (p *Optional) Filter(fn func(e interface{}) bool) *Optional {
	if fn == nil {
		panic("fn is nil")
	}
	if fn(p.value) {
		return p
	}
	return empty()
}

func (p *Optional) Map(fn func(e interface{}) interface{}) *Optional {
	if fn == nil {
		panic("fn is nil")
	}
	val := fn(p.value)
	return &Optional{val}
}

func (p *Optional) FlatMap(fn func(e interface{}) *Optional) *Optional {
	if fn == nil {
		panic("fn is nil")
	}
	val := fn(p.value)
	return val
}

func (p *Optional) OrElse(def interface{}) interface{} {
	if p.value != nil {
		switch p.value.(type) {
		case string, int, int8, int16, int32, float32, float64:
			if p.value == "" || p.value == 0 {
				return def
			}
			return p.value
		}
	}
	return def
}

func (p *Optional) OrElseGet(fn func() interface{}) interface{} {
	if p.value != nil {
		switch p.value.(type) {
		case string, int, int8, int16, int32, float32, float64:
			if p.value == "" || p.value == 0 {
				return fn()
			}
			return p.value
		}
	}
	return fn()
}

func TestOptional(t *testing.T) {

}
