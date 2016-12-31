// +build drivers,!plugins

package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/akutz/go-dp/lib"

	// load the drivers
	_ "github.com/akutz/go-dp/drivers"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 1, 4, 8, ' ', 0)
	defer w.Flush()

	// print the lib package's random int
	fmt.Fprintf(w, "main.lib.RandInt64\t\t\t\t%d\n", lib.RandInt64)

	// create and initialize a new driver. it will print the same random
	// number that's already been printed above since it shared the same
	// `lib` package
	d, err := lib.NewDriver("andretti")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	d.Init(nil, w)
}
