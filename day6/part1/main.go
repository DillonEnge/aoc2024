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

	// sum := 0

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	ss := strings.Split(strings.TrimSpace(string(b)), "\n")

	x, y := findStart(ss)

	n := NewNavigator(ss)
	n.SetDirection(UP)
	n.SetPosition(x, y)

	positions := []string{fmt.Sprintf("%d,%d", x, y)}

	for n.Next() {
		s := n.String()

		if s == "#" {
			n.Previous()
			n.RotateClockwise()
		} else {
			x, y := n.Position()
			positions = append(positions, fmt.Sprintf("%d,%d", x, y))
		}
	}

	slices.Sort(positions)
	positions = slices.Compact(positions)

	fmt.Printf("distinct position count: %d\n", len(positions))

	return nil
}

func findStart(ss []string) (x, y int) {
	for i := range ss {
		if index := strings.Index(ss[i], "^"); index > -1 {
			return index, i
		}
	}

	return -1, -1
}
