package largs

import (
	"testing"
)

var largsErr error
var l *largs

func TestNew(t *testing.T) {
	if l, largsErr = New(); largsErr != nil {
		t.Fail()
	}
}

func TestParse(t *testing.T) {
	if largsErr = l.parse([]string{"a", "b", "c"}); largsErr != nil {
		t.Fail()
	}
}
