package main

// import "github.com/saurabh3460/cli_cobra_1/cmd"
// func main() {
// 	cmd.Execute()
// }

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dadjoke",
		Short: "Get some dad jokes",
		Long:  "Now days we don't listen dad's jokes, so use this cli to do so!!",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello there!", args)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
