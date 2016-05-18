package src

import t "github.com/VathsalaM/GoSecret/Tile"

type board struct {
	id    int
	table map[int]map[int]t.Tile
}

type Position struct{
	column int
	row int
}

type Element interface {
	currentPosition()(Position)
	NextPosition()(Position)
	Die()(bool)
	Kill()(bool)
}

func createTable(length int,width int ) (map[int]map[int]tile) {
	colours := []string{ "white","black"}
	columns := make(map[int](map[int]tile),0)
	for j := 0; j < length; j++ {
		var row = make(map[int]tile,0)
		for i :=0; i < width; i++ {
			row[i] = NewTile((j*width)+i, colours[i % 2])
		}
		columns[j] = row
	}
	return columns
}

func New(id int) (newboard board) {
	newboard.id = id
	newboard.table = createTable(8,8)
	return
}