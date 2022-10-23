package betwixt

import (
	"fmt"

	"github.com/asciifaceman/betwixt/betwixt/conf"
	"github.com/asciifaceman/betwixt/betwixt/lifecycle"
	"github.com/asciifaceman/betwixt/betwixt/provisioner"
)

type Betwixt struct {
	Provisioner provisioner.Provisioner
	Lifecycle   lifecycle.Lifecycle
	Global      *conf.Global
	Local       *conf.Local
	State       *conf.State
}

func New(optFnc ...Option) *Betwixt {
	options := &Options{}
	for _, opt := range optFnc {
		opt(options)
	}

	b := &Betwixt{}

	b.Lifecycle = lifecycle.NewLifecycle(options.Lifecycle)
	b.Provisioner = provisioner.NewProvisioner(options.Provisioner)

	return b
}

func (b *Betwixt) IngestConfigs(local *conf.Local, global *conf.Global) error {
	fmt.Println("Ingesting configs...")

	b.Local = local
	b.Global = global

	p, err := conf.NewProvisionerConfig(b.Global, b.Local)
	if err != nil {
		return err
	}

	l, err := conf.NewLifecycleConfig(b.Global, b.Local)
	if err != nil {
		return err
	}

	err = b.Lifecycle.LoadConfig(l)
	if err != nil {
		return err
	}

	err = b.Provisioner.LoadConfig(p)
	if err != nil {
		return err
	}

	return nil
}

// CycleTest is the entrypoint for a test run cycle
func (b *Betwixt) CycleTest() error {
	return nil
}

// CycleClean is the entrypoint for a clean cycle
func (b *Betwixt) CycleClean() error {
	return nil
}
