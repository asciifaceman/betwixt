package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/asciifaceman/betwixt/betwixt/csl"
)

var (
	provisioners = []string{"ansible"}
	lifecycles   = []string{"aws"}
)

func NewLocal(path string) *Local {
	l := &Local{
		Filename: fmt.Sprintf("%s/betwixt", path),
	}

	return l
}

// Local represents a project-local configuration
type Local struct {
	Filename    string            `json:"-"`
	Provisioner string            `json:"provisioner"`
	Lifecycle   string            `json:"lifecycle"`
	AWS         *AwsConfiguration `json:"aws,omitempty"`
}

func (l *Local) Init() error {
	l.Provisioner = csl.OptionsPrompt("Which provisioner will this project use?", provisioners)
	l.Lifecycle = csl.OptionsPrompt("Which Lifecycle will this project use?", lifecycles)

	switch l.Lifecycle {
	case "aws":
		l.AWS = NewAWSConfig()
	default:
	}
	return l.Write()
}

func (l *Local) Write() error {
	bin, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(l.Filename, bin, 0755)

}

func (l *Local) Read() error {
	file, err := os.Open(l.Filename)
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
