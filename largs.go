// This package aims to provide a simple way to quickly setup and parse command
// line parameters with go.
package largs

import (
	"fmt"
	"os"
)

//                                                                       PUBLIC
///////////////////////////////////////////////////////////////////////////////

type Largs struct {
	Group
}

// Create a new largs instance to customize with groups, flags,
// named/positional arguments etc.
func New() (l *Largs, err error) {
	l = &Largs{}
	return
}

// This function attempts to parse the command line arguments according to the
// constructed largs instance.  If err is returned, it will be a string
// containing the reason the parameters weren't correct, followed by the help
// text (so just print it).
func (l *Largs) ParseArgs() (err error) {
	err = l.parse(os.Args)
	return
}

//                                                                      PRIVATE
///////////////////////////////////////////////////////////////////////////////

// This function expects to be passed the full slice returned by os.Args, and
// and is broken out here to allow testing.
func (l *Largs) parse(a []string) (err error) {
	for i, v := range a {
		if i != 0 {
			fmt.Println(v)
		}
	}

	return
}
