package components

type elephant struct{
	position int
	id int
	colour string
}

func NewElephant(id int,colour string)(newElephant elephant){
	newElephant.colour = colour
	newElephant.id = id
	newElephant.position= nil
	return
}


