package largs

import (
	"testing"
)

var err error
var l *largs

func TestNew(t *testing.T) {
	if l, err = New(); err != nil {
		t.Fail()
	}
}

func TestParse(t *testing.T) {

	if err = l.Parse([]string{"a", "b", "c"}); err != nil {
		t.Fail()
	}
}
