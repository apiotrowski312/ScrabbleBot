package gaddag

import (
	"testing"

	"github.com/bmizerany/assert"
)

func Test_GraphGet(t *testing.T) {

	t.Run("Get children when nil", func(t *testing.T) {
		g := &Node{
			children: map[rune]Node{},
		}
		_, isOk := g.get('w')

		if isOk {
			t.Errorf("Getting root child has error. There should not be any child")
			return
		}
	})

	t.Run("Get children when nil", func(t *testing.T) {
		g := &Node{
			children: map[rune]Node{
				'w': Node{
					isWord: false,
				},
			},
		}
		_, isOk := g.get('w')
		if !isOk {
			t.Errorf("Getting root child has error. There should be a child")
			return
		}
	})
}

func Test_GraphAdd(t *testing.T) {
	t.Run("Add children when nil", func(t *testing.T) {
		g := Node{
			children: map[rune]Node{},
		}

		exampleNode := Node{
			isWord:   false,
			children: map[rune]Node{},
		}
		g.add('w', Node{isWord: false})

		assert.Equal(t, g.children['w'], exampleNode)
	})

	t.Run("Add children to children", func(t *testing.T) {
		g := Node{
			children: map[rune]Node{},
		}
		exampleNode := Node{
			isWord: false,
			children: map[rune]Node{
				'w': Node{
					isWord:   true,
					children: map[rune]Node{},
				},
			},
		}
		child := g.add('w', Node{})
		child.add('w', Node{isWord: true})

		assert.Equal(t, g.children['w'], exampleNode)
	})
}
