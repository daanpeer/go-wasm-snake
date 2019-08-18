package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
	"time"
)

type Vector struct {
	x int
	y int
}

type Food struct {
	position Vector
}

type Snake struct {
	body      []Vector
	direction string
	size      int
	digesting bool
}

const (
	Up    = "Up"
	Down  = "Down"
	Left  = "Left"
	Right = "Right"
)

func NewSnake() *Snake {
	return &Snake{size: 10, direction: Left, body: []Vector{Vector{x: 40, y: 90}}}
}

func (s *Snake) Direction(direction string) {
	s.direction = direction
}

func (s *Snake) Head() Vector {
	return s.body[0]
}

func (s *Snake) Tail() Vector {
	return s.body[len(s.body)-1]
}

func (s *Snake) Eat() {
	fmt.Println("eat: eating")
	s.digesting = true
}

func (s *Snake) Move() {
	head := s.Head()
	newHead := Vector{head.x, head.y}
	// create new head based original head
	// move new head
	// put new head on top

	switch dir := s.direction; dir {
	case Up:
		newHead.y -= s.size
	case Right:
		newHead.x += s.size
	case Left:
		newHead.x -= s.size
	default:
	case Down:
		newHead.y += s.size
	}

	s.body = append([]Vector{newHead}, s.body...)

	if !s.digesting {
		s.body = s.body[:len(s.body)-1]
	} else {
		s.digesting = false
	}
}

func drawFood(f *Food, ctx *js.Value) {
	fmt.Println("Drawing food")
	ctx.Call("fillText", "🌮", f.position.x, f.position.y)
}

func drawSnake(s *Snake, ctx *js.Value) {
	fmt.Println("Drawing snake", s)
	for _, part := range s.body {
		ctx.Call("fillText", "🐛", part.x, part.y)
	}
}

func clear(ctx *js.Value, width float64, height float64) {
	fmt.Println(width, height)
	ctx.Call("clearRect", 0, 0, width, height)
}

func random(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func collide(a Vector, b Vector) bool {
	return a.x == b.x && a.y == b.y
}

func main() {
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "canvas")
	width := doc.Get("body").Get("clientWidth").Float()
	height := doc.Get("body").Get("clientHeight").Float()
	canvasEl.Set("width", width)
	canvasEl.Set("height", height)
	ctx := canvasEl.Call("getContext", "2d")
	done := make(chan struct{}, 0)

	ctxPtr := &ctx

	snake := NewSnake()
	food := &Food{position: Vector{x: 20, y: 90}}

	doc.Call("addEventListener", "keyup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		key := args[0].Get("key").String()
		switch k := key; k {
		case "w":
			snake.Direction(Up)
		case "a":
			snake.Direction(Left)
		case "s":
			snake.Direction(Down)
		case "d":
			snake.Direction(Right)
		}

		if key == "l" {
			fmt.Println("loop")
			clear(ctxPtr, width, height)
			drawSnake(snake, ctxPtr)
			drawFood(food, ctxPtr)
			if collide(food.position, snake.Head()) {
				snake.Eat()

			}
			snake.Move()
		}
		return nil
	}))

	fmt.Println(ctx, snake)

	var renderFrame js.Func

	var loop int
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		loop++

		if loop == 20 {
			loop = 0
			fmt.Println("loop")
			clear(ctxPtr, width, height)
			snake.Move()
			drawSnake(snake, ctxPtr)
			drawFood(food, ctxPtr)

			fmt.Println(food.position, snake.Head(), collide(food.position, snake.Head()))
			if collide(food.position, snake.Head()) {
				snake.Eat()
			}
		}

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release()

	js.Global().Call("requestAnimationFrame", renderFrame)

	fmt.Println("Hello, WebAssembly!")

	<-done
}
