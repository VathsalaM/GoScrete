package components

import "github.com/VathsalaM/GoScrete/Interface"

type elephant struct{
	position Interface.Position
	id int
	colour string
}

func NewElephant(id int,colour string)(newElephant elephant){
	newElephant.colour = colour
	newElephant.id = id
	newElephant.position= Interface.Position{}
	return
}


