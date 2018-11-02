package main

import (
	"math/big"
	"syscall/js"
)

func fib(n int64) *big.Int {
	x, y, f := big.NewInt(0), big.NewInt(1), big.NewInt(0)

	var i int64
	for i = 1; i < n; i++ {
		f.Add(x, y)
		x.Set(y)
		y.Set(f)
	}
	return f
}

func main() {
	js.Global().Get("document").Call("getElementById", "title").Set("innerHTML", "Hello, WebAssembly!")

	var i int64
	for i = 0; i < 100; i++ {
		js.Global().Get("document").Call("getElementById", "result").Call("insertAdjacentHTML", "beforeend", fib(i).String() + "<br>")
	}
}
