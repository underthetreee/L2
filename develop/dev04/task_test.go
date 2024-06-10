package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	expectedResult := map[string][]string{
		"акптя":  {"пятак", "пятка", "тяпка"},
		"иклост": {"листок", "слиток", "столик"},
	}

	result := findAnagrams(words)

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("findAnagrams() = %v, want %v", result, expectedResult)
	}
}
