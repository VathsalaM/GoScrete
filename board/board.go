package src

import(
	t "github.com/VathsalaM/GoSecret/tile"
	"github.com/VathsalaM/GoSecret/Interface"
)

type board struct {
	id    int
	table map[int]map[int]t.Tile
}

type Element interface {
	currentPosition()(Interface.Position)
	NextPosition()(Interface.Position)
	Die()(bool)
	Kill()(bool)
}

func createTable(length int,width int ) (map[int]map[int]t.Tile) {
	colours := []string{ "white","black"}
	columns := make(map[int](map[int]t.Tile),0)
	for j := 0; j < length; j++ {
		var row = make(map[int]t.Tile,0)
		for i :=0; i < width; i++ {
			row[i] = t.Tile{Id:Interface.Position{},Colour:colours[i%2]}
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