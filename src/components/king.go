package components

type king struct{
	position int
	id int
	colour string
}

func NewKing(id int,colour string)(newKing king){
	newKing.colour = colour
	newKing.id = id
	newKing.position= nil
	return
}


