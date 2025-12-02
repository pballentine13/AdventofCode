package main

import (
	"testing"
)

func TestFindCode(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		expected  int
		expected2 int
	}{
		{"Test Input", "input_test.txt", 3, 6},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Dial := dial{CurrentValue: 50, MaxValue: 100, CountZero: 0, PassedZero: 0}
			Dial.FindCode(test.filename)
			if Dial.CountZero != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, Dial.CountZero)
			}
			if Dial.PassedZero != test.expected2 {
				t.Errorf("Expected %d, got %d", test.expected2, Dial.PassedZero)
			}
		})
	}
}

// func TestSpinLeft(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		start    int
// 		number   int
// 		expected int
// 	}{
// 		{"Left from 0", 0, 10, 90},
// 		{"Left Across 0", 5, 10, 95},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			Dial := dial{CurrentValue: test.start, MaxValue: 100}
// 			Dial.spinLeft(test.number)
// 			if Dial.CurrentValue != test.expected {
// 				t.Errorf("Expected %d, got %d", test.expected, Dial.CurrentValue)
// 			}
// 		})
// 	}
// }
//
// func TestSpinRight(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		start    int
// 		number   int
// 		expected int
// 	}{
// 		{"Right from 0", 0, 10, 10},
// 		{"Right Across 0", 95, 10, 5},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			Dial := dial{CurrentValue: test.start, MaxValue: 100}
// 			Dial.spinRight(test.number)
// 			if Dial.CurrentValue != test.expected {
// 				t.Errorf("Expected %d, got %d", test.expected, Dial.CurrentValue)
// 			}
// 		})
// 	}
// }
