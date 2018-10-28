package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeatedString(t *testing.T) {
	assert.Equal(t, int64(7), repeatedString("aba", 10))
	assert.Equal(t, int64(1000000000000), repeatedString("a", 1000000000000))
}
