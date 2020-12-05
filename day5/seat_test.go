package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSeat(t *testing.T) {
	seat := NewSeat("FBFBBFFRLR")
	assert.NotNil(t, seat)
	assert.Equal(t, 5, seat.Column)
	assert.Equal(t, 44, seat.Row)
	assert.Equal(t, 357, seat.ID)
}
