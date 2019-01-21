#include "devices.h"

#include <stdio.h>

int main(int argc, char* argv[])
{
    char* device= "/dev/video0";

    if (argc >= 2) {
        device = argv[1];
    }

    bool valid = isLoobackDevice(device);

    fprintf(stdout, "Device '%s' is loopback: '%s'\n",
        device, valid ? "true" : "false");

    return 0;
}