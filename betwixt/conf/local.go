package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/asciifaceman/betwixt/betwixt/csl"
)

var (
	provisioners = []string{"ansible"}
	lifecycles   = []string{"aws"}
)

func NewLocal(path string) *Local {
	l := &Local{
		Filename: fmt.Sprintf("%s/betwixtfile", path),
	}

	return l
}

// Need to implement an "architecture" parameter that allows layout definition
/*** ex.
"architecture": {
  "a": {
    "count": 1,
    "variables": {
      "target": "read"
    }
  },
  "b": {
    "count": 1,
    "variables": {
      "target": "write"
    }
  }
}
**/

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

	if l.Lifecycle == "aws" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		dname := strings.Split(wd, "/")
		l.AWS = &AwsConfiguration{}
		l.AWS.Tags = append(l.AWS.Tags, &AwsTag{
			Key:   "Name",
			Value: fmt.Sprintf("%s-betwixt", dname[len(dname)-1]),
		})
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
