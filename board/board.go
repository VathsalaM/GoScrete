package src

import (
	"github.com/VathsalaM/GoSecret/Interface"
	t "github.com/VathsalaM/GoSecret/tile"
)

type board struct {
	id    int
	table map[int]map[int]t.Tile
}

func createTable(length int, width int) map[int]map[int]t.Tile {
	colours := []string{"white", "black"}
	columns := make(map[int](map[int]t.Tile), 0)
	for j := 0; j < length; j++ {
		var row = make(map[int]t.Tile, 0)
		for i := 0; i < width; i++ {
			row[i] = t.Tile{Id: Interface.Position{j, i}, Color: colours[i%2]}
		}
		columns[j] = row
	}
	return columns
}

func New(id int) (newboard board) {
	newboard.id = id
	newboard.table = createTable(8, 8)
	return
}

func (b *board) Place(element Interface.Element, position Interface.Position) {
	tile := b.table[position.Row][position.Column]
	tile.Place(element)
}

func (b *board) MoveElementForward(element Interface.Element){
	currentPosition := element.Position()
	b.Place(element,Interface.Position{currentPosition.Row,currentPosition.Column+1})
}
