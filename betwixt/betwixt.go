package betwixt

import (
	"os"

	"github.com/asciifaceman/betwixt/betwixt/lifecycle"
	"github.com/asciifaceman/betwixt/betwixt/state"
)

// New returns a new bootstrapped betwixt ready to do work
func New() (*Betwixt, error) {
	b := &Betwixt{}
	return b, b.Bootstrap()
}

type Betwixt struct {
	ProjectDirectory string
	State            *state.State
	Lifecycle        *lifecycle.Lifecycle
}

// Launch launches a betwixt instance based on the booted
// configuration state (lifecycle, merged configs, etc)

// Apply runs the booted configuration (provisioner, merged configs)
// against the betwixt instance if booted/exists

// Clean destroys the betwixt instance if exists
// and attempts to clean up locally

// Bootstrap manages launch configuration and setup
func (b *Betwixt) Bootstrap() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	b.ProjectDirectory = wd

	return nil
}

//
