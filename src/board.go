package src

type board struct {
	id    int
	table map[int]map[int]tile
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

func createTable(size int, length int) (map[int]map[int]tile) {
	colours := []string{ "white","black"}
	columns := make(map[int](map[int]tile),0)
	for j := 0; j < length; j++ {
		var row = make(map[int]tile,0)
		for i :=0; i < 8; i++ {
			row[i] = NewTile((j*8)+i, colours[i % 2])
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