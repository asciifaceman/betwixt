package provisioner

type Provision int

const (
	None Provision = iota
	AnsibleProvisioner
)

func StringToProvision(prov string) Provision {
	switch prov {
	case "ansible":
		return AnsibleProvisioner
	default:
		return None
	}
}

func (p Provision) New() Provisioner {
	switch p {
	case None:
		return nil
	default:
		return nil
	case AnsibleProvisioner:
		return &Ansible{}
	}
}

type Provisioner interface {
	LoadConfig(interface{}) error
	Run() error
	Clean() error
}
