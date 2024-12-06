package main

import (
	"bufio"
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

	pages := map[string][]string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "|") {
			vals := strings.Split(strings.TrimSpace(line), "|")

			rules := []string{}
			if v, ok := pages[vals[0]]; ok {
				rules = v
			}

			rules = append(rules, vals[1])

			pages[vals[0]] = rules
		} else if strings.Contains(line, ",") {
			vals := strings.Split(strings.TrimSpace(line), ",")

			if valid(pages, vals) {
				v, err := middleNum(vals)
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

func valid(pages map[string][]string, vals []string) bool {
	readVals := []string{}

	for i := range vals {
		for _, v := range pages[vals[i]] {
			if slices.Contains(readVals, v) {
				return false
			}
		}

		readVals = append(readVals, vals[i])
	}

	return true
}

func middleNum(vals []string) (int, error) {
	v := vals[len(vals)/2]

	return strconv.Atoi(v)
}
