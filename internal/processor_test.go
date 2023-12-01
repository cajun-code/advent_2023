package internal

import "testing"

func TestProcessFile(t *testing.T) {
	var expected int64 = 142
	val := ProcessFile("../day01/test.txt")
	if expected != val {
		t.Fatalf("got %d expected %d", val, expected)
	}
}

func TestProcessFile2(t *testing.T) {
	var expected int64 = 281
	val := ProcessFile("../day01/test2.txt")
	if expected != val {
		t.Fatalf("got %d expected %d", val, expected)
	}
}
