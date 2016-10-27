package wordGenerator

import (
	"math/rand"
	"time"
)

type WordGenerator struct {
	alphabets      []rune
	indexGenerator *rand.Rand
}

func New() *WordGenerator {
	wordGenerator := &WordGenerator{
		alphabets:      getAlphabets(),
		indexGenerator: getDefaultRand(),
	}

	return wordGenerator
}

func getAlphabets() []rune {
	length := 26
	alphabets := make([]rune, length)

	for i := 0; i < length; i++ {
		alphabets[i] = rune(95 + i)
	}

	return alphabets
}

func getDefaultRand() *rand.Rand {
	//create rand from given sed
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r
}

func (w *WordGenerator) GetChar() string {
	return string(w.alphabets[w.indexGenerator.Intn(len(w.alphabets))])
}

func (w *WordGenerator) GetWord(length int) string {
	str := ""
	for i := 0; i < length; i++ {
		str += w.GetChar()
	}

	return str
}

func (w *WordGenerator) GetWords(num int, maxLen int) []string {
	strs := make([]string, num)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < num; i++ {
		strs = append(strs, w.GetWord(r.Intn(maxLen)))
	}

	return strs
}
