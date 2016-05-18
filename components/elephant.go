package components

import "github.com/VathsalaM/GoSecret/Interface"

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

func (e *elephant)UpdatePosition(position Interface.Position){
	e.position = position
}

func (e *elephant)Position()Interface.Position{
	return e.position
}
