#include "devices.h"

#include <fcntl.h>
#include <linux/videodev2.h>
#include <sys/ioctl.h>
#include <sys/stat.h>

#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#include <unistd.h>

static int openDevice(char deviceName[])
{
    // Open device
    int fd = open(deviceName, O_RDONLY /* required */ | O_NONBLOCK, 0);

    if (-1 == fd) {
        fprintf(stderr, "Cannot open '%s': %d, %s\n",
            deviceName, errno, strerror(errno));
        exit(EXIT_FAILURE);
    }

    return fd;
}

bool testDevice(char deviceName[])
{
    struct v4l2_capability cap;

    // Open device
    int fd = openDevice(deviceName);
    if (-1 == fd) {
        fprintf(stderr, "Could not open '%s': %d, %s\n",
            deviceName, errno, strerror(errno));
        exit(EXIT_FAILURE);
    }

    int err = ioctl(fd, VIDIOC_QUERYCAP, &cap);
    close(fd);

    if (-1 == err) {
        fprintf(stderr, "Could not open '%s': %d, %s\n",
            deviceName, errno, strerror(errno));
        exit(EXIT_FAILURE);
    }
    bool valid = false;
    char correctDriverName[] = "v4l2 loopback";
    if (0 != strcmp(cap.driver, correctDriverName)) {
        fprintf(stderr, "Device driver is '%s' not '%s'\n",
            cap.driver, correctDriverName);
        valid = false;
    }
    valid = true;

    return valid;
}

bool isLoobackDevice(char deviceName[])
{
    struct stat st;

    if (-1 == stat(deviceName, &st)) {
        fprintf(stderr, "Could not process '%s': %d, %s\n",
            deviceName, errno, strerror(errno));
        exit(EXIT_FAILURE);
    }

    // Test if character device
    if (!S_ISCHR(st.st_mode)) {
        fprintf(stderr, "%s is no character device\n", deviceName);
        exit(EXIT_FAILURE);
    }

    // Test if device is valid
    bool valid = false;
    valid = testDevice(deviceName);

    return valid;
}