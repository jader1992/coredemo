package demo

const KEY = "jade:demo"

type Service interface {
	GetFoo() Foo
}

type Foo struct {
	Name string
}