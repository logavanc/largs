// This package aims to provide a simple way to quickly setup and parse command
// line parameters with go.
package largs

import (
	"fmt"
	_ "os"
)

type largs struct {
	defaultGroup *group
}

type group struct {
	name string
}

func New() (l *largs, err error) {
	l = new(largs)
	l.defaultGroup = new(group)

	return
}

// This function expects to be passed the full slice returned by os.Args
func (l *largs) Parse(a []string) (err error) {
	for i, v := range a {
		if i != 0 {
			fmt.Println(v)
		}
	}

	return
}
