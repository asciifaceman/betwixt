package provisioner

import "github.com/asciifaceman/betwixt/betwixt/conf"

type Ansible struct {
	Config *conf.AnsibleConfiguration
}

func NewAnsible() *Ansible {
	return &Ansible{}
}

func (a *Ansible) Run() error {
	return nil
}

func (a *Ansible) LoadConfig(conf *conf.ProvisionerConfig) error {
	a.Config = conf.Ansible
	return nil
}
