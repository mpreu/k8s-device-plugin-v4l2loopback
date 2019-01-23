package main

import (
	"os"
	"os/signal"
)

func newOSSignalWatcher(sig ...os.Signal) chan os.Signal {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, sig...)

	return channel
}
