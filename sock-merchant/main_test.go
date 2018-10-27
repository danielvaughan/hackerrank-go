package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSockMerchant(t *testing.T) {
	assert.Equal(t, int32(3), sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}))
	assert.Equal(t, int32(4), sockMerchant(10, []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}))
}
