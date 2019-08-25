package main

import (
	"math/rand"
	"time"
)

func randNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func collide(a *Vector, b *Vector) bool {
	return a.x == b.x && a.y == b.y
}
