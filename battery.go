package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// Battery provides the information for a battery
type Battery struct {
	Consumption      float64 // (watts) current power flowing into/out of the battery
	Name             string
	Filepath         string
	Percentage       float64       // battery charge percentage
	PercentageDesign float64       // absolute battery charge percentage
	Remaining        time.Duration // remaining time for charging or discharging
	Status           string
}

// New instantiates a Battery object for the passed in string representing a system's battery
func New(name string) (*Battery, error) {
	battery := &Battery{
		Name:     name,
		Filepath: path.Join("/sys", "class", "power_supply", name),
	}
	// make sure the battery exists
	if _, err := os.Stat(battery.Filepath); os.IsNotExist(err) {
		return battery, fmt.Errorf("Battery %s does not exist, please specify a battery name with -name", name)
	}

	return battery, nil
}

func getFloat(filepath string) (value float64, err error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return value, err
	}
	return strconv.ParseFloat(strings.TrimSpace(string(file)), 64)
}

func getInt(filepath string) (value int, err error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return value, err
	}
	return strconv.Atoi(strings.TrimSpace(string(file)))
}

// GetStatus updates the Battery object with the current state of the battery
func (battery *Battery) GetStatus() (err error) {
	// get status
	file, err := ioutil.ReadFile(path.Join(battery.Filepath, "status"))
	if err != nil {
		return err
	}
	battery.Status = strings.TrimSpace(string(file))

	// decide how we will calculate
	if _, err := os.Stat(path.Join(battery.Filepath, "power_now")); err == nil {
		// get consumption
		battery.Consumption, err = getFloat(path.Join(battery.Filepath, "power_now"))
		if err != nil {
			return err
		}

		// get percentages
		energyNow, err := getFloat(path.Join(battery.Filepath, "energy_now"))
		if err != nil {
			return err
		}
		energyFull, err := getFloat(path.Join(battery.Filepath, "energy_full"))
		if err != nil {
			return err
		}
		energyFullDesign, err := getFloat(path.Join(battery.Filepath, "energy_full_design"))
		if err != nil {
			return err
		}
		battery.Percentage = (energyNow / energyFull) * 100
		battery.PercentageDesign = (energyNow / energyFullDesign) * 100

		// get remaining
		powerNow, err := getFloat(path.Join(battery.Filepath, "power_now"))
		if err != nil {
			return err
		}
		if battery.Status == "Discharging" {
			battery.Remaining = time.Duration(int((energyNow/powerNow)*(60*60))) * time.Second
		} else {
			battery.Remaining = time.Duration(int(((energyFull-energyNow)/powerNow)*(60*60))) * time.Second
		}
	} else {
		// get consumption
		voltage, err := getFloat(path.Join(battery.Filepath, "voltage_now"))
		if err != nil {
			return err
		}
		current, err := getFloat(path.Join(battery.Filepath, "current_now"))
		if err != nil {
			return err
		}
		battery.Consumption = voltage * current

		// get percentages
		chargeNow, err := getFloat(path.Join(battery.Filepath, "charge_now"))
		if err != nil {
			return err
		}
		chargeFull, err := getFloat(path.Join(battery.Filepath, "charge_full"))
		if err != nil {
			return err
		}
		chargeFullDesign, err := getFloat(path.Join(battery.Filepath, "charge_full_design"))
		if err != nil {
			return err
		}
		battery.Percentage = (chargeNow / chargeFull) * 100
		battery.PercentageDesign = (chargeNow / chargeFullDesign) * 100

		// get remaining
		if battery.Status == "Discharging" {
			battery.Remaining = time.Duration(int((chargeNow/current)*(60*60))) * time.Second
		} else {
			battery.Remaining = time.Duration(int(((chargeFull-chargeNow)/current)*(60*60))) * time.Second
		}
	}

	if battery.Consumption > 0.1 && battery.Percentage < 99.9 {
		if battery.Status != "Discharging" {
			battery.Status = "Charging"
		}
	} else if battery.Consumption == 0 && battery.Percentage == 0.00 {
		battery.Status = "Depleted"
	} else {
		battery.Status = "Charged"
		battery.Remaining = time.Duration(0) * time.Second
	}

	return nil
}

func (battery *Battery) String() string {
	batteryString := `Battery %s
Status:            %s
Percentage:        %f
Percentage Design: %f
Remaining:         %s
Consumption:       %f`

	return fmt.Sprintf(batteryString, battery.Name, battery.Status, battery.Percentage, battery.PercentageDesign, battery.Remaining.String(), battery.Consumption)
}
