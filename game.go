package main

import (
	"fmt"
	"math"
)

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
	g.Food = NewFood(&Vector{x: x, y: y})
	fmt.Println("SpawnFood: ", g.Food)
}

func (g *Game) Loop() {
	if !g.Snake.IsDead {
		g.Snake.Move()
	}

	// check if hitting walls
	head := g.Snake.head
	if head.x < 0 || head.y < 0 || head.x > int(g.Width) || head.y > int(g.Height) {
		g.Snake.Dead()
	}

	dx := head.x - g.Food.position.x
	dy := head.y - g.Food.position.y

	distance := math.Sqrt(float64(dx*dx + dy*dy))

	if distance < float64(g.Snake.size/2+g.Food.size/2) {
		g.Snake.Eat()
		g.SpawnFood()
	}
}
