package pgpass

import (
	"bufio"
	"errors"
	"io"
)

// ErrNotEnoughFields indicates line doesn't contain enough fields
var ErrNotEnoughFields = errors.New("Not enough fields")

// Entry represents single entry in pgpass file
type Entry struct {
	Hostname, Port, Database, Username, Password string
}

// EntryReader reads entries from pgpass file
type EntryReader struct {
	s     *bufio.Scanner
	entry Entry
	err   error
}

// NewEntryReader returns new entry reader from provided r
func NewEntryReader(r io.Reader) *EntryReader {
	return &EntryReader{s: bufio.NewScanner(r)}
}

// Next reads in next entry and returns true on success
func (er *EntryReader) Next() (ok bool) {
	// Get next line, but skip comments and empty lines
	var line string
	for {
		if !er.s.Scan() {
			return
		}
		line = er.s.Text()
		if len(line) > 0 && line[0] != '#' {
			break
		}
	}

	fs := getFields(line)
	if len(fs) < 5 {
		er.err = ErrNotEnoughFields
		return
	}

	er.entry = Entry{fs[0], fs[1], fs[2], fs[3], fs[4]}
	return true
}

// Entry returns last read entry
func (er *EntryReader) Entry() Entry {
	return er.entry
}

// Err returns underlying error if any.
// Should be checked after unsuccessful call to Next().
func (er *EntryReader) Err() error {
	err := er.err
	if err == nil {
		err = er.s.Err()
	}
	return err
}

// getFields splits line into fields delimited by colon (:).
// All fields are properly unescaped.
func getFields(s string) []string {
	fs := make([]string, 0, 5)
	f := make([]rune, 0, len(s))

	var esc bool
	for _, c := range s {
		switch {
		case esc:
			f = append(f, c)
			esc = false
		case c == '\\':
			esc = true
		case c == ':':
			fs = append(fs, string(f))
			f = f[:0]
		default:
			f = append(f, c)
		}
	}
	return append(fs, string(f))
}
