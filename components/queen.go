package components

	type queen struct{
	position int
	id int
	colour string
}

func NewQueen(id int,colour string)(newQueen queen){
	newQueen.colour = colour
	newQueen.id = id
	newQueen.position= nil
	return
}


