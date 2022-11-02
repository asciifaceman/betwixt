package conf

import "github.com/imdario/mergo"

type AnsibleConfiguration struct {
	VaultPasswordFile string `json:"vaultPasswordFile,omitempty"`
}

func (a *AnsibleConfiguration) Merge(other *AnsibleConfiguration) error {
	if other == nil {
		return nil
	}

	return mergo.Merge(a, other, mergo.WithOverride)
}
