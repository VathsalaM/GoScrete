package src

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/VathsalaM/GoSecret/Interface"
	c "github.com/VathsalaM/GoSecret/components"
)

func TestBoardCreation(t *testing.T) {
	board := New(1)
	assert.Equal(t, 8, len(board.table))
}

func TestArrangedTiles(t *testing.T) {
	board := New(1)
	assert.Equal(t,"white",board.table[0][0].Color)
	assert.Equal(t,"black",board.table[0][1].Color)
	assert.Equal(t,"white",board.table[7][6].Color)
	assert.Equal(t,"black",board.table[7][7].Color)
	assert.Equal(t,Interface.Position{Row:7, Column:7},board.table[7][7].Id)
}

func TestMoveForward(t *testing.T) {
	board := New(1)
	s := c.NewSoldier(1,"black")
	assert.Equal(t,Interface.Position{},s.Position())
	board.Place(&s,Interface.Position{1,1})
	assert.Equal(t,Interface.Position{1,1},s.Position())
	board.MoveElementForward(&s)
	assert.Equal(t,Interface.Position{1,2},s.Position())
}
