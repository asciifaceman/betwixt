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
	Editor string            `json:"editor"`
	AWS    *AwsConfiguration `json:"aws"`
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
