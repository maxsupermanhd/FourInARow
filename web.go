//go:build js && wasm

package main

import (
	"fmt"
	"os"
	"syscall/js"
)

func jsoutput(str string) {
	js.Global().Get("document").Call("getElementById", "appoutput").Set("value", js.Global().Get("document").Call("getElementById", "appoutput").Get("value").String()+str)
	js.Global().Get("document").Call("getElementById", "appoutput").Set("scrollTop", js.Global().Get("document").Call("getElementById", "appoutput").Get("scrollHeight"))
}

func textinputevent(this js.Value, args []js.Value) interface{} {
	fmt.Printf("%#v", args[0].Get("key").String())
	if args[0].Get("key").String() == "Enter" {
		input <- js.Global().Get("document").Call("getElementById", "sendinput").Get("value").String()
		js.Global().Get("document").Call("getElementById", "sendinput").Set("value", "")
	}
	return nil
}

func shutdownjsinput(code int) {
	js.Global().Get("document").Call("getElementById", "sendinput").Set("disabled", true)
	os.Exit(code)
}

func init() {
	js.Global().Get("document").Call("getElementById", "sendinput").Call("addEventListener", "keydown", js.FuncOf(textinputevent))
	output = jsoutput
	exitfunc = shutdownjsinput
}
