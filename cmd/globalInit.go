/*
Copyright © 2022 Charles

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/asciifaceman/betwixt/betwixt/conf"
	"github.com/asciifaceman/betwixt/betwixt/csl"
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
		csl.Note("Initializing global config, this only needs done once for a user...")
		g := conf.Global{}
		err := g.Init()
		if err != nil {
			csl.Error(err.Error())
			os.Exit(1)
		}
		csl.Note(fmt.Sprintf("Config generated. You can find it in [%s]", g.GetFilename()))
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
