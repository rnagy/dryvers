package dryvers

import "os"

type Battery struct{}

// IsLow returns true if the battery is low (less than 10%) or
// an error if there was a problem.
func (b *Battery) IsLow() (bool, error) {
	v, err := b.Get()
	if err != nil {
		return false, err
	}

	return v < 0.1, nil
}

func pickChargeOrEnergy() (string, string) {
	_, err := os.Stat("/sys/class/power_supply/BAT0/charge_now")
	if err != nil {
		return "/sys/class/power_supply/BAT0/energy_now", "/sys/class/power_supply/BAT0/energy_full"
	}
	return "/sys/class/power_supply/BAT0/charge_now", "/sys/class/power_supply/BAT0/charge_full"
}

// NewBattery creates a new instance of the batter driver.
func NewBattery() *Battery {
	return &Battery{}
}
