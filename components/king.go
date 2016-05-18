package components

import (
	"github.com/VathsalaM/GoSecret/Interface"
)

type king struct{
	position Interface.Position
	id int
	colour string
}

func NewKing(id int,colour string)(newKing king){
	newKing.colour = colour
	newKing.id = id
	newKing.position= Interface.Position{}
	return
}

func (k *king)UpdatePosition(position Interface.Position){
	k.position = position
}

func (k *king)Position()Interface.Position{
	return k.position
}

