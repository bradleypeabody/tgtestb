// +bulid wasm

package main

import (
	"syscall/js"
)

func main() {

	println("main() function here")

	ch := make(chan bool, 10)

	fn := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		println("got to callback function")
		ch <- true
		println("successfully pushed a bool onto the channel")
		return nil
	})
	js.Global().Get("window").Call("setCallback", fn)

	for v := range ch {
		println("got value from ch", v)
	}

}

// func main() {
// 	ch := make(chan bool)
// 	js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		go func() {
// 			println("sending")
// 			ch <- true
// 			println("sent")
// 		}()
// 		return nil
// 	}), 2000)
// 	println("waiting")
// 	<-ch
// 	println("done")
// }
