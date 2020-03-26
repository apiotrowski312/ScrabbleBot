package str_manipulator

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_Reverse(t *testing.T) {
	toReverse := Reverse("Hello")
	pattern := "olleH"
	assert.Equal(t, toReverse, pattern, "The two words should be the same.")

	toReverse = Reverse("")
	pattern = ""
	assert.Equal(t, toReverse, pattern, "The two words should be the same.")

	toReverse = Reverse("!@#$#$")
	pattern = "$#$#@!"
	assert.Equal(t, toReverse, pattern, "The two words should be the same.")
}
