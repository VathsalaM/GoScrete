package components

import (
	"github.com/VathsalaM/GoSecret/Interface"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoldier_UpdatePosition(t *testing.T) {
	s := NewSoldier(1, "black")
	assert.Equal(t, Interface.Position{}, s.Position())
	s.UpdatePosition(Interface.Position{2, 3})
	assert.Equal(t, Interface.Position{2, 3}, s.Position())
}

