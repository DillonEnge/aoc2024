package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"math"
	"os"
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

	s1, s2 := []int{}, []int{}

	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()

		ss := strings.Split(t, "   ")

		i1, err := strconv.Atoi(ss[0])
		if err != nil {
			return err
		}

		s1 = append(s1, i1)

		i2, err := strconv.Atoi(ss[1])
		if err != nil {
			return err
		}

		s2 = append(s2, i2)
	}

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})

	sum := 0
	for i := range s1 {
		sum += int(math.Abs(float64(s1[i] - s2[i])))
	}

	fmt.Println(sum)

	return nil
}
