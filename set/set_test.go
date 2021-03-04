package set

import (
	"github.com/xwzy/godan/assert"
	"testing"
)

func TestSet(t *testing.T) {
	s1 := DefaultSet()
	s2 := DefaultSet()

	s1.Add(1)
	s2.Add(2)
	s1.Add(3)
	s2.Add(3)

	assert.True(t, s1.InterSection(s2).Exist(3))
}
