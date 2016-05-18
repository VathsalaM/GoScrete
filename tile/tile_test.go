package src

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/VathsalaM/GoScrete/components"
	"github.com/VathsalaM/GoScrete/Interface"
)

func TestForTileColour(t *testing.T) {
	newTile:=NewTile(1,"black")
	assert.Equal(t,"black",newTile.color)
}

func TestCanPlace(t *testing.T) {
	newSoldier := components.NewSoldier(1,"black")
	newHorse := components.NewHorse(3,"white")
	newTile:=NewTile(1,"black")
	assert.Equal(t,Interface.Position{},newSoldier.Position())
	newTile.Place(newSoldier)
	//assert.Equal(t,Interface.Position{},newSoldier.Positon())
	newTile.Place(newHorse)
	//assert.Equal(t,Interface.Position{},newSoldier.Positon())
}