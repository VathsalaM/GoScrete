package components

type camel struct{
	position int
	id int
	colour string
}

func NewCamel(id int,colour string)(newCamel camel){
	newCamel.colour = colour
	newCamel.id = id
	newCamel.position= nil
	return
}


