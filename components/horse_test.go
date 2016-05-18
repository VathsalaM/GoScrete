package components

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/VathsalaM/GoSecret/Interface"
)

func TestHorse_UpdatePosition(t *testing.T) {
	h := NewQueen(1,"black")
	assert.Equal(t,Interface.Position{}, h.Position())
	h.UpdatePosition(Interface.Position{2,3 })
	assert.Equal(t,Interface.Position{2,3}, h.Position())
}

