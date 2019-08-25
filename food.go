package main

import (
	"math/rand"
)

type Food struct {
	Position *Vector
	Emoji    string
	Size     int
}

func NewFood(position *Vector) *Food {
	food := []string{"🍌", "🍔", "🌮", "🌯"}
	emoji := food[rand.Intn(len(food))]
	return &Food{Position: position, Emoji: emoji, Size: 40}
}
