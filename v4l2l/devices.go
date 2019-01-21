package v4l2l

// #include "devices.h"
import "C"

import (
	"log"
	"path/filepath"
)

// Device describes a device file for a v4l2loopback device.
type Device struct {
	Name string // Device name, e.g. 'video1'
	Path string // Device path, e.g. '/dev/video1'
}

// DeviceList describes a list of device.
type DeviceList []Device

// GetDeviceList returns the v4l2loopback device list.
func GetDeviceList() DeviceList {
	return parseDevices()
}

// parseDevices parses the '/dev' directory looking for video devices
// ('video0', 'video1', ...). It checks if a video device is a valid
// video4linux2loopback device and return a list of V4l2lDevice.
func parseDevices() DeviceList {
	matches, err := filepath.Glob("/dev/video*")
	if err != nil {
		log.Fatal(err)
	}

	var devices DeviceList
	for _, dev := range matches {
		if isLoopbackDevice(dev) {
			device := Device{
				Name: filepath.Base(dev),
				Path: dev,
			}
			devices = append(devices, device)
		}
	}

	return devices
}

// isLoopbackDevice tests if a device is a valid video4linux2loopback device.
// Input parameter has to be an absolute path to a device file.
func isLoopbackDevice(device string) bool {
	// Convert to C compatible string type
	cs := C.CString(device)

	result := C.isLoobackDevice(cs)

	return bool(result)
}
