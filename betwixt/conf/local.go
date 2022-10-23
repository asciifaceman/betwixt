package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var validProvisioners = []string{"ansible"}
var validLifecycles = []string{"aws"}

func isValidProvisioner(prov string) bool {
	for _, val := range validProvisioners {
		if val == prov {
			return true
		}
	}
	return false
}

func isValidLifecycle(lc string) bool {
	for _, val := range validLifecycles {
		if val == lc {
			return true
		}
	}
	return false
}

func InitLocal() error {
	l := &Local{}

	var provinput string
	var lcinput string

	fmt.Println("Creating skeleton config for local project.")
	fmt.Println("Which provisioner will this project use?")

	for {
		fmt.Printf("Supported options include (%s): ", strings.Join(validProvisioners, ","))
		fmt.Scanln(&provinput)
		lowered := strings.ToLower(provinput)

		if isValidProvisioner(lowered) {
			l.Provisioner = lowered
			break
		} else {
			continue
		}
	}

	fmt.Println("Which Lifecycle will this project use? (typically your cloud provider)")

	for {
		fmt.Printf("Supported options include (%s): ", strings.Join(validLifecycles, ","))
		fmt.Scanln(&lcinput)
		lowered := strings.ToLower(lcinput)

		if isValidLifecycle(lowered) {
			l.Lifecycle = lowered
			break
		} else {
			continue
		}
	}

	err := l.Write()
	if err != nil {
		return err
	}

	fmt.Println("Successfully created ./betwix")
	return nil
}

// Local represents a project-local configuration
type Local struct {
	Provisioner string                     `json:"provisioner"`
	Lifecycle   string                     `json:"lifecycle"`
	AWS         *LocalAwsConfiguration     `json:"aws,omitempty"`
	Ansible     *LocalAnsibleConfiguration `json:"ansible,omitempty"`
}

func (l *Local) Write() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	localConfigPath := fmt.Sprintf("%s/betwixt", wd)

	bin, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(localConfigPath, bin, 0755)

}

func (l *Local) Read(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, l)

}

type State struct {
	Filepath string
	AWS      *AwsState `json:"aws,omitempty"`
}

func (s *State) Read() error {
	file, err := os.Open(s.Filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &s)
	return err
}

func (s *State) Write() error {
	bin, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(s.Filepath, bin, 0755)

	return err
}
