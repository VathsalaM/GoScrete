package components

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/VathsalaM/GoSecret/Interface"
)

func TestKing_UpdatePosition(t *testing.T) {
	k := NewQueen(1,"black")
	assert.Equal(t,Interface.Position{}, k.Position())
	k.UpdatePosition(Interface.Position{2,3 })
	assert.Equal(t,Interface.Position{2,3}, k.Position())
}

