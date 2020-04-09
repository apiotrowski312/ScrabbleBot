package str_manipulator

import (
	"testing"

	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/bmizerany/assert"
)

func Test_RemoveCharacters(t *testing.T) {
	toReverse := RemoveCharacters("Hello", "l")
	expect := test_utils.GetGoldenFileString(t, toReverse, t.Name()+"_1", *update)
	assert.Equal(t, expect, toReverse)

	toReverse = RemoveCharacters("Hello", "helo")
	expect = test_utils.GetGoldenFileString(t, toReverse, t.Name()+"_2", *update)
	assert.Equal(t, expect, toReverse)

	toReverse = RemoveCharacters("Hello", "Helo")
	expect = test_utils.GetGoldenFileString(t, toReverse, t.Name()+"_3", *update)
	assert.Equal(t, expect, toReverse)

	toReverse = RemoveCharacters("Hello", "")
	expect = test_utils.GetGoldenFileString(t, toReverse, t.Name()+"_4", *update)
	assert.Equal(t, expect, toReverse)

	toReverse = RemoveCharacters("Very long string which has a lot of differnet characters %$#@!", " ")
	expect = test_utils.GetGoldenFileString(t, toReverse, t.Name()+"_5", *update)
	assert.Equal(t, expect, toReverse)

	toReverse = RemoveCharacters("Very long string which has a lot of differnet characters %$#@!", "long%@#$")
	expect = test_utils.GetGoldenFileString(t, toReverse, t.Name()+"_6", *update)
	assert.Equal(t, expect, toReverse)
}
