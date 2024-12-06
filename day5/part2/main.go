package main

import (
	"bufio"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strconv"
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

	sum := 0

	after := map[string][]string{}
	before := map[string][]string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			vals := strings.Split(strings.TrimSpace(line), "|")

			rules := []string{}
			if v, ok := after[vals[0]]; ok {
				rules = v
			}

			rules = append(rules, vals[1])

			after[vals[0]] = rules

			rules = []string{}
			if v, ok := before[vals[1]]; ok {
				rules = v
			}

			rules = append(rules, vals[0])

			before[vals[1]] = rules
		} else if strings.Contains(line, ",") {
			vals := strings.Split(strings.TrimSpace(line), ",")

			if !valid(after, vals) {

				correctedVals := ncorrect(after, before, vals)
				if len(vals) != len(correctedVals) {
					panic(errors.New("incorrect length"))
				}

				v, err := middleNum(correctedVals)
				if err != nil {
					return err
				}

				sum += v
			}
		}
	}

	fmt.Println(sum)

	return nil
}

func valid(after map[string][]string, vals []string) bool {
	readVals := []string{}

	for i := range vals {
		for _, v := range after[vals[i]] {
			if slices.Contains(readVals, v) {
				return false
			}
		}

		readVals = append(readVals, vals[i])
	}

	return true
}

func correct(after map[string][]string, _ map[string][]string, vals []string) []string {
	correctedVals := []string{}

	for i := range vals {
		targetIndex := 0
		if i == 0 {
			correctedVals = append(correctedVals, vals[i])
			continue
		}

		for j := range correctedVals {
			for _, v := range after[vals[i]] {
				if slices.Contains(correctedVals[:j], v) {
					targetIndex = j - 1
				}
			}
		}
		correctedVals = slices.Insert(correctedVals, targetIndex, vals[i])
	}

	fmt.Printf("Incorrect: %v\n", vals)
	fmt.Printf("Corrected: %v\n\n", correctedVals)
	return correctedVals
}

func ncorrect(after map[string][]string, _ map[string][]string, vals []string) []string {
	correctedVals := make([]string, len(vals))
	copy(correctedVals, vals)

	// Keep swapping until no changes are needed
	changed := true
	for changed {
		changed = false
		for i := 0; i < len(correctedVals)-1; i++ {
			current := correctedVals[i]
			next := correctedVals[i+1]

			// Check if next must come before current
			mustSwap := false
			for _, v := range after[next] {
				if v == current {
					mustSwap = true
					break
				}
			}

			if mustSwap {
				correctedVals[i], correctedVals[i+1] = correctedVals[i+1], correctedVals[i]
				changed = true
			}
		}
	}

	fmt.Printf("Incorrect: %v\n", vals)
	fmt.Printf("Corrected: %v\n\n", correctedVals)
	return correctedVals
}
func middleNum(vals []string) (int, error) {
	v := vals[len(vals)/2]

	return strconv.Atoi(v)
}
