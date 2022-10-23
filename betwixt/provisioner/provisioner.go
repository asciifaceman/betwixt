package provisioner

import "github.com/asciifaceman/betwixt/betwixt/conf"

type ProvisionerOption int

const (
	None ProvisionerOption = iota
	AnsibleProvisioner
)

func StringToProvisioner(prov string) ProvisionerOption {
	if prov == "ansible" {
		return AnsibleProvisioner
	}

	return None
}

type Provisioner interface {
	LoadConfig(*conf.ProvisionerConfig) error
	Run() error
}

func NewProvisioner(opt ProvisionerOption) Provisioner {
	if opt == AnsibleProvisioner {
		return NewAnsible()
	}

	return nil
}
