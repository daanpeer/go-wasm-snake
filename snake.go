package main

import "fmt"

type Snake struct {
	Head      *Vector
	Direction string
	Size      int
	IsDead    bool
	Parts     []Part
	Path      []Vector
}

func NewSnake() *Snake {
	return &Snake{Size: 40, Direction: Left, Head: &Vector{90, 90}}
}

func (s *Snake) SetDirection(direction string) {
	s.Direction = direction
}

func (s *Snake) Eat() {
	fmt.Println("eat: eating")
	s.Parts = append(s.Parts, Part{})
}

func (s *Snake) moveParts() {
	var from int
	var part Part
	for r := 1; r < len(s.Parts)+1; r++ {
		from = r * 10
		part = s.Parts[r-1]
		if len(s.Path) > from {
			part.Position = s.Path[len(s.Path)-(from)]
			s.Parts[r-1] = part
		}
	}
}

func (s *Snake) Move() {
	switch dir := s.Direction; dir {
	case Up:
		s.Head.y -= 4
	case Right:
		s.Head.x += 4
	case Left:
		s.Head.x -= 4
	default:
	case Down:
		s.Head.y += 4
	}
	s.Path = append(s.Path, *s.Head)
	s.moveParts()
}

func (s *Snake) Dead() {
	fmt.Println("Died!")
	s.IsDead = true
}
