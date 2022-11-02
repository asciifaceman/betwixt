package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/asciifaceman/betwixt/betwixt/csl"
)

var (
	configDir  string
	configFile string
)

// Global represents the system wide configuration for
// betwixt which contains sane defaults to be overriden
// at the project level, as well as some handys
type Global struct {
	Editor  string                `json:"editor"`
	AWS     *AwsConfiguration     `json:"aws"`
	Ansible *AnsibleConfiguration `json:"ansible"`
}

func (g *Global) GetFilename() string {
	return configFile
}

func (g *Global) Init() error {
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		return err
	}

	pathEditor := os.Getenv("EDITOR")
	// TODO: Need a different option prompt for this that can support os.Getenv default
	editor := csl.OptionsPrompt("Which editor do you prefer?", []string{pathEditor, "vim", "nano"})
	g.Editor = editor

	prov := csl.OptionsPrompt("What is your default provisioner?", []string{"ansible"})
	if prov == "ansible" {
		g.Ansible = &AnsibleConfiguration{}
		csl.Info("Do you want CHANGEME defaults written to the ansible configuration?")
		editInPlace := csl.YesNoPrompt()
		if editInPlace {
			g.Ansible.VaultPasswordFile = "changeme"
		}
	}

	life := csl.OptionsPrompt("What is your default lifecycle?", []string{"aws"})
	if life == "aws" {
		g.AWS = &AwsConfiguration{}
		csl.Info("Do you want CHANGEME defaults written to the AWS configuration?")
		editInPlace := csl.YesNoPrompt()
		if editInPlace {
			//var required = "REQUIRED"
			var changeme = "CHANGEME"
			g.AWS.InstanceType = changeme
			g.AWS.KeypairName = changeme
			g.AWS.SubnetID = changeme
			g.AWS.Tags = append(g.AWS.Tags, &AwsTag{
				Key:   changeme,
				Value: changeme,
			})
			g.AWS.IamInstanceProfile = changeme
			g.AWS.Region = changeme
			g.AWS.AMI = changeme
			g.AWS.SecurityGroups = []string{changeme}
			g.AWS.SSHUsername = changeme
			g.AWS.SSHPrivateKeyPath = changeme
		}
	}

	return g.Write()
}

func (g *Global) Open() error {
	cmd := exec.Command(g.Editor, configFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (g *Global) Read() error {
	file, err := os.Open(configFile)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, &g)
}

func (g *Global) Write() error {
	bin, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFile, bin, 0755)

}

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		csl.Error("Failed to discover users homedir, can't read or write config file!")
		os.Exit(1)
	}

	configDir = fmt.Sprintf("%s/.config/betwixt", homedir)
	configFile = fmt.Sprintf("%s/betwixt.conf", configDir)
}
