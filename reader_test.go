package pgpass

import (
	"bytes"
	"testing"
)

func TestEntryReader(t *testing.T) {
	data := bytes.NewBufferString(
		`host:5432:db:root:god123
# Comment line
host:5432:db:root:god\:123
host:5432:db:root:
host:5432::root:god 123
*:*:*:root:god123

host:5432:db:ro\:ot:god123
host:5432:db:root\\:god123
`)

	expected := []Entry{
		Entry{"host", "5432", "db", "root", "god123"},
		Entry{"host", "5432", "db", "root", "god:123"},
		Entry{"host", "5432", "db", "root", ""},
		Entry{"host", "5432", "", "root", "god 123"},
		Entry{"*", "*", "*", "root", "god123"},
		Entry{"host", "5432", "db", "ro:ot", "god123"},
		Entry{"host", "5432", "db", `root\`, "god123"},
	}

	r := NewEntryReader(data)
	for n, exp := range expected {
		if !r.Next() {
			t.Fatalf("No more entries (%d): %s", n, r.Err())
		}
		if e := r.Entry(); e != exp {
			t.Errorf("%d: expected %v, got %v", n, exp, e)
		}
	}
	if err := r.Err(); err != nil {
		t.Error(err)
	}
}
