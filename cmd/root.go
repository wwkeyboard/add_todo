package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wwkeyboard/goorg"
)

var orgFile string
var modHeading string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "add_todo",
	Short: "Adds todo items to an org-mode file",
	Long: `Parses an org-mode file and adjusts is
This is meant to have an interaction similar to taskwarrior.
`,
	Run: run,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&orgFile, "file", "", "org file to operate on")
	RootCmd.PersistentFlags().StringVar(&modHeading, "heading", "INBOX", "Heading to add to")

	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}

func run(cmd *cobra.Command, args []string) {
	org, err := goorg.FromFile(orgFile)
	if err != nil {
		fmt.Println("Couldn't read org-mode file")
	}

	fmt.Printf("Read %v toplevel headlines\n", len(org.Headlines))
}
