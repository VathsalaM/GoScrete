package src

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/VathsalaM/GoSecret/Interface"
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
	assert.Equal(t,Interface.Position{Row:0, Column:0},board.table[7][7].Id)
}