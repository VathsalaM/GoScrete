package src


type tile struct{
	id int
	color string
}

func NewTile(id int,color string)(newTile tile){
	newTile.color = color
	newTile.id = id
	return
}