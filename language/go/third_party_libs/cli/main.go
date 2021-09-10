package main

import (
	"fmt"
	"math/rand"

	"runtime"

	"time"

	term "github.com/gesquive/cli"
	"github.com/spf13/cobra"
)

const (
	CLIName = "runner"
)

type RunnerFlags struct {
	debug bool
	user  string
}

var option RunnerFlags

func main() {
	var rootCmd = &cobra.Command{
		Use:   CLIName,
		Short: "razorci build runner",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	var runCmd = &cobra.Command{
		Use: "run",
		Run: func(cmd *cobra.Command, args []string) {
			runtime.GOMAXPROCS(runtime.NumCPU())
			rand.Seed(time.Now().UnixNano())
			if option.debug {
				term.SetPrintLevel(term.LevelDebug)
			}
			fmt.Println(args)
			//var err error
			if len(args) < 1 {
				runSingleJob("", option)
			} else {
				runSingleJob(args[0], option)
			}

		},
	}

	var detailCmd = &cobra.Command{
		Use: "detail",
		Short: "use to get detail on topic",
		Long: "use detail to get detail description of a topics",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("detail:", args)
			fmt.Println(cmd.Flag("name").Value)
			fmt.Println(cmd.Flag("get").Value)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&option.debug, "debug", "v", false, "debug")

	runCmd.Flags().String("command", "ls", "Pass command to run")

	runCmd.AddCommand(detailCmd)
	rootCmd.AddCommand(runCmd)


	rootCmd.Execute()
}

func runSingleJob(path string, opts RunnerFlags) {
	fmt.Println(path, opts)
}