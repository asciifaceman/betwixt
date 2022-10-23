package conf

type ProvisionerConfig struct {
	Ansible *AnsibleConfiguration
}

func NewProvisionerConfig(global *Global, local *Local) (*ProvisionerConfig, error) {
	p := &ProvisionerConfig{}

	// ansible
	p.Ansible = global.Ansible
	err := p.Ansible.MergeLocal(local.Ansible)
	if err != nil {
		return nil, err
	}

	return p, nil
}
