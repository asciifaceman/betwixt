package betwixt

import (
	"github.com/asciifaceman/betwixt/betwixt/lifecycle"
	"github.com/asciifaceman/betwixt/betwixt/provisioner"
)

type Options struct {
	Provisioner provisioner.ProvisionerOption
	Lifecycle   lifecycle.LifecycleOption
}

type Option func(*Options)

func WithProvisioner(prov string) Option {
	return func(opt *Options) {
		opt.Provisioner = provisioner.StringToProvisioner(prov)
	}
}

func withLifecycler(lc string) Option {
	return func(opt *Options) {
		opt.Lifecycle = lifecycle.StringToLifecycle(lc)
	}
}

func WithAnsible() Option {
	return func(opt *Options) {
		opt.Provisioner = provisioner.AnsibleProvisioner
	}
}

func WithAWS() Option {
	return func(opt *Options) {
		opt.Lifecycle = lifecycle.AWSLifecycle
	}
}
