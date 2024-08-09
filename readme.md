# Arctis 9 Battery Level Checker *(arctis9bat)*

## Overview
The **Arctis 9 Battery Level Checker** is a command-line tool designed to solve a common problem
faced by Linux users: the lack of official support for checking the battery level of SteelSeries Arctis
9 headsets.\
Without native Linux software, users are left without an easy way to monitor their
headset's battery status.\
This project was born out of necessity and curiosity, as it's my first
attempt at reverse engineering, where i learned about the world of HID (Human Interface Device)
protocols to retrieve battery information **directly** from the headset.

## Features
- **Battery Level Monitoring**: Query and display the battery level of your SteelSeries Arctis 9 
headset on Linux.
- **Cross-Platform Compatibility**: While primarily aimed at Linux users, the tool is cross-platform and can run on any system where the HIDAPI library is supported.
- **Extendable**: The project is designed with future enhancements in mind, allowing easy addition of support for other features and headsets.

## Installation
To get started, clone the repository and build the project:
```bash
git clone https://github.com/Catfish1210/arctis9bat.git
cd arctis9bat
go build -o arctis9bat
chmod +x arctis9bat
```

## Prerequisites
- **Go**: Ensure you have Go version 1.22.3 or newer installed on your system.
- **HIDAPI**: Depending on your OS, you may need to install HIDAPI. On Debian-based systems,
you can install it with:
```bash
sudo apt-get install libhidapi-libusb0
```

## Dependencies
- **[github.com/karalabe/hid](github.com/karalabe/hid)**

## Usage
After building the project, you can run the tool as follows:
```bash
./arctis9bat
# alternatively without udev rules:
sudo ./arctis9bat
```
Make sure to run the tool with the appropriate permissions, especially on Linux.\
If you don't want to use `sudo`, follow the instructions below to configure udev rules.

### Setting Up udev Rules for Non-root Access (Linux)
To allow non-root users to access the headset, create a udev rule:
>**1. Create a new rule file**:
>>```bash
>>sudo touch /etc/udev/rules.d/99-steelseries.rules
>>```
>
>**2. Add the following lines**:
>>```bash
>>SUBSYSTEM=="usb", ATTR{idVendor}=="1038", ATTR{idProduct}=="12c2", MODE="0666", GROUP="plugdev"
>>SUBSYSTEM=="usb", ATTR{idVendor}=="1038", ATTR{idProduct}=="12c4", MODE="0666", GROUP="plugdev"
>>```
>>>**ATTR{idVendor}** and **ATTR{idProduct}** should match your device's vendor and product ID. Adjust these if needed.\
>>>These can be verified by using `lsusb`:
>>>```bash
>>>lsusb
>>># Where similar lines should be present with matching idVendor and idProduct
>>>#Bus 000 Device 019: ID 1038:12c2 SteelSeries ApS SteelSeries Arctis 9
>>>#Bus 000 Device 020: ID 1038:12c4 SteelSeries ApS SteelSeries Arctis 9
>>>```
>>>MODE="0666" sets read and write permissions for everyone.\
>>>GROUP="plugdev" is the group that should have access. Ensure your user is part of this group or use GROUP="users".
>
>**3. Reload udev rules**:
>>```bash
>>sudo udevadm control --reload-rules
>>sudo udevadm trigger
>>```
>**4. Some devices might require unpluging and repluging your headset.**

## Contributing
If you have a different headset and would like to see it supported, feel free to open an issue or
submit a pull request. This project is designed with extendability in mind, and I'm happy to
collaborate to add support for more features/devices.

## Acknowledgements
Special thanks to the open-source community and the creators of HIDAPI, whose work made this
project possible.

## Licence
This project is licensed under the MIT Licence. See the [LICENCE](./LICENCE) file for details.<br><br>

---
Feel free to reach out if you have any questions or suggestions. I'm excited to see how this project
can grow and help others in the community!