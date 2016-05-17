package components

type horse struct{
	position int
	id int
	colour string
}

func NewHorse(id int,colour string)(newHorse horse){
	newHorse.colour = colour
	newHorse.id = id
	newHorse.position= nil
	return
}

