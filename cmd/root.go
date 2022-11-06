/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	workflowGroup = &cobra.Group{
		ID:    "workflow",
		Title: "Workflow",
	}
	supportGroup = &cobra.Group{
		ID:    "support",
		Title: "Support",
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "betwixt",
	Short: "Applied cloud test flow for configuration management",
	Long: `
______        _              _        _    
| ___ \      | |            (_)      | |   
| |_/ /  ___ | |_ __      __ _ __  __| |_  
| ___ \ / _ \| __|\ \ /\ / /| |\ \/ /| __| 
| |_/ /|  __/| |_  \ V  V / | | >  < | |_  
\____/  \___| \__|  \_/\_/  |_|/_/\_\ \__| 

	Betwixt is a CLI application that allows one
to launch ec2 instances, apply configuration management,
and then run tests to support the development flow of automation
with a single utility purpose built to this workflow.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.betwixt.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddGroup(workflowGroup)
	rootCmd.AddGroup(supportGroup)

}
