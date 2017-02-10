package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/wwkeyboard/goorg"
)

// todoCmd represents the todo command
var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Run,
}

func init() {
	RootCmd.AddCommand(todoCmd)

	//	orgFile, err := RootCmd.PersistentFlags().GetString("file")
	//	if err != nil {
	//		fmt.Println("Must set file to operate on")
	//		os.Exit(1)
	//	}

	// todoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// todoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Run(cmd *cobra.Command, args []string) {
	org, err := goorg.FromFile(orgFile)
	if err != nil {
		fmt.Println("Couldn't read org-mode file")
	}

	var ibox *goorg.Headline

	for _, h := range org.Headlines {
		if h.Title == modHeading {
			ibox = h
		}
	}

	if ibox == nil {
		fmt.Println("couldn't find Headline")
		os.Exit(1)
	}

	fmt.Println("found it")
}
