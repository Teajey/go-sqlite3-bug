package bug

import "testing"

func FatalAssertEq[C comparable](t *testing.T, context string, expected, actual C) {
	if expected != actual {
		t.Fatalf("%s: %#v != %#v", context, expected, actual)
	}
}

func FailIfErr(t *testing.T, err error, context string) {
	if err != nil {
		t.Fatalf("%s: %s\n", context, err)
	}
}
