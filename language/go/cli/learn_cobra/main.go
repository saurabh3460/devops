package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	force   bool
	rootCmd = &cobra.Command{
		Use:   "learncode",
		Short: "learn go code via cli",
		Long:  "bla bla bla blaa blaa..",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from root", force)

		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "set to force the action")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
