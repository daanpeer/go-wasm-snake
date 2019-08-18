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

func collide(a Vector, b Vector) bool {
	return a.x == b.x && a.y == b.y
}

type Game struct {
	Snake  *Snake
	Food   *Food
	Width  float64
	Height float64
}

func NewGame(height float64, width float64) *Game {
	snake := NewSnake()
	game := &Game{Snake: snake, Width: width, Height: height}
	game.SpawnFood()
	return game
}

func (g *Game) SpawnFood() {
	x := randNumber(int(g.Width)/10) * 10
	y := randNumber(int(g.Height)/10) * 10
	g.Food = &Food{position: Vector{x: x, y: y}}
	fmt.Println("SpawnFood: ", g.Food)
}

func (g *Game) Loop() {
	g.Snake.Move()
	// check if hitting walls

	if collide(g.Food.position, g.Snake.Head()) {
		g.Snake.Eat()
		g.SpawnFood()
	}
}

func randNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
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
	g := NewGame(height, width)

	doc.Call("addEventListener", "keyup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		key := args[0].Get("key").String()
		switch k := key; k {
		case "w":
			g.Snake.Direction(Up)
		case "a":
			g.Snake.Direction(Left)
		case "s":
			g.Snake.Direction(Down)
		case "d":
			g.Snake.Direction(Right)
		}
		return nil
	}))

	var renderFrame js.Func

	var loop int
	renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		loop++

		if loop == 5 {
			loop = 0
			fmt.Println("loop")
			clear(ctxPtr, width, height)
			g.Loop()
			drawSnake(g.Snake, ctxPtr)
			drawFood(g.Food, ctxPtr)
		}

		js.Global().Call("requestAnimationFrame", renderFrame)
		return nil
	})
	defer renderFrame.Release()

	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done
}
