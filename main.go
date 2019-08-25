package main

import (
	"strconv"
	"syscall/js"
)

// dead screen
// follow cursor
// fix hit yourself
// add points
// add fps

func main() {
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "canvas")
	width := canvasEl.Get("clientWidth").Float()
	height := canvasEl.Get("clientHeight").Float()
	canvasEl.Set("width", width)
	canvasEl.Set("height", height)
	ctx := canvasEl.Call("getContext", "2d")
	done := make(chan struct{}, 0)

	ctxPtr := &ctx
	g := NewGame(height, width)
	d := NewDrawing(ctxPtr, g)

	doc.Call("addEventListener", "keyup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		key := args[0].Get("key").String()
		switch k := key; k {
		case "w":
			g.Snake.SetDirection(Up)
		case "a":
			g.Snake.SetDirection(Left)
		case "s":
			g.Snake.SetDirection(Down)
		case "d":
			g.Snake.SetDirection(Right)
		case "p":
			g.Snake.Eat()
		}
		return nil
	}))

	points := doc.Call("getElementById", "points")

	var renderFrame js.Func
	var prevLength int
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		g.Loop()
		d.DrawGame()

		l := len(g.Snake.Parts)
		if l != prevLength {
			points.Set("innerHTML", strconv.Itoa(l))
			prevLength = l
		}

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release()

	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}
