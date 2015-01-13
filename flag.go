package largs

import (
	_ "fmt"
	_ "os"
)

//                                                                       PUBLIC
///////////////////////////////////////////////////////////////////////////////

type Flag struct {
	short rune
	long  string
	count bool
}

// Set the short flag name (-v).
func (f *Flag) SetShort(s rune) (err error) {
	f.short = s
	return
}

// Set the long flag name (--verbose)
func (f *Flag) SetLong(s string) (err error) {
	f.long = s
	return
}

//                                                                      PRIVATE
///////////////////////////////////////////////////////////////////////////////

func newFlag() (f *Flag, err error) {
	f = &Flag{}
	return
}
