package b

import aService "example/user/hello/a"

type B struct{}

func NewB() *B {
	return &B{}
}

var doSomething = aService.DoSomething

func (b B) Results() string {
	return doSomething()
}
