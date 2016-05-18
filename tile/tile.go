package src

import (
	"github.com/VathsalaM/GoSecret/Interface"
	"fmt"
)

type Tile struct{
	Id Interface.Position
	Color string
	Element Interface.Element
}

func (t *Tile)Place(element Interface.Element){
	prevElement := t.Element
	if prevElement != nil{
		prevElement.UpdatePosition(Interface.Position{})
	}
	element.UpdatePosition(t.Id)
	fmt.Println(element)
	t.Element = element
}