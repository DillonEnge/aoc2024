package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
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

	sum := 0

	for i := range ss {
		for j := range ss[i] {
			if string(ss[i][j]) != "^" && string(ss[i][j]) != "#" {
				nss := make([]string, len(ss))
				copy(nss, ss)

				if j == len(nss[i])-1 {
					nss[i] = nss[i][:j] + "#"
				} else {
					nss[i] = nss[i][:j] + "#" + nss[i][j+1:]
				}

				t := time.Now()

				n := NewNavigator(nss)
				n.SetDirection(UP)
				n.SetPosition(x, y)

				for n.Next() {
					s := n.String()

					if s == "#" {
						n.Previous()
						n.RotateClockwise()
					}

					if time.Since(t) > 5*time.Millisecond {
						sum++
						break
					}
				}
			}
		}
	}

	fmt.Printf("sum: %d\n", sum)

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
