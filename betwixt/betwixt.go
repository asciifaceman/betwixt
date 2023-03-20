// betwixt is the pillar struct of the infrastructure automation test suite
package betwixt

import (
	"os"

	"github.com/asciifaceman/betwixt/betwixt/conf"
	"github.com/asciifaceman/betwixt/betwixt/csl"
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
	State            state.State
	Lifecycle        lifecycle.Lifecycle
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

	// Read in global config
	global := &conf.Global{}
	err = global.Read()
	if err != nil {
		csl.Help("Did you run [ betwixt config init ] to create glboal config?")
		return err
	}

	// Read in local config
	local := conf.NewLocal(b.ProjectDirectory)
	err = local.Read()
	if err != nil {
		csl.Help("Did you run [ betwixt init ] in this directory?")
		return err
	}

	// Create Lifecycle and Provisioner and merge respective configs
	lc := lifecycle.StringToLifecycle(local.Lifecycle)
	lifecycle := lc.New(local.AWS)
	b.Lifecycle = lifecycle

	state, err := lc.StateManager(b.ProjectDirectory)
	if err != nil {
		return err
	}
	b.State = state

	// Return configured for task assignment

	return nil
}

//
