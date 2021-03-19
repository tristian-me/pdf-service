package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"pdf-service/utils"
	"pdf-service/web"
)

// CliArgs is a simple struct for the CLI Args (flags)
type CliArgs struct {
	Port    int
	TempDir string
}

// DefaultTempDir is the default temporary directory
const (
	DefaultTempDir         = "/tmp/pdf-service"
	DefaultSecondsInterval = 5
)

func main() {
	// Check Chromium path
	checkChromium()

	// Fetch flags
	args := CliArgs{}
	flag.IntVar(&args.Port, "port", 8765, "Command line arguments")
	flag.StringVar(&args.TempDir, "temp-dir", DefaultTempDir, "The pemparary directory")
	flag.Parse()

	fmt.Println("Starting PDF Service")
	fmt.Println("Running on port: ", args.Port)

	err := setupDirectories(args)
	if err != nil {
		log.Fatal("Cannot create directory, maybe not writable?")
	}

	// Is the event tool for removing
	go utils.GarbageCollection(args.TempDir, DefaultSecondsInterval)

	web.RunServer(args.Port, args.TempDir)
}

func setupDirectories(args CliArgs) error {
	var err error

	if args.TempDir == DefaultTempDir {
		err = os.MkdirAll(DefaultTempDir, 0777)
	} else {
		err = os.MkdirAll(args.TempDir, 0777)
	}

	return err
}

func checkChromium() {
	_, err := exec.LookPath("chromium")
	if err != nil {
		log.Fatal(err)
	}
}
