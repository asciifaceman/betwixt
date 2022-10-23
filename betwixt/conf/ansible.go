package conf

type AnsibleConfiguration struct {
	VaultPasswordFile string `json:"vaultPasswordFile"`
}

func (a *AnsibleConfiguration) MergeLocal(local *LocalAnsibleConfiguration) error {
	if local == nil {
		return nil
	}

	// TODO: Make this smarter / reflect maybe? Mergo won't work
	if local.VaultPasswordFile != "" {
		a.VaultPasswordFile = local.VaultPasswordFile
	}

	return nil
}

type LocalAnsibleConfiguration struct {
	VaultPasswordFile string `json:"vaultPasswordFile,omitempty"`
}
