package str_manipulator

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func Test_Reverse(t *testing.T) {
	toReverse := Reverse("Hello")
	pattern := "olleH"
	assert.Equal(t, toReverse, pattern)

	toReverse = Reverse("")
	pattern = ""
	assert.Equal(t, toReverse, pattern)

	toReverse = Reverse("!@#$#$")
	pattern = "$#$#@!"
	assert.Equal(t, toReverse, pattern)
}
