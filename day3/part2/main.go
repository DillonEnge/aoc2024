package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"regexp"
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

	re := regexp.MustCompile(`(mul\([0-9]+,[0-9]+\))|(do\(\))|(don't\(\))`)

	matches := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()

		if m := re.FindAllString(t, -1); m != nil {
			matches = append(matches, m...)
		}
	}

	enabled := true
	for _, m := range matches {
		if m == "do()" {
			enabled = true
		} else if m == `don't()` {
			enabled = false
		} else {
			if !enabled {
				continue
			}

			f, err := factors(m)
			if err != nil {
				return err
			}

			sum += f[0] * f[1]
		}
	}

	fmt.Println(sum)

	return nil
}

func factors(s string) ([]int, error) {
	ss := strings.Split(strings.Replace(strings.Split(s, "(")[1], ")", "", -1), ",")

	ns := []int{}

	for _, v := range ss {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		ns = append(ns, i)
	}

	return ns, nil
}
