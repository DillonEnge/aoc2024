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

		if safeish(is) {
			sum++
		}
	}

	fmt.Println(sum)

	return nil
}

func safeish(s []int) bool {
	if safe(s) {
		return true
	}

	fmt.Printf("Original unsafe: %v\n", s)
	for i := range s {
		ns := make([]int, len(s))
		copy(ns, s)
		ns = append(ns[:i], ns[i+1:]...)
		fmt.Printf("Testing: %v\n", ns)
		if len(ns) > 0 {
			if safe(ns) {
				fmt.Printf("Passed. %v\n", ns)
				return true
			}
			fmt.Printf("Failed. %v\n", ns)
		}
	}

	return false
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
