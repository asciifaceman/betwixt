/*
Copyright © 2022 Charles

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type TestA struct {
	Thing  string
	Thing3 []string
	Thing4 []*TestB
}

type TestB struct {
	Name string
}

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Run provisioner and tests without changing infrastructure",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apply called")
		// run ansible against remote host
		// run goss tests against remote host

	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	applyCmd.GroupID = workflowGroup.ID
}
