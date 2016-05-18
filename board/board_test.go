package src

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBoardCreation(t *testing.T) {
	board := New(1)
	assert.Equal(t, 8, len(board.table))
}

func TestArrangedTiles(t *testing.T) {
	board := New(1)
	assert.Equal(t,"white",board.table[0][0].color)
	assert.Equal(t,"black",board.table[0][1].color)
	assert.Equal(t,"white",board.table[7][6].color)
	assert.Equal(t,"black",board.table[7][7].color)
	assert.Equal(t,63,board.table[7][7].id)
}