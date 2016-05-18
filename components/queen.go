package components

import (
	"github.com/VathsalaM/GoSecret/Interface"
)

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

func (q *queen)UpdatePosition(position Interface.Position){
	q.position = position
}

func (q *queen)Position()Interface.Position{
	return q.position
}



