package main

import (
	"math/rand"
)

type Food struct {
	position *Vector
	Emoji    string
	size     int
}

func NewFood(position *Vector) *Food {
	food := []string{"🍌", "🍔", "🌮", "🌯"}
	emoji := food[rand.Intn(len(food))]
	return &Food{position: position, Emoji: emoji, size: 40}
}
