package src

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/VathsalaM/GoScrete/src/components"
)

func TestForTileColour(t *testing.T) {
	newTile:=NewTile(1,"black")
	assert.Equal(t,"black",newTile.color)
}

func TestCanPlace(t *testing.T) {
	newSoldier := components.NewSoldier(1,"black")
	newHorse := components.NewHorse(3,"white")
	newTile:=NewTile(1,"black")
	assert.Equal(t,Position{},newSoldier.Positon())
	newTile.Place(newSoldier)
	assert.Equal(t,Position{row:1,column:1},newSoldier.Positon())
	newTile.Place(newHorse)
	assert.Equal(t,Position{row:nil,column:nil},newSoldier.Positon())
}