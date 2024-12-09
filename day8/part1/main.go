package main

import (
	"fmt"
	"io"
	"log/slog"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Encountered err", "err", err)
	}
}

func run() error {
	f, err := os.Open("../input.txt")
	if err != nil {
		return err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	// sum := 0

	antennae := map[string][][2]int{}

	re, err := regexp.Compile(`[a-z]|[A-Z]|[0-9]`)
	if err != nil {
		return err
	}

	for i, s := range strings.Split(string(b), "\n") {
		for j := range s {
			char := string(s[j])
			if re.Match([]byte(char)) {
				v, ok := antennae[char]
				if ok {
					v = append(v, [2]int{i, j})
				} else {
					v = [][2]int{{i, j}}
				}
				antennae[char] = v
			}
		}
	}

	fmt.Printf("%v", antennae)

	return nil
}

func getResonantCoords(coords [][2]int) [][2]int {
	for _, v := range coords {

	}
}

func distanceBetweenPoints(s1, s2 [2]int) int {
	return math.Sqrt(math.Pow(float64(s1[0]-s2[0]), 2) + math.Pow(float64(s1[1]-s2[1]), 2))
}
