package wordGenerator_test

import (
	"testing"

	"github.com/zhexuany/wordGenerator"
)

func TestWordGenerator_GetChar(t *testing.T) {
	if len(wordGenerator.GetChar()) != 1 {
		t.Errorf("Failed to test GetChar(). The length of its return shall be 1.")
	}
}

func TestWordGenerator_GetWord(t *testing.T) {
	if len(wordGenerator.GetWord(5)) != 5 {
		t.Errorf("Failed to test GetWord().")
	}
}

func TestWordGenerator_GetWords(t *testing.T) {
	strs := wordGenerator.GetWords(10, 20)
	if len(strs) != 10 {
		t.Errorf("Failed to test GetWords(length) %d", len(strs))
	}

	var max int
	for _, v := range strs {
		if max < len(v) {
			max = len(v)
		}
	}

	if max > 20 {
		t.Errorf("Failed to test GetWords(num, maxLen)")
	}

}
