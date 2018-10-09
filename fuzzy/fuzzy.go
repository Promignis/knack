package fuzzy

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/promignis/knack/utils"
)

type StrMetric struct {
	inputStr   string
	DictStr    string
	ratio      float64
	Levenstein int
}

type StrMetrics []StrMetric

func (s StrMetrics) Len() int { return len(s) }

func (s StrMetrics) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s StrMetrics) Less(i, j int) bool { return s[i].ratio <= s[j].ratio }

// TODO: make it concurrent with goroutines
func GetFuzzyCandidates(searchTerm string, dict []string, maxDistance int) StrMetrics {

	candidates := []StrMetric{}

	for _, word := range dict {
		distance := EditDistance(searchTerm, word)
		ratio := levensteinRatio(distance, searchTerm, word)
		if distance <= maxDistance {
			strMetric := StrMetric{searchTerm, word, ratio, distance}
			candidates = append(candidates, strMetric)
		}

	}

	sort.Sort(StrMetrics(candidates))

	return candidates
}

// TODO: load it at start
func GetDictWords() []string {

	var dictPath string

	if utils.IsUnixBased() {
		dictPath = "/usr/share/dict/words"
	} else {
		dictPath = "../data/words"
	}

	fileData, err := ioutil.ReadFile(dictPath)

	if err != nil {
		panic(err)
	}

	strFileData := string(fileData)

	return strings.Split(strFileData, "\n")

}

func runFuzzyStdin() {

	reader := bufio.NewReader(os.Stdin)

	words := GetDictWords()

	for {
		text, _ := reader.ReadString('\n')

		text = strings.Replace(text, "\n", "", -1)

		candidates := GetFuzzyCandidates(text, words, 2)
		viewCandidates(candidates, " ")
	}

}

func viewCandidates(candidates StrMetrics, delim string) {
	for _, c := range candidates {
		fmt.Printf("\033[32m %s%s\033[0m", c.DictStr, delim)
	}
	fmt.Print("\n\n")
}

func levensteinRatio(distance int, s1, s2 string) float64 {
	return float64(distance) / math.Max(float64(len(s1)), float64(len(s2)))
}
