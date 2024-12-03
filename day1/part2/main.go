package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
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

	s1, m1 := []int{}, map[int]int{}

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

		m1[i2]++
	}

	sum := 0
	for _, v := range s1 {
		sum += v * m1[v]
	}

	fmt.Println(sum)

	return nil
}
