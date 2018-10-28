package main

import (
	"testing"
)

func TestRansomNote1(t *testing.T) {
	checkMagazine([]string{"give", "me", "one", "grand", "today", "night"}, []string{"give", "one", "grand", "today"})
	// Output: Yes
}

func TestRansomNote2(t *testing.T) {
	checkMagazine([]string{"two", "times", "three", "is", "not", "four"}, []string{"two", "times", "two", "is", "four"})
	// Output: No
}
