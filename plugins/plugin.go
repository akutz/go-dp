package main

import "C"

import (
	"context"

	// load the drivers
	_ "github.com/akutz/go-dp/drivers"
)

// DefaultContext is the plug-in's default context.
var DefaultContext context.Context = context.WithValue(
	context.Background(), "name", "plugin.andretti")
