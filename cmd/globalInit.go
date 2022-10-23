/*
Copyright © 2022 Charles

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initGlobalCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize global config",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("globalinit called")
		/*
			var editor string
			fmt.Println("Preferred editor: This should be something in your PATH to open config files with (ex. code, vim)")
			fmt.Print("Editor: ")
			fmt.Scanln(&editor)

			sampleTag := &conf.AwsTag{
				Key:   "managedBy",
				Value: "betwixt",
			}

			g := &conf.Global{
				Editor: editor,
				AWS: &conf.AwsConfiguration{
					Tags: []*conf.AwsTag{
						sampleTag,
					},
					SecurityGroups: make([]string, 0),
				},
			}
			err := g.Init()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		*/
	},
}

func init() {
	configCmd.AddCommand(initGlobalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
