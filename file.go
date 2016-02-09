package pgpass

import (
	"os"
	"os/user"
	"path"
)

// OpenDefault opens default pgpass file, which is ~/.pgpass.
// Current homedir will be retrieved by calling user.Current
// or using $HOME on failure.
func OpenDefault() (f *os.File, err error) {
	var homedir = os.Getenv("HOME")
	usr, err := user.Current()
	if err == nil {
		homedir = usr.HomeDir
	} else if homedir == "" {
		return
	}
	// TODO: check file permission is 0600
	return os.Open(path.Join(homedir, ".pgpass"))
}
