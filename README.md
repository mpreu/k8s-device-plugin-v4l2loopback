# v4l2loopback Device Plugin for Kubernetes

## Build Status
`master`:   [![Build Status](https://travis-ci.org/mpreu/k8s-device-plugin-v4l2loopback.svg?branch=master)](https://travis-ci.org/mpreu/k8s-device-plugin-v4l2loopback)

`develop`:  [![Build Status](https://travis-ci.org/mpreu/k8s-device-plugin-v4l2loopback.svg?branch=develop)](https://travis-ci.org/mpreu/k8s-device-plugin-v4l2loopback)

## Intro
`v4l2loopback` is a kernel module which allows you to create "virtual video devices" based on the `video4linux` API. The module allows to consume video input from applications as usual while the content itself is provided by other applications and not from real camera devices.

The module is very useful in the context of simulations (e.g. virtual electrical control units in cars) where you want to provide simulated data to simulated hardware devices.

## Dependencies on the Host
To be able to consume the `v4l2loopback` devices the corresponding kernel module has to be installed on Kubernetes nodes. While the kernel module is provided by some distributions (at least Debian) on most platforms it has to be build manually. See the guide in the [official repository](https://github.com/umlaeute/v4l2loopback).
