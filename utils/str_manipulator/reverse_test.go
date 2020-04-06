package str_manipulator

import (
	"flag"
	"testing"

	"github.com/bmizerany/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

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
