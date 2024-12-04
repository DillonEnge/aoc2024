package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"slices"
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
	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	s := string(b)

	ss := strings.Split(strings.TrimSpace(s), "\n")

	for i := range ss {
		l := ss[i]
		for j := range l {
			if string(l[j]) == "A" {
				if check(ss, i, j) {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)

	return nil
}

func check(s []string, i, j int) bool {
	if i+1 >= len(s) || i-1 < 0 || j+1 >= len(s[i]) || j-1 < 0 {
		return false
	}

	ls := string(s[i+1][j+1]) + string(s[i+1][j-1]) + string(s[i-1][j-1]) + string(s[i-1][j+1])
	if !slices.Contains([]string{"MMSS", "SMMS", "SSMM", "MSSM"}, ls) {
		return false
	}

	return true
}
