package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBag(t *testing.T) {
	bag := NewBag("light red bags contain 1 bright white bag, 2 muted yellow bags.")
	assert.NotNil(t, bag)
	assert.Equal(t, "light red", bag.Color)
	assert.Equal(t, 1, bag.Contains["bright white"])
	assert.Equal(t, 2, bag.Contains["muted yellow"])
	bag2 := NewBag("faded blue bags contain no other bags.")
	assert.NotNil(t, bag2)
	assert.Equal(t, "faded blue", bag2.Color)
	assert.Equal(t, 0, bag2.Contains["other"])
}
