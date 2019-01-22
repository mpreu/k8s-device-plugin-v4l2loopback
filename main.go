package main

import (
	"flag"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/mpreu/k8s-device-plugin-v4l2loopback/v4l2l"
)

func main() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagLogLevel := flag.String("log-level", "info", "Define the logging level: info, debug.")
	flag.Parse()

	switch *flagLogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	}

	log.Println("Starting k8s-device-plugin-v4l2loopback")
	log.Println("Searching for devices ...")
	devices := v4l2l.GetDeviceList()

	if len(devices) == 0 {
		log.Println("No devices found")
		return
	}

	log.Debugf("Devices found: %v", devices)
}
