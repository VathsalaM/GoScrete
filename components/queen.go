package components

import "github.com/VathsalaM/GoScrete/Interface"

type queen struct{
	position Interface.Position
	id int
	colour string
}

func NewQueen(id int,colour string)(newQueen queen){
	newQueen.colour = colour
	newQueen.id = id
	newQueen.position= Interface.Position{}
	return
}


