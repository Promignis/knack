package fuzzy

import (
	"math/rand"
	"testing"
)

//TODO: write tests to test accuracy of fuzzy search

func TestFuzzySearch(t *testing.T) {
	dict := []string{"abc", "aebcde", "dabcde", "a123b"}
	testWord := "ab"
	candidates := GetFuzzyCandidates(testWord, dict, 3)
	_ = candidates
	// viewCandidates(candidates, "\n")
}

func genRandString(size int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	byteArr := make([]byte, size)
	for i := range byteArr {
		byteArr[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(byteArr)
}

func benchmarkFuzzySearch(distance int, b *testing.B) {
	dict := GetDictWords()

	for n := 0; n < b.N; n++ {
		word := genRandString(rand.Intn(10))
		candidates := GetFuzzyCandidates(word, dict, distance)
		_ = candidates
	}
}

func BenchmarkFuzzySearch1(b *testing.B) { benchmarkFuzzySearch(1, b) }
func BenchmarkFuzzySearch2(b *testing.B) { benchmarkFuzzySearch(2, b) }
func BenchmarkFuzzySearch3(b *testing.B) { benchmarkFuzzySearch(3, b) }
