// https://leetcode.com/problems/print-foobar-alternately
// Non-channel solution
package main

import (
	"sync"
)

type FooBar struct {
	n    int
	MU   sync.Mutex
	b    bool
	cond *sync.Cond
}

func NewFooBar(n int) *FooBar {
	fb := FooBar{n: n}
	fbcond := sync.NewCond(&fb.MU)
	fb.cond = fbcond
	return &fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		fb.MU.Lock()
		for fb.b {
			fb.cond.Wait()
		}
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.b = !fb.b
		fb.MU.Unlock()
		fb.cond.Signal()
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		fb.MU.Lock()
		for !fb.b {
			fb.cond.Wait()
		}
		// printBar() outputs "bar". Do not change or remove this line.
		printBar()
		fb.b = !fb.b
		fb.MU.Unlock()
		fb.cond.Signal()
	}
}

// func printFoo() {
// 	fmt.Print("foo")
// }

// func printBar() {
// 	fmt.Print("bar")
// }

// func main() {
// 	fb := NewFooBar(5)
// 	go fb.Foo(printFoo)
// 	go fb.Bar(printBar)
// 	time.Sleep(time.Second * 1)
// }
