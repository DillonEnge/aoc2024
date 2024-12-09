package main

import "testing"

func TestParseMathExp(t *testing.T) {
	v, err := parseMathExp("1 + 1")
	if err != nil {
		t.Errorf("failed with err: %s", err.Error())
	}
	if int(v) != 2 {
		t.Errorf("failed, expected: %d, actual: %d", 2, int(v))
	}

	v, err = parseMathExp("10 * 10 + 2")
	if err != nil {
		t.Errorf("failed with err: %s", err.Error())
	}
	if int(v) != 102 {
		t.Errorf("failed, expected: %d, actual: %d", 102, int(v))
	}
}

func TestGetAllCombinations(t *testing.T) {
	combos := getAllCombinations([]string{"*", "+"}, 5)
	t.Logf("%v", combos)
	if len(combos) != 32 {
		t.Error("failed")
	}
}
