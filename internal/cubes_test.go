package internal_test

import (
	"github.com/cajun-code/advent/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlayCubeFile(t *testing.T) {
	var expected int64 = 8
	val, _ := internal.PlayCubeGame("../day02/test.txt")
	assert.Equal(t, expected, val)
}

func TestPowerCube(t *testing.T) {
	var expected int64 = 2286
	_, val := internal.PlayCubeGame("../day02/test.txt")
	assert.Equal(t, expected, val)
}
