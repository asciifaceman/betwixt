package lifecycle

import (
	"fmt"

	"github.com/asciifaceman/betwixt/betwixt/conf"
	"github.com/asciifaceman/betwixt/betwixt/csl"
)

type AWSRemote struct {
	Config *conf.AwsConfiguration
}

func (a *AWSRemote) LoadConfig(config interface{}) error {
	c, ok := config.(*conf.AwsConfiguration)
	if !ok {
		return fmt.Errorf("expected AWSRemote config for LoadConfig")
	}
	a.Config = c
	return nil
}

func (a *AWSRemote) Init() error {
	return nil
}

func (a *AWSRemote) Launch() error {
	csl.Info("Launching instance in %s\n", a.Config.Region)
	csl.Success("not really this is just a dead end test")
	return nil
}
