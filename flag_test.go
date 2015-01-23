package largs

import (
	"testing"
)

func TestNewFlag(t *testing.T) {

	type InitOption struct {
		Description string
		Options     []flagOption
		ShouldError bool
	}

	args := []InitOption{
		InitOption{
			Description: "none",
			Options:     []flagOption{},
			ShouldError: true},
		InitOption{
			Description: "short",
			Options:     []flagOption{Short('s')},
			ShouldError: false},
		InitOption{
			Description: "long",
			Options:     []flagOption{Long("Long")},
			ShouldError: false},
		InitOption{
			Description: "countable",
			Options:     []flagOption{Countable},
			ShouldError: true},
		InitOption{
			Description: "required",
			Options:     []flagOption{Required},
			ShouldError: true},
		InitOption{
			Description: "short, long",
			Options:     []flagOption{Short('s'), Long("Long")},
			ShouldError: false},
		InitOption{
			Description: "short, countable",
			Options:     []flagOption{Short('s'), Countable},
			ShouldError: false},
		InitOption{
			Description: "short, required",
			Options:     []flagOption{Short('s'), Required},
			ShouldError: false},
		InitOption{
			Description: "long, countable",
			Options:     []flagOption{Long("Long"), Countable},
			ShouldError: false},
		InitOption{
			Description: "long, required",
			Options:     []flagOption{Long("Long"), Required},
			ShouldError: false},
		InitOption{
			Description: "countable, required",
			Options:     []flagOption{Countable, Required},
			ShouldError: true},
	}

	for _, v := range args {
		_, err := NewFlag(v.Options...)

		if v.ShouldError && err == nil {
			t.Log("Was supposed to error but did not with options:")
			t.Log(v.Description)
			t.Fail()

		} else if !v.ShouldError && err != nil {
			t.Log("Was not supposed to error but did with options:")
			t.Log(v.Description)
			t.Fail()
		}
	}
}
