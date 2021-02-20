package trie

import (
	"github.com/xwzy/godan/assert"
	"testing"
)

func TestTrie_Search(t *testing.T) {
	tree := DefaultTrie()
	tree.Insert("hello")
	assert.True(t, tree.Search("hello"))
	assert.False(t, tree.Search("helloo"))
	assert.False(t, tree.Search("hell"))

	tree.Insert("helloo")
	assert.True(t, tree.Search("helloo"))
	assert.False(t, tree.Search("he"))

	tree.Insert("你好")
	assert.True(t, tree.Search("你好"))
	assert.False(t, tree.Search("你好啊"))
	assert.False(t, tree.Search("你"))

	tree.Insert("你好啊")
	assert.True(t, tree.Search("你好"))
	assert.True(t, tree.Search("你好啊"))
	assert.False(t, tree.Search("你"))
}

func TestTrie_StartsWith(t *testing.T) {
	tree := DefaultTrie()
	tree.Insert("hello")
	assert.True(t, tree.StartsWith("hello"))
	assert.False(t, tree.StartsWith("helloo"))
	assert.True(t, tree.StartsWith("hell"))

	tree.Insert("helloo")
	assert.True(t, tree.StartsWith("helloo"))
	assert.True(t, tree.StartsWith("he"))

	tree.Insert("你好")
	assert.True(t, tree.StartsWith("你好"))
	assert.False(t, tree.StartsWith("你好啊"))
	assert.True(t, tree.StartsWith("你"))

	tree.Insert("你好啊")
	assert.True(t, tree.StartsWith("你好"))
	assert.True(t, tree.StartsWith("你好啊"))
	assert.True(t, tree.StartsWith("你"))
}
