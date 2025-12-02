package main

import (
	"testing"
)

func TestFixGiftShop(t *testing.T) {
	answer := fixGiftShop("input_test.txt")
	if answer != 1227775554 {
		t.Errorf("answer should be 1227775554, got %d", answer)
	}
}
