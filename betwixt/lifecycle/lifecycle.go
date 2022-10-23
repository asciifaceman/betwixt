package lifecycle

import "github.com/asciifaceman/betwixt/betwixt/conf"

type LifecycleOption int

const (
	None LifecycleOption = iota
	AWSLifecycle
)

func StringToLifecycle(lc string) LifecycleOption {
	if lc == "aws" {
		return AWSLifecycle
	}

	return None
}

type Lifecycle interface {
	LoadConfig(conf *conf.LifecycleConfig) error
	Launch() error
	Destroy() error
}

func NewLifecycle(opt LifecycleOption) Lifecycle {
	if opt == AWSLifecycle {
		return NewAWSRemote()
	}
	return nil
}
