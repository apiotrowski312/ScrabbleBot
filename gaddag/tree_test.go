package gaddag

import (
	"reflect"
	"testing"
)

func Test_GraphGet(t *testing.T) {
	g := &node{
		children: map[rune]node{},
	}

	t.Run("Get children when nil", func(t *testing.T) {
		_, isOk := g.get('w')

		if isOk {
			t.Errorf("Getting root child has error. There should not be any child")
			return
		}
	})

	g = &node{
		children: map[rune]node{
			'w': node{
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
	g := node{
		children: map[rune]node{},
	}

	exampleNode := node{
		isWord:   false,
		children: map[rune]node{},
	}

	t.Run("Add children when nil", func(t *testing.T) {
		g.add('w', node{isWord: false})

		if reflect.DeepEqual(g.children['w'], exampleNode) != true {
			t.Errorf("Adding child to graph failed. Child: %v, example: %v", g.children['w'], exampleNode)
			return
		}
	})

	g = node{
		children: map[rune]node{},
	}
	exampleNode = node{
		isWord: false,
		children: map[rune]node{
			'w': node{
				isWord:   true,
				children: map[rune]node{},
			},
		},
	}

	t.Run("Add children to children", func(t *testing.T) {
		child := g.add('w', node{})
		child.add('w', node{isWord: true})

		if reflect.DeepEqual(g.children['w'], exampleNode) != true {
			t.Errorf("Adding child to graph failed. Child: %v, example: %v", g.children['w'], exampleNode)
			return
		}
	})
}
