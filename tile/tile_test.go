package src

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/VathsalaM/GoSecret/components"
	"github.com/VathsalaM/GoSecret/Interface"
)

func TestCanPlace(t *testing.T) {
	newSoldier := components.NewSoldier(1,"black")
	newHorse := components.NewHorse(3,"white")
	newTile := Tile{Id:Interface.Position{1,1},Color:"Black"}
	assert.Equal(t,Interface.Position{},newSoldier.Position())
	newTile.Place(&newSoldier)
	assert.Equal(t,Interface.Position{1,1},newSoldier.Position())
	newTile.Place(&newHorse)
	assert.Equal(t,Interface.Position{},newSoldier.Position())
	assert.Equal(t,Interface.Position{1,1},newHorse.Position())
}