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

	// d = &SystemDevice{}
	// d.unitLocation = seed.Location
	// d.unitName = seed.Name
	// d.commandOn = seed.CommandOn
	// d.commandOff = seed.CommandOff
	// d.claimSeconds = seed.ClaimSeconds
	// d.executor = seed.Executor

	// if seed.Logger == nil {

	// 	// Lifted straight from https://golang.org/src/pkg/log/log.go
	// 	d.logger = log.New(
	// 		os.Stderr,
	// 		fmt.Sprintf("%s/%s: ", seed.Name, seed.Location),
	// 		log.LstdFlags)

	// } else {
	// 	d.logger = seed.Logger
	// }

	// // Create mutexes.
	// d.stateLock = new(sync.Mutex)

	// // Synchronize initialized device state.
	// // powered bool initializes to false, so make the hardware match...
	// d.setPowerOff()

	// // Launch timer go routines.
	// go d.claimTimer()
	// go d.powerTimer()

	// d.logger.Println("New device created.")

	return
}
