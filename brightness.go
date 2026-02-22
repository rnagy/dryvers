package dryvers

import (
	"errors"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type brightType int

const (
	noBacklight brightType = iota
	xbacklight
	brightnessctl
)

// Brightness is a type that handles querying screen brightness.
// It supports brightnessctl and xbacklight supported devices.
type Brightness struct {
	mode brightType
}

// Get returns the current screen brightness (between 0 and 1) or
// an error if there was a problem.
func (b *Brightness) Get() (float64, error) {
	switch b.mode {
	case brightnessctl:
		out, err := exec.Command("brightnessctl", "get").Output()
		if err != nil {
			return 0, err
		}
		maxOut, _ := exec.Command("brightnessctl", "max").Output()
		val, err := strconv.Atoi(strings.TrimSpace(string(out)))
		if err != nil {
			return 0, err
		}
		max, _ := strconv.Atoi(strings.TrimSpace(string(maxOut)))
		return float64(val) / float64(max), nil
	default:
		out, err := exec.Command("xbacklight").Output()
		if err != nil {
			return 0, err
		}

		if strings.TrimSpace(string(out)) == "" {
			return 0, errors.New("no back-lit screens found")
		}
		ret, err := strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
		if err != nil {
			return 0, err
		}
		return ret / 100, nil
	}
}

// Set attempts to set the screen brightness - 1 for maximum and 0 for off.
// An error is returned in the case of a problem.
func (b *Brightness) Set(value float64) error {
	if value < 0 {
		value = 0
	} else if value > 1 {
		value = 1
	}

	percentStr := strconv.Itoa(int(value * 100))
	switch b.mode {
	case brightnessctl:
		return exec.Command("brightnessctl", "set", percentStr+"%").Run()
	case xbacklight:
		return exec.Command("xbacklight", "-set", percentStr).Run()
	default:
		return nil
	}
}

func brightnessMode() brightType {
	cmd := exec.Command("xbacklight")
	err := cmd.Run()
	if err != nil || cmd.ProcessState.ExitCode() != 0 {
		err = exec.Command("brightnessctl").Run()
		if err != nil {
			log.Println("Could not launch xbacklight or brightnessctl", err)
			return noBacklight
		} else {
			return brightnessctl
		}
	}
	return xbacklight
}

func NewBrightness() *Brightness {
	return &Brightness{mode: brightnessMode()}
}
