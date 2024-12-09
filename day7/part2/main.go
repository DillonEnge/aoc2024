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

	sum := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()

		ss := strings.Split(t, ": ")
		result, err := strconv.Atoi(ss[0])
		if err != nil {
			return err
		}

		ok, err := valid(ss[1], result)
		if err != nil {
			return err
		}
		if ok {
			sum += result
		}
	}

	fmt.Println(sum)

	return nil
}

func valid(s string, result int) (bool, error) {
	f := strings.Fields(s)

	combos := getAllCombinations([]string{"+", "*", "||"}, len(f)-1)

	for _, ops := range combos {
		stmt := combineValsAndOps(f, ops)
		total, err := parseMathExp(strings.Join(stmt, " "))
		if err != nil {
			return false, err
		}
		if int(total) == result {
			return true, nil
		}
	}

	return false, nil
}

func getAllCombinations(opts []string, n int) [][]string {
	result := [][]string{{}}
	for range n {
		newResult := [][]string{}
		for _, v := range result {
			cv := make([]string, len(v))
			copy(cv, v)
			for _, opt := range opts {
				newResult = append(newResult, append(cv, opt))
			}
		}

		result = newResult
	}

	return result
}

func combineValsAndOps(vals []string, ops []string) []string {
	c := []string{}

	for i := range ops {
		c = append(c, vals[i], ops[i])
	}

	c = append(c, vals[len(vals)-1])

	return c
}

func parseMathExp(exp string) (int64, error) {
	f := strings.Fields(exp)
	var result int64
	lastOp := ""
	for i := range f {
		switch f[i] {
		case "*":
			fallthrough
		case "+":
			fallthrough
		case "||":
			lastOp = f[i]
		default:
			v, err := strconv.Atoi(f[i])
			if err != nil {
				return result, err
			}
			switch lastOp {
			case "*":
				result *= int64(v)
			case "+":
				result += int64(v)
			case "||":
				s1 := strconv.Itoa(int(result))
				s2 := strconv.Itoa(v)
				s3 := s1 + s2
				ni, err := strconv.Atoi(s3)
				if err != nil {
					return result, err
				}
				result = int64(ni)
			default:
				result += int64(v)
			}

			lastOp = f[i]
		}
	}

	return result, nil
}
