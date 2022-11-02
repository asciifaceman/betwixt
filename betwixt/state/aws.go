package state

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	awsStateFileName = "aws.json"
)

type AwsState struct {
	ProjectDirectory string `json:"-"`
	InstanceID       string `json:"instanceId"`
}

// Init sets up the state handler
func (a *AwsState) Init(wd string) error {
	a.ProjectDirectory = wd
	err := os.Mkdir(a.ProjectDirectory, 0755)
	return err
}

// Load reads in an existing state file for AWS
func (a *AwsState) Load() error {
	stateFile := fmt.Sprintf("%s/%s", a.ProjectDirectory, awsStateFileName)

	file, err := os.Open(stateFile)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, a)
}

// Write writes out the state file for AWS
func (a *AwsState) Write() error {
	stateFile := fmt.Sprintf("%s/%s", a.ProjectDirectory, awsStateFileName)

	bin, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(stateFile, bin, 0755)
}

// GetID returns the currently stored instance id
func (a *AwsState) GetID() string {
	return a.InstanceID
}

// SetID stores a given instance ID and writes it to the state file
func (a *AwsState) SetID(id string) error {
	a.InstanceID = id
	return a.Write()
}
