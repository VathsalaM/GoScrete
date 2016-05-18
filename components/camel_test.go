package components

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/VathsalaM/GoSecret/Interface"
)

func TestCamel_UpdatePosition(t *testing.T) {
	q := NewQueen(1,"black")
	assert.Equal(t,Interface.Position{}, q.Position())
	q.UpdatePosition(Interface.Position{2,3 })
	assert.Equal(t,Interface.Position{2,3}, q.Position())
}

