package str_manipulator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
