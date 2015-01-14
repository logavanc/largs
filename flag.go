package largs

import (
	_ "fmt"
	_ "os"
)

//                                                                       PUBLIC
///////////////////////////////////////////////////////////////////////////////

type Flag struct {
	short     rune
	long      string
	countable bool
}

// Get the short flag name (-v).
func (f *Flag) Short() (s rune) {
	return f.short
}

// Get the long flag name (--version).
func (f *Flag) Long() (l string) {
	return f.long
}

// Get the flag's countable setting.
func (f *Flag) Countable() (c bool) {
	return f.countable
}

//                                                                      PRIVATE
///////////////////////////////////////////////////////////////////////////////

// Create a new (countable) short flag object and return a pointer to it.
func CountableFlag(short rune) (f *Flag, err error) {
	f = &Flag{
		short:     short,
		long:      "",
		countable: true,
	}

	return
}

// Create a new short flag object and return a pointer to it.
func ShortFlag(short rune) (f *Flag, err error) {
	f = &Flag{
		short:     short,
		long:      "",
		countable: false,
	}

	return
}

// Create a new long flag object and return a pointer to it.
func LongFlag(long string) (f *Flag, err error) {
	f = &Flag{
		short:     0,
		long:      long,
		countable: false,
	}

	return
}

// Create a new short flag object and return a pointer to it.
func Normalflag(short rune, long string) (f *Flag, err error) {
	f = &Flag{
		short:     short,
		long:      long,
		countable: false,
	}

	return
}
