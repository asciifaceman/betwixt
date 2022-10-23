package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	configPath string
	configFile string
)

// Global represents a system wide configuration for
// betwixt which informs new local project initializations
// as well as your sane defaults
type Global struct {
	Editor  string                `json:"editor"`
	AWS     *AwsConfiguration     `json:"aws"`
	Ansible *AnsibleConfiguration `json:"ansible"`
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

// Init sets up the global config path and
// writes the current global to it
func (g *Global) Init() error {
	fmt.Printf("Creating %s and writing default config...\n", configPath)
	err := os.MkdirAll(configPath, 0755)
	if err != nil {
		return err
	}

	err = g.Write()
	if err != nil {
		return err
	}

	var confirmEdit string
	fmt.Println("Do you want to edit your global config now? (it will be mostly empty)")

	for {
		fmt.Print("(y/n) or ctl+c to exit: ")
		fmt.Scanln(&confirmEdit)

		confirm := strings.ToLower(confirmEdit)

		if confirm == "" || confirm != "y" && confirm != "n" {
			fmt.Println("Input must be y/Y or n/N")
			continue
		}

		if confirm == "y" {
			g.Open()
		}

		break

	}

	fmt.Printf("Global config created at %s", configFile)

	return err
}

type SSHConfig struct {
	Username       string `json:"username"`
	PrivateKeyPath string `json:"privateKeyPath"`
}

func init() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to discover users homedir, can't read or write global config")
		os.Exit(1)
	}

	configPath = fmt.Sprintf("%s/.config/betwixt", homedir)
	configFile = fmt.Sprintf("%s/betwixt.conf", configPath)
}
