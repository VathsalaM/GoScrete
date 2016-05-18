package components

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/VathsalaM/GoSecret/Interface"
)

func TestElephant_UpdatePosition(t *testing.T) {
	e := NewQueen(1,"black")
	assert.Equal(t,Interface.Position{}, e.Position())
	e.UpdatePosition(Interface.Position{2,3 })
	assert.Equal(t,Interface.Position{2,3}, e.Position())
}

