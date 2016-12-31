// +build go1.8
// +build plugins,!drivers

package main

import (
	"context"
	"fmt"
	"os"
	"plugin"
	"text/tabwriter"

	"github.com/akutz/go-dp/lib"
)

func main() {
	// print the lib package's random int
	w := tabwriter.NewWriter(os.Stdout, 1, 4, 8, ' ', 0)
	defer w.Flush()

	// print the lib package's random int
	fmt.Fprintf(w, "main.lib.RandInt64\t\t\t\t%d\n", lib.RandInt64)

	// if there are fewer than two command-line arguments then there a path
	// to the plug-in was not provided; abort with error
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s PLUGIN\n", os.Args[0])
		return
	}

	// load the plug-in. this registers the driver
	p, err := plugin.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// get the plug-in's default context
	ctxSymbol, err := p.Lookup("DefaultContext")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// type assert the plug-in's default context. even though a context is an
	// interface type, exported symbols are always available as pointers. so
	// the type assertion must be for a *context.Context, not a context.Context.
	pCtx, ok := ctxSymbol.(*context.Context)
	if !ok {
		fmt.Fprintf(
			os.Stderr,
			"error: invalid context type for symbol \"DefaultContext\": %T\n",
			ctxSymbol)
		os.Exit(1)
	}

	// get the context pointed to by pCtx
	ctx := *pCtx

	// create and initialize a new driver using the plug-in's default
	// context. the driver's Init function will print the same random number
	// that's already been printed above since it shared the same `lib` pacakge
	d, err := lib.NewDriver("andretti")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	d.Init(ctx, w)
}
