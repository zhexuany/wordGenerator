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

func GetChar() string {
	return string(wordGenerator.alphabets[wordGenerator.indexGenerator.Intn(len(wordGenerator.alphabets))])
}

func GetWord(length int) string {
	str := ""
	for i := 0; i < length; i++ {
		str += GetChar()
	}

	return str
}

func GetWords(num int, maxLen int) []string {
	strs := make([]string, num)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < num; i++ {
		strs = append(strs, GetWord(r.Intn(maxLen)))
	}

	return strs
}
