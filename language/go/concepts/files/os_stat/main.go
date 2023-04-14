package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

var Version = "development"
var GitCommit string
var GitBranch string
var BuildDate string

func main() {
	fmt.Println("Version:\t", Version, os.Args[0])
	fmt.Printf("BUILD-INFO: commit:%v, branch: %v, date: %v, version: %v\n",
		GitCommit, GitBranch, BuildDate, Version)
	if finfo, err := os.Stat(os.Args[0]); err == nil {
		stat := finfo.Sys().(*syscall.Stat_t)
		fmt.Printf("Build Time: %v", time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec)))
	}
}
