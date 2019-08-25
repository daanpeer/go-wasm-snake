package main

import "fmt"

type Snake struct {
	head      Vector
	direction string
	size      int
	IsDead    bool
	parts     []Part
	path      []Vector
	steps     int
}

func NewSnake() *Snake {
	return &Snake{size: 40, direction: Left, head: Vector{90, 90}}
}

func (s *Snake) Direction(direction string) {
	s.direction = direction
}

func (s *Snake) Eat() {
	fmt.Println("eat: eating")
	s.parts = append(s.parts, Part{})
}

func (s *Snake) moveParts() {
	var from int
	var part Part
	for r := 1; r < len(s.parts)+1; r++ {
		from = r * 10
		part = s.parts[r-1]
		if len(s.path) > from {
			part.position = s.path[len(s.path)-(from)]
			s.parts[r-1] = part
		}
	}
}

func (s *Snake) Move() {
	switch dir := s.direction; dir {
	case Up:
		s.head.y -= 4
	case Right:
		s.head.x += 4
	case Left:
		s.head.x -= 4
	default:
	case Down:
		s.head.y += 4
	}
	s.path = append(s.path, s.head)
	s.moveParts()
}

func (s *Snake) Dead() {
	fmt.Println("Died!")
	s.IsDead = true
}
