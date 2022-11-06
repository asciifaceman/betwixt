package provisioner

import (
	"fmt"

	"github.com/asciifaceman/betwixt/betwixt/conf"
)

type Ansible struct {
	Config *conf.AnsibleConfiguration
}

func (a *Ansible) LoadConfig(config interface{}) error {
	c, ok := config.(*conf.AnsibleConfiguration)
	if !ok {
		return fmt.Errorf("expected conf.AnsibleConfiguration for LoadConfig")
	}
	a.Config = c
	return nil

}

func (a *Ansible) Run() error {
	return nil
}

func (a *Ansible) Clean() error {
	return nil
}
