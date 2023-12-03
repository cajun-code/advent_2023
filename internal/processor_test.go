package internal_test

import (
	"github.com/cajun-code/advent/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessFile(t *testing.T) {
	var expected int64 = 142
	val := internal.ProcessFile("../day01/test.txt")
	assert.Equal(t, expected, val)
}

func TestProcessFile2(t *testing.T) {
	var expected int64 = 281
	val := internal.ProcessFile("../day01/test2.txt")
	if expected != val {
		t.Fatalf("got %d expected %d", val, expected)
	}
}

func TestProcessFile3(t *testing.T) {
	var expected int64 = 88
	val := internal.ProcessFile("../day01/test3.txt")
	if expected != val {
		t.Fatalf("got %d expected %d", val, expected)
	}
}
