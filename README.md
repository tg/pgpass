# pgpass [![GoDoc](https://godoc.org/github.com/tg/pgpass?status.svg)](https://godoc.org/github.com/tg/pgpass) [![Build Status](https://travis-ci.org/tg/pgpass.svg?branch=master)](https://travis-ci.org/tg/pgpass)
Access PostreSQL password file as described in [libpq documentation](http://www.postgresql.org/docs/9.4/static/libpq-pgpass.html).
This libarary will allow you to iterate all the entries in pgpass file as well as easily access passwords for a given host and username pair.

## Default file
When applicable, the default pgpass file being used is read from `~/.pgpass` location.
This is not going to work on Windows, I uderstand, but I don't own one and cannot test. Appropriate patch reading from `%APPDATA%\postgresql\pgpass.conf` should be quite trivial though.

## Examples
### Get password
```go
// Input:
// localhost:5432:db:tg:letmein
// *:2345:db:tg:trustno1
// *:*:db:*:superman

package main

import (
	"fmt"

	"github.com/tg/pgpass"
)

func main() {
	for _, host := range []string{"localhost", "remotehost:2345", "spacehost"} {
		pass, err := pgpass.Password(host, "tg")
		if err != nil {
			panic(err)
		}		
		fmt.Println(pass)		
	}
}

// Output:
// letmein
// trustno1
// superman
```
### Inject password into an URL
```go
package main

import (
	"fmt"

	"github.com/tg/pgpass"
)

func main() {
	u, _ := pgpass.UpdateURL("postgres://tg@localhost:5432/db?sslmode=disable")
	fmt.Println(u) // postgres://tg:letmein@localhost:5432/db?sslmode=disable
	// Now you can call sql.Open("postgres", u)
}
```
## Why?
Because my password is in `~/.pgpass` already, for use with `psql`.
Because I don't want to duplicate it in the config file for each process.
Because someone will send the config file to his mum by email one day,
or store it in github.
