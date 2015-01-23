package largs

import (
	"fmt"
	_ "os"
)

type flagOption func(*Flag) error

type Flag struct {
	short       rune
	long        string
	description string
	required    bool
	countable   bool
}

//                                                                     GETTERS
///////////////////////////////////////////////////////////////////////////////

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

// Get the flag's required setting.
func (f *Flag) Required() (c bool) {
	return f.required
}

//                                                                        INIT
///////////////////////////////////////////////////////////////////////////////

func NewFlag(options ...flagOption) (f *Flag, err error) {
	f = &Flag{}

	for _, option := range options {
		err := option(f)
		if err != nil {
			return nil, err
		}
	}

	// Either short or long parameter must be defined.
	if f.short == 0 && f.long == "" {
		return nil, fmt.Errorf(
			"NewFlag() must receive at least a Short() or Long() parameter.")
	}

	return
}

//                                                                     OPTIONS
///////////////////////////////////////////////////////////////////////////////

func Short(s rune) func(*Flag) error {
	return func(f *Flag) (err error) {
		f.short = s
		return
	}
}

func Long(l string) func(*Flag) error {
	return func(f *Flag) (err error) {
		f.long = l
		return
	}
}

func Description(d string) func(*Flag) error {
	return func(f *Flag) (err error) {
		f.description = d
		return
	}
}

func Countable(f *Flag) (err error) {
	f.countable = true
	return
}

func Required(f *Flag) (err error) {
	f.required = true
	return
}
