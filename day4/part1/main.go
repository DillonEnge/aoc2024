package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
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
			if string(l[j]) == "X" {
				sum += surroundingWordCount(ss, i, j)
			}
		}
	}

	fmt.Println(sum)

	return nil
}

func surroundingWordCount(s []string, i, j int) int {
	count := 0

	if j+3 < len(s[i]) {
		//X***
		if string(s[i][j+1])+string(s[i][j+2])+string(s[i][j+3]) == "MAS" {
			count++
		}
	}
	if j-3 >= 0 {
		//***X
		if string(s[i][j-1])+string(s[i][j-2])+string(s[i][j-3]) == "MAS" {
			count++
		}
	}
	if i+3 < len(s) {
		//X
		//*
		//*
		//*
		if string(s[i+1][j])+string(s[i+2][j])+string(s[i+3][j]) == "MAS" {
			count++
		}
		if j+3 < len(s[i]) {
			//X
			//-*
			//--*
			//---*
			if string(s[i+1][j+1])+string(s[i+2][j+2])+string(s[i+3][j+3]) == "MAS" {
				count++
			}
		}
		if j-3 >= 0 {
			//---X
			//--*
			//-*
			//*
			if string(s[i+1][j-1])+string(s[i+2][j-2])+string(s[i+3][j-3]) == "MAS" {
				count++
			}
		}
	}
	if i-3 >= 0 {
		//*
		//*
		//*
		//X
		if string(s[i-1][j])+string(s[i-2][j])+string(s[i-3][j]) == "MAS" {
			count++
		}
		if j+3 < len(s[i]) {
			//---*
			//--*
			//-*
			//X
			if string(s[i-1][j+1])+string(s[i-2][j+2])+string(s[i-3][j+3]) == "MAS" {
				count++
			}
		}
		if j-3 >= 0 {
			//*
			//-*
			//--*
			//---X
			if string(s[i-1][j-1])+string(s[i-2][j-2])+string(s[i-3][j-3]) == "MAS" {
				count++
			}
		}
	}

	return count
}
