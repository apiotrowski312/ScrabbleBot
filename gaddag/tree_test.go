package gaddag

import (
	"reflect"
	"testing"
)

func Test_GraphGet(t *testing.T) {
	g := &Node{
		children: map[rune]Node{},
	}

	t.Run("Get children when nil", func(t *testing.T) {
		_, isOk := g.get('w')

		if isOk {
			t.Errorf("Getting root child has error. There should not be any child")
			return
		}
	})

	g = &Node{
		children: map[rune]Node{
			'w': Node{
				isWord: false,
			},
		},
	}

	t.Run("Get children when nil", func(t *testing.T) {
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

		if reflect.DeepEqual(g.children['w'], exampleNode) != true {
			t.Errorf("Adding child to graph failed. Child: %v, example: %v", g.children['w'], exampleNode)
			return
		}
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

		if reflect.DeepEqual(g.children['w'], exampleNode) != true {
			t.Errorf("Adding child to graph failed. Child: %v, example: %v", g.children['w'], exampleNode)
			return
		}
	})
}
