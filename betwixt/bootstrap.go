package betwixt

import (
	"fmt"
	"os"
	"strings"

	"github.com/asciifaceman/betwixt/betwixt/conf"
)

func Bootstrap() (*Betwixt, error) {

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	localConfigPath := fmt.Sprintf("%s/betwixt", wd)

	localConfig := &conf.Local{}

	f, err := os.Stat(localConfigPath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Do you want to create?")
		var createPrompt string
		for {
			fmt.Print("(y/n) or ctl+c to exit: ")
			fmt.Scanln(&createPrompt)

			confirm := strings.ToLower(createPrompt)

			if confirm == "" || confirm != "y" && confirm != "n" {
				fmt.Println("Input must be y/Y or n/N")
				continue
			}

			if confirm == "y" {
				err = conf.InitLocal()
				return nil, err
			}

			return nil, fmt.Errorf("File does not exist and did not create it. See betwixt init.")
		}

	}
	if f.IsDir() {
		return nil, fmt.Errorf("./betwixt is a directory, not a config file")
	}

	err = localConfig.Read(localConfigPath)
	if err != nil {
		return nil, err
	}

	globalConfig := &conf.Global{}
	err = globalConfig.Read()
	if err != nil {
		return nil, err
	}

	b := New(withLifecycler(localConfig.Lifecycle), WithProvisioner(localConfig.Provisioner))
	err = b.IngestConfigs(localConfig, globalConfig)
	if err != nil {
		return nil, err
	}

	return b, nil
}
