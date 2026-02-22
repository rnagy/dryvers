//go:build openbsd || freebsd || netbsd
// +build openbsd freebsd netbsd

package dryvers

import "syscall"

func (b *battery) PluggedIn() (bool, error) {
	val, err := syscall.Sysctl("hw.acpi.acline")
	if err != nil {
		return true, err
	}

	return val[0] == 1, nil
}

// Get returns the current battery level (between 0 and 1) or
// an error if there was a problem.
func (b *battery) Get() (float64, error) {
	val, err := syscall.Sysctl("hw.acpi.battery.life")
	if err != nil {
		return 0, err
	}

	percent := int(val[0])
	if percent == 0 { // avoid 0/100 below
		return 0, nil
	}

	return float64(percent) / 100, nil
}
