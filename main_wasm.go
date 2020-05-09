// +bulid wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {

	fmt.Printf("main() function here\n")

	ch := make(chan bool, 10)

	fn := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Printf("got to callback function\n")
		ch <- true
		fmt.Printf("successfully pushed a bool onto the channel\n")
		return nil
	})
	js.Global().Get("window").Call("setCallback", fn)

	for v := range ch {
		fmt.Printf("got value from ch: %t\n", v)
	}

}
