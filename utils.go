package main

import (
	"math/rand"
	"time"
)

type Vector struct {
	x int
	y int
}

type Part struct {
	Position Vector
}

const (
	Up    = "Up"
	Down  = "Down"
	Left  = "Left"
	Right = "Right"
)

func randNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func collide(a *Vector, b *Vector) bool {
	return a.x == b.x && a.y == b.y
}
