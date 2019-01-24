package main

import (
	"flag"
	"os"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/fsnotify/fsnotify"
	"github.com/mpreu/k8s-device-plugin-v4l2loopback/v4l2l"
	api "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
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

	log.Println("Starting OS signals watcher")
	sig := newOSSignalWatcher(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	log.Println("Starting filesystem event watcher")
	watcher, err := newFSWatcher(api.DevicePluginPath)

	if err != nil {
		log.Errorf("Could not create filesystem watcher: %v", err)
		return
	}

	defer watcher.Close()

	log.Println("Starting device plugin server")
	devicePlugin := NewV4l2lDevicePlugin()

	err = devicePlugin.Serve()

	if err != nil {
		log.Errorf("Plugin server error: %v", err)
		return
	}

	for {
		// Wait for channels
		select {
		// Termination signals
		case s := <-sig:
			log.Debugf("Termination signal received: %v", s)
			devicePlugin.StopServer()
		// Filesystem events
		case event := <-watcher.Events:
			if event.Name == api.KubeletSocket && event.Op&fsnotify.Create == fsnotify.Create {
				log.Infof("fsnotify: %s created", api.KubeletSocket)
				devicePlugin := NewV4l2lDevicePlugin()
				err = devicePlugin.Serve()

				if err != nil {
					log.Errorf("Plugin server error: %v", err)
					return
				}
			}
			if event.Name == api.KubeletSocket && event.Op&fsnotify.Remove == fsnotify.Remove {
				log.Infof("fsnotify: %s removed", pluginSocket)
				devicePlugin.StopServer()
			}

		}
	}
}
