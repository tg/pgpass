package pgpass

import (
	"os"
	"os/user"
	"path"
)

// OpenDefault opens default pgpass file, which is ~/.pgpass
func OpenDefault() (f *os.File, err error) {
	usr, err := user.Current()
	if err != nil {
		return
	}
	// TODO: check file permission is 0600
	return os.Open(path.Join(usr.HomeDir, ".pgpass"))
}
