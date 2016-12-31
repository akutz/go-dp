package lib

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var (
	driverMap = map[string]func() Driver{}
	r         = rand.New(rand.NewSource(time.Now().UnixNano()))

	// RandInt64 is a random, 64-bit integer that is initialized once when
	// the package is initialized.
	RandInt64 = r.Int63()
)

// Driver is the interface that describes a driver-pattern type.
type Driver interface {
	// Initialize the driver.
	Init(ctx context.Context, w io.Writer) error

	// Name returns the name of the driver.
	Name() string
}

// Register registers a new driver constructor.
func Register(name string, ctor func() Driver) {
	driverMap[name] = ctor
}

// NewDriver returns a new driver instance for the specified driver name.
func NewDriver(name string) (Driver, error) {
	ctor, ok := driverMap[name]
	if !ok {
		return nil, fmt.Errorf("error: invalid driver name: %s", name)
	}
	return ctor(), nil
}
