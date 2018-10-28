package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingValleys(t *testing.T) {
	assert.Equal(t, int32(1), countingValleys(8, "UDDDUDUU"))
	assert.Equal(t, int32(2), countingValleys(12, "DDUUDDUDUUUD"))
}
