package pgpass

import (
	"bytes"
	"testing"
)

func TestPasswordFrom(t *testing.T) {
	data := `
localhost:5432:db:root:god123
*:5000:db:root:god5000
localhost:1000:db:root:god1000
localhost:*:db:root:local_god
localhost:*:db:*:buddy
*:*:*:*:anything
`

	tests := []struct {
		host, user, pass string
	}{
		{"localhost:5432", "root", "god123"},
		{"localhost", "root", "god123"},
		{"localhost:5000", "root", "god5000"},
		{"localhost:1000", "root", "god1000"},
		{"localhost:123", "root", "local_god"},
		{"localhost:5432", "user", "buddy"},
		{"otherhost:5432", "root", "anything"},

		{"noport", "root", "anything"},
		{"", "", "anything"},
	}

	for _, tt := range tests {
		pass, err := PasswordFrom(tt.host, tt.user, bytes.NewBufferString(data))
		if err != nil {
			t.Error(tt, err)
		}
		if pass != tt.pass {
			t.Error(tt, "Wrong password: ", pass)
		}
	}
}
