package str_manipulator

import (
	"testing"

	"github.com/apiotrowski312/goldtest"
)

func Test_RemoveCharacters(t *testing.T) {
	toReverse := RemoveCharacters("Hello", "l")
	goldtest.Assert(t, toReverse, "testdata/"+t.Name()+"_1")

	toReverse = RemoveCharacters("Hello", "helo")
	goldtest.Assert(t, toReverse, "testdata/"+t.Name()+"_2")

	toReverse = RemoveCharacters("Hello", "Helo")
	goldtest.Assert(t, toReverse, "testdata/"+t.Name()+"_3")

	toReverse = RemoveCharacters("Hello", "")
	goldtest.Assert(t, toReverse, "testdata/"+t.Name()+"_4")

	toReverse = RemoveCharacters("Very long string which has a lot of differnet characters %$#@!", " ")
	goldtest.Assert(t, toReverse, "testdata/"+t.Name()+"_5")

	toReverse = RemoveCharacters("Very long string which has a lot of differnet characters %$#@!", "long%@#$")
	goldtest.Assert(t, toReverse, "testdata/"+t.Name()+"_6")

}
