package lcd

import "testing"
import "github.com/stretchr/testify/assert"

func TestAlwaysReturn1(t *testing.T) {
	result := AlwaysReturn1()
	assert.Equal(t, 1, result, "Should be 1")
}
