package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"math"
	"os"
	"slices"
	"sort"
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

	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()

		ss := strings.Split(t, " ")

		is := []int{}

		for _, v := range ss {
			i, err := strconv.Atoi(v)
			if err != nil {
				return err
			}

			is = append(is, i)
		}

		if safe(is) {
			sum++
		}
	}

	fmt.Println(sum)

	return nil
}

func safe(s []int) bool {
	if !decreasing(s) && !increasing(s) {
		return false
	}

	for i := range s {
		if i == 0 {
			continue
		}

		diff := int(math.Abs(float64(s[i] - s[i-1])))

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func decreasing(s []int) bool {
	s2 := make([]int, len(s))
	copy(s2, s)
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] > s2[j]
	})

	return slices.Equal(s, s2)
}

func increasing(s []int) bool {
	s2 := make([]int, len(s))
	copy(s2, s)
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})

	return slices.Equal(s, s2)
}
