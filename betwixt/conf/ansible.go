package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/asciifaceman/betwixt/betwixt/csl"
	"github.com/imdario/mergo"
)

type AnsibleConfiguration struct {
	VaultPasswordFile string `json:"vaultPasswordFile,omitempty"`
}

func (a *AnsibleConfiguration) Init(path string, editor string) error {
	csl.Info("Do you wish to edit the ansible configuration before saving?")
	edit := csl.YesNoPrompt()
	if !edit {
		return nil
	}

	a.VaultPasswordFile = "changeme"

	tmpAnsible, err := os.CreateTemp(path, "ansible*")
	if err != nil {
		return err
	}

	tmpfile := tmpAnsible.Name()

	bin, err := json.MarshalIndent(a, "", " ")
	if err != nil {
		return err
	}

	ioutil.WriteFile(tmpfile, bin, 0755)
	tmpAnsible.Sync()
	tmpAnsible.Close()

	csl.Open(editor, tmpfile)

	file, err := os.Open(tmpfile)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, a)
	if err != nil {
		return err
	}
	file.Close()

	return os.Remove(tmpfile)
}

func (a *AnsibleConfiguration) Merge(other *AnsibleConfiguration) error {
	if other == nil {
		return nil
	}

	return mergo.Merge(a, other, mergo.WithOverride)
}
