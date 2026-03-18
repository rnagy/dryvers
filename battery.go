package dryvers

import (
    "github.com/distatus/battery"
)

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

// NewBattery creates a new instance of the batter driver.
func NewBattery() *Battery {
	return &Battery{}
}

func (b *Battery) PluggedIn() (bool, error) {
	batteries, err := battery.GetAll()
	if err != nil {
		return true, err // assume power if no battery info
	}
	for _, battery := range batteries {
		// if any battery is discharging, we are not plugged in
		if (battery.State.String() == "Discharging") {
			return false, nil
		}
	}
	return true, nil
}

// Get returns the current battery level (between 0 and 1) or
// an error if there was a problem.
func (b *Battery) Get() (float64, error) {
	batteries, err := battery.GetAll()
	var ctotal, ftotal float64
	if err != nil {
		return 0, nil
	}
	for _, battery := range batteries {
		ctotal += battery.Current
		ftotal += battery.Full
	}
	return ctotal / ftotal, nil
}
