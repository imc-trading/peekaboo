package sensors

import (
	"os"
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type Sensors []Sensor

type Sensor struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Status string `json:"status"`
	Entity string `json:"entity"`
	Value  string `json:"value"`
}

func Get() (Sensors, error) {
	// Create sdr dump, too speed up consequtive queries
	sdr := "/tmp/peekaboo.sdr"
	if _, err := os.Stat(sdr); os.IsNotExist(err) {
		_, err := parse.Exec("ipmitool", []string{"sdr", "dump", sdr})
		if err != nil {
			return Sensors{}, err
		}
	}

	// Query for sdr (sensor data records)
	o, err := parse.Exec("ipmitool", []string{"sdr", "-S", sdr, "elist", "all"})
	if err != nil {
		return Sensors{}, err
	}

	s := Sensors{}
	for _, line := range strings.Split(o, "\n") {
		a := strings.Split(line, "|")
		if len(a) < 5 {
			continue
		}

		s = append(s, Sensor{
			Name:   strings.TrimSpace(a[0]),
			ID:     strings.TrimSpace(a[1]),
			Status: strings.TrimSpace(a[2]),
			Entity: strings.TrimSpace(a[3]),
			Value:  strings.TrimSpace(a[4]),
		})
	}

	return s, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
