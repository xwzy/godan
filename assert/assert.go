package assert

import "testing"

func True(t *testing.T, a interface{}) {
	if a != true {
		t.Fatalf("%s != true", a)
	}
}

func False(t *testing.T, a interface{}) {
	if a != false {
		t.Fatalf("%s != false", a)
	}
}

func Equal(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func NotEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Fatalf("%s == %s", a, b)
	}
}

func SliceEqual(t *testing.T, a []interface{}, b []interface{}) {
	if len(a) != len(b) {
		t.Fatalf("size not equal: %d != %d", len(a), len(b))
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Fatalf("%v != %v", a, b)
		}
	}
}
