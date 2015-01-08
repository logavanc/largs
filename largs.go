package largs

import (
	"fmt"
	"os"
)

type Largs struct {
	stateLock *sync.Mutex

	Args []Arg
}

type Arg struct {
	Long  string
	Short string
}

// A system device must be created with a DeviceSeed.
func New() (d *SystemDevice, err error) {
	if err = validateSeed(seed); err != nil {
		return
	}

	return
}
