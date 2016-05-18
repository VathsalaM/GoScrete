package components

import "github.com/VathsalaM/GoScrete/components"

type soldier struct{
	position int
	id int
	colour string
}

func NewSoldier(id int,colour string)(newSoldier soldier){
	newSoldier.colour = colour
	newSoldier.id = id
	newSoldier.position= nil
	return
}

func (s *soldier)UpdatePosition(element Element){

}

