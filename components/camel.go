package components

import "github.com/VathsalaM/GoSecret/Interface"

type camel struct{
	position Interface.Position
	id int
	colour string
}

func NewCamel(id int,colour string)(newCamel camel){
	newCamel.colour = colour
	newCamel.id = id
	newCamel.position= Interface.Position{}
	return
}

func (c *camel)UpdatePosition(position Interface.Position){
	c.position = position
}

func (c *camel)Position()Interface.Position{
	return c.position
}

