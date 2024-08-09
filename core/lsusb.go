package core

import (
	"os/exec"
)

func RunLSUSB() (string, error) {
	cmd := exec.Command("lsusb")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
