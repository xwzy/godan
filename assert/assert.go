package assert

import (
	"runtime/debug"
	"testing"
)

func True(t *testing.T, a interface{}) {
	if a != true {
		debug.PrintStack()
		t.Fatalf("%s != true", a)
	}
}

func False(t *testing.T, a interface{}) {
	if a != false {
		debug.PrintStack()
		t.Fatalf("%s != false", a)
	}
}

func Equal(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		debug.PrintStack()
		t.Fatalf("%s != %s", a, b)
	}
}

func NotEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		debug.PrintStack()
		t.Fatalf("%s == %s", a, b)
	}
}

func SliceEqual(t *testing.T, a []interface{}, b []interface{}) {
	if len(a) != len(b) {
		debug.PrintStack()
		t.Fatalf("size not equal: %d != %d", len(a), len(b))
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			debug.PrintStack()
			t.Fatalf("%v != %v", a, b)
		}
	}
}
