package main

import (
	"os"
	"os/signal"

	"github.com/fsnotify/fsnotify"
)

func newOSSignalWatcher(sig ...os.Signal) chan os.Signal {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, sig...)

	return channel
}

func newFSWatcher(files ...string) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		return nil, err
	}

	for _, f := range files {
		err = watcher.Add(f)
		if err != nil {
			watcher.Close()
			return nil, err
		}
	}

	return watcher, nil
}
