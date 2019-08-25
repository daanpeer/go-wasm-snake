package main

import (
	"strconv"
	"syscall/js"
)

type Drawing struct {
	ctx  *js.Value
	Game *Game
}

func NewDrawing(ctx *js.Value, g *Game) *Drawing {
	return &Drawing{ctx: ctx, Game: g}
}

func (d *Drawing) DrawFood() {
	food := d.Game.Food
	d.ctx.Set("font", strconv.Itoa(food.size)+"px Georgia")
	d.ctx.Call("fillText", food.Emoji, food.position.x, food.position.y)
}

func (d *Drawing) DrawSnake() {
	snake := d.Game.Snake
	d.ctx.Call("fillText", "🔴", snake.head.x, snake.head.y)
	for _, part := range snake.parts {
		d.ctx.Call("fillText", "🔵", part.position.x, part.position.y)
	}
}

func (d *Drawing) DrawPath() {
	for _, p := range d.Game.Snake.path {
		d.ctx.Call("fillText", "*", p.x, p.y)
	}
}

func (d *Drawing) Clear() {
	d.ctx.Call("clearRect", 0, 0, d.Game.Width, d.Game.Height)

}

func (d *Drawing) Background() {
	d.ctx.Set("fillStyle", "black")
	d.ctx.Call("fillRect", 0, 0, d.Game.Width, d.Game.Height)
}

func (d *Drawing) GameOver() {
	d.ctx.Set("font", "30px Arial")
	d.ctx.Set("fillStyle", "black")
	d.ctx.Call("fillRect", 0, 0, d.Game.Width, d.Game.Height)
	d.ctx.Set("fillStyle", "white")
	d.ctx.Call("fillText", "Dead 😞", d.Game.Width/2, d.Game.Height/2)
}

func (d *Drawing) DrawPos(v *Vector) {
	d.ctx.Call("fillText", "x: "+strconv.Itoa(v.x)+"y: "+strconv.Itoa(v.y), v.x, v.y-40)
}

func (d *Drawing) Debug() {
	d.ctx.Set("font", "20px Arial")
	d.ctx.Set("fillStyle", "white")

	head := d.Game.Snake.head
	d.DrawPos(&head)

	d.DrawPos(d.Game.Food.position)
}

func (d *Drawing) DrawGame() {
	d.Clear()

	if d.Game.Snake.IsDead {
		d.GameOver()
		return
	}

	d.Background()
	d.DrawFood()
	d.DrawSnake()
	d.Debug()
}
