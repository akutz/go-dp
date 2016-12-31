package drivers

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/akutz/go-dp/lib"
)

const name = "andretti"

func init() {
	lib.Register(name, func() lib.Driver { return &driver{} })
}

type driver struct{}

func (d *driver) Init(ctx context.Context, w io.Writer) error {
	var (
		ok bool
		v  = fmt.Sprintf("driver.%s", name)
	)
	if ctx != nil {
		vo := ctx.Value("name")
		if v, ok = vo.(string); !ok {
			fmt.Fprintf(
				os.Stderr,
				"error: invalid type for context key \"name\": %T\n",
				vo)
			os.Exit(1)
		}
	}
	fmt.Fprintf(w, "%s.lib.RandInt64\t\t\t\t%d\n", v, lib.RandInt64)
	return nil
}

func (d *driver) Name() string { return name }
