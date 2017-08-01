package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/util"
	_ "github.com/uzem/inspeqtor-patched/notifiers"
)

const (
	RuleFailed    inspeqtor.EventType = "RuleFailed"
	RuleRecovered                     = "RuleRecovered"
)

var version = "undefined"

var (
	configDir       string
	socketPath      string
	logLevel        string
	testConfig      bool
	testAlertRoutes bool

	vrsn bool
)

func init() {
	inspeqtor.Name = "Inspeqtor Patched"
	inspeqtor.Events = append(inspeqtor.Events, RuleFailed, RuleRecovered)

	flag.StringVar(&configDir, "c", "/etc/inspeqtor", "Location of inspeqtor config files")
	flag.StringVar(&socketPath, "s", "/var/run/inspeqtor.sock", "Connects to inspeqtor using the socket provided")
	flag.StringVar(&logLevel, "l", "info", "Logging level (warn, info, debug, verbose)")
	flag.BoolVar(&testConfig, "tc", false, "Test configuration and exit")
	flag.BoolVar(&testAlertRoutes, "ta", false, "Test alert routes and exit")
	flag.BoolVar(&vrsn, "v", false, "Print version and exit")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s - %s\n", inspeqtor.Name, version)
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrsn {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	util.SetLogLevel(logLevel)
	setupLogging()
}

func main() {
	inq, err := inspeqtor.New(configDir, socketPath)
	if err != nil {
		log.Fatalln(err)
	}

	if err := inq.Parse(); err != nil {
		log.Fatalln(err)
	}

	if testConfig {
		util.Info("Configuration parsed ok.")
		os.Exit(0)
	} else if testAlertRoutes {
		inq.TestAlertRoutes()
	} else {
		inq.Start()
		inspeqtor.HandleSignals()
	}
}

func setupLogging() {
	log.SetPrefix("inspeqtor: ")
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}
