package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
	assert.Equal(t, int32(4), jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
	assert.Equal(t, int32(3), jumpingOnClouds([]int32{0, 0, 0, 1, 0, 0}))
}
