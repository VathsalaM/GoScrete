package components

import "github.com/VathsalaM/GoScrete/Interface"

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


