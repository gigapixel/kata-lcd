package lcd

import "testing"
import "github.com/stretchr/testify/assert"

func TestAlwaysReturn1(t *testing.T) {
	result := AlwaysReturn1()
	assert.Equal(t, 1, result, "Should be 1")
}

func TestNumber1(t *testing.T) {
	result := LCD(1)
	assert.Equal(t, result[0], "   ")
	assert.Equal(t, result[1], "  |")
	assert.Equal(t, result[2], "  |")
}

func TestNumber2(t *testing.T) {
	result := LCD(2)
	assert.Equal(t, result[0], " _ ")
	assert.Equal(t, result[1], " _|")
	assert.Equal(t, result[2], "|_ ")
}

func TestNumber3(t *testing.T) {
	result := LCD(3)
	assert.Equal(t, result[0], " _ ")
	assert.Equal(t, result[1], " _|")
	assert.Equal(t, result[2], " _|")
}


func TestNumber4(t *testing.T) {
	result := LCD(4)
	assert.Equal(t, result[0], "   ")
	assert.Equal(t, result[1], "|_|")
	assert.Equal(t, result[2], "  |")
}

func TestNumber5(t *testing.T) {
	result := LCD(5)
	assert.Equal(t, result[0], " _ ")
	assert.Equal(t, result[1], "|_ ")
	assert.Equal(t, result[2], " _|")
}

func TestNumber6(t *testing.T) {
	result := LCD(6)
	assert.Equal(t, result[0], " _ ")
	assert.Equal(t, result[1], "|_ ")
	assert.Equal(t, result[2], "|_|")
}

func TestNumber7(t *testing.T) {
	result := LCD(7)
	assert.Equal(t, result[0], " _ ")
	assert.Equal(t, result[1], "  |")
	assert.Equal(t, result[2], "  |")
}