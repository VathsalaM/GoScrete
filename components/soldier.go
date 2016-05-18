package components

import (
	"github.com/VathsalaM/GoSecret/Interface"
	"fmt"
)

type soldier struct{
	position Interface.Position
	id int
	colour string
}

func NewSoldier(id int,colour string)(newSoldier soldier){
	newSoldier.colour = colour
	newSoldier.id = id
	newSoldier.position= Interface.Position{}
	return
}

func (s *soldier)UpdatePosition(position Interface.Position){
	fmt.Println("Before : ",s.position)
	s.position = position
	fmt.Println("After : ",s.position)
}

func (s *soldier)Position()Interface.Position{
	return s.position
}

