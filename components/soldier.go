package components

import "github.com/VathsalaM/GoSecret/Interface"

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

//func (s *soldier)UpdatePosition(position Interface.Position){

//}

func (s *soldier)Position()Interface.Position{
	return s.position
}

