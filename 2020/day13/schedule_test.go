package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchedule(t *testing.T) {
	s := NewSchedule("7,13,x,x,59,x,31,19")
	assert.Equal(t, int64(1068781), s.findEarliestTime())
	s = NewSchedule("17,x,13,19")
	assert.Equal(t, int64(3417), s.findEarliestTime())
	s = NewSchedule("67,7,59,61")
	assert.Equal(t, int64(754018), s.findEarliestTime())
	// s := NewSchedule("7,13,x,x,59,x,31,19")
	// assert.Equal(t, int64(1068781), s.findEarliestTime())
}
