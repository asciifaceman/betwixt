/*
Copyright © 2022 Charles

*/
package cmd

import (
	"os"

	"github.com/asciifaceman/betwixt/betwixt"
	"github.com/asciifaceman/betwixt/betwixt/csl"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Launch a test cycle from scratch",
	Long: `test launches a betwixt run from scratch,
starting with a clean that destroys any current instance
tied to the local project, then launching it fresh, running
the defined privisioner, and applying tests at the end.`,
	Run: func(cmd *cobra.Command, args []string) {
		csl.Note("Starting test cycle...")

		b, err := betwixt.New()
		if err != nil {
			csl.Error(err.Error())
			os.Exit(1)
		}

		StateID := b.State.GetID()
		if StateID == "" {
			csl.Note("No remote instances exist...")
			err := b.Lifecycle.Launch()
			if err != nil {
				csl.Error(err.Error())
				csl.Info("We really should now loop with a check for remote host state with a failout")
			}
		}

		/*
			b, err := betwixt.Bootstrap()
			if err != nil {
				fmt.Printf("Failed to start betwixt: %v\n", err)
				os.Exit(1)
			}

			b.Lifecycle.Launch()

			// read in global and local config
			// destroy local ansible cache (downloaded roles)
			// destroy remote instance if exists
			// create remote instance
			// run ansible against instance
			// run goss tests
			// report
		*/
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	testCmd.GroupID = workflowGroup.ID
}
