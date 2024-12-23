package nanoid

import (
	"strings"
	"testing"
)

func TestDefaults(t *testing.T) {
	got := Nanoid(DefaultLength, DefaultCharset)
	if len(got) != DefaultLength {
		t.Errorf("Nanoid(DefaultLength, DefaultCharset) length = %d; want %d", len(got), DefaultLength)
	}
	for _, c := range got {
		if !strings.ContainsRune(DefaultCharset, c) {
			t.Errorf("Nanoid(DefaultLength, DefaultCharset) contains char %c; it shouldn't.", c)
		}
	}
}

func TestLength200(t *testing.T) {
	got := Nanoid(200, DefaultCharset)
	if len(got) != 200 {
		t.Errorf("Nanoid(200, DefaultCharset) length = %d; want 200", len(got))
	}
}

func TestCharsetABC(t *testing.T) {
	got := Nanoid(DefaultLength, "ABC")
	for _, c := range got {
		if !strings.ContainsRune("ABC", c) {
			t.Errorf("Nanoid(DefaultLength, \"ABC\") contains a char not in \"ABC\": %c.", c)
		}
	}
}
