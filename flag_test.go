package largs

import (
	"testing"
)

var flagsErr error
var f *Flag

func TestNewFlag(t *testing.T) {
	if f, flagsErr = newFlag(); flagsErr != nil {
		t.Fail()
	}
}

func TestShort(t *testing.T) {
	test := []rune{
		'a',
		'%',
		'0',
		'⚝',
		'\u0416',
	}

	for _, v := range test {
		if flagsErr = f.SetShort(v); flagsErr != nil {
			t.Fail()
		}
	}
}

func TestLong(t *testing.T) {
	test := []string{
		"abcdefg",
		"%#$%^#$%^",
		"0123456789",
	}

	for _, v := range test {
		if flagsErr = f.SetLong(v); flagsErr != nil {
			t.Fail()
		}
	}
}
