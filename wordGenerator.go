package wordGenerator

import (
	"math/rand"
	"time"
)

type WordGenerator struct {
	alphabets      []rune
	indexGenerator *rand.Rand
}

var wordGenerator *WordGenerator

func init() {
	wordGenerator = &WordGenerator{
		alphabets:      getAlphabets(),
		indexGenerator: getDefaultRand(),
	}

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

//GetChar will return a single character from Alphabets
func GetChar() string {
	return string(wordGenerator.alphabets[wordGenerator.indexGenerator.Intn(len(wordGenerator.alphabets))])
}

//GetWord will return a word whose length is defined by parameter
func GetWord(length int) string {
	str := ""
	for i := 0; i < length; i++ {
		str += GetChar()
	}

	return str
}

//GetWords will return a slice of string.
func GetWords(num int, maxLen int) []string {
	strs := make([]string, num)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < num; i++ {
		strs[i] = GetWord(r.Intn(maxLen) + 1)
	}

	return strs
}
