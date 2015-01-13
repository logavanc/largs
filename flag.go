package largs

import (
	_ "fmt"
	_ "os"
)

//                                                                       PUBLIC
///////////////////////////////////////////////////////////////////////////////

// Set the short flag name (-v).
func (f *flag) SetShort(s rune) (err error) {
	f.short = s
	return
}

// Set the long flag name (--verbose)
func (f *flag) SetLong(s string) (err error) {
	f.long = s
	return
}

//                                                                      PRIVATE
///////////////////////////////////////////////////////////////////////////////

type flag struct {
	short rune
	long  string
	count bool
}

func newFlag() (f *flag, err error) {
	f = new(flag)
	return
}
