package src

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestForTileColour(t *testing.T) {
	newTile:=NewTile(1,"black")
	assert.Equal(t,"black",newTile.color)
}