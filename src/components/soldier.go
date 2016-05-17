package components

import "github.com/VathsalaM/GoScrete/src"

type soldier src.Element{
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
