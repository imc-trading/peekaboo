package ipmi

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type IPMI struct {
	IpmitoolInstalled bool     `json:"ipmitoolInstalled"`
	IpmitoolVersion   string   `json:"ipmitoolVersion"`
	Fans              Fans     `json:"fans"`
	InletTempDegrC    *int     `json:"inletTempDegrC,omitempty"`
	ExhaustTempDegrC  *int     `json:"exhaustTempDegrC,omitempty"`
	Current1          *float32 `json:"current1Amps,omitempty"`
	Current2          *float32 `json:"current2Amps,omitempty"`
	Voltage1          *int     `json:"voltage1Volts,omitempty"`
	Voltage2          *int     `json:"voltage2Volts,omitempty"`
	PowerCons         *int     `json:"powerConsWatts,omitempty"`
}

type Fans []Fan

type Fan struct {
	Name     string `json:"name"`
	SpeedRPM int    `json:"speedRPM"`
}

func strToIntPtr(m map[string]string, f string) (*int, error) {
	if v, ok := m[f]; ok {
		arr := strings.Split(v, " ")
		i, err := strconv.Atoi(arr[0])
		if err != nil {
			return nil, fmt.Errorf("failed parsing field: %s error: %s", f, err.Error())
		}

		return &i, nil
	}
	return nil, nil
}

func strToFloat32Ptr(m map[string]string, f string) (*float32, error) {
	if v, ok := m[f]; ok {
		arr := strings.Split(v, " ")
		f64, err := strconv.ParseFloat(arr[0], 64)
		f32 := float32(f64)
		if err != nil {
			return nil, fmt.Errorf("failed parsing field: %s error: %s", f, err.Error())
		}

		return &f32, nil
	}
	return nil, nil
}

func Get() (IPMI, error) {
	i := IPMI{}

	// ipmitool
	if err := parse.Exists("ipmitool"); err == nil {
		i.IpmitoolInstalled = true

		o, err := parse.Exec("ipmitool", []string{"-V"})
		if err != nil {
			return IPMI{}, err
		}
		arr := strings.Split(o, " ")
		i.IpmitoolVersion = arr[2]
	} else {
		i.IpmitoolInstalled = false
		return i, nil
	}

	// Fans
	m, err := parse.ExecRegexpMap("ipmitool", []string{"sdr", "type", "fan"}, "\\|.*\\|", "\\|\\sok")
	if err != nil {
		return IPMI{}, err
	}

	i.Fans = Fans{}
	for k, v := range m {
		matched, _ := regexp.MatchString("Fan[0-9].*", k)
		if !matched {
			continue
		}

		arr := strings.Split(v, " ")
		rpm, err := strconv.Atoi(arr[0])
		if err != nil {
			return IPMI{}, fmt.Errorf("failed parsing field: %s error: %s", k, err.Error())
		}

		i.Fans = append(i.Fans, Fan{
			Name:     k,
			SpeedRPM: rpm,
		})
	}

	// Voltage
	m, err = parse.ExecRegexpMap("ipmitool", []string{"sdr", "type", "voltage"}, "\\|.*\\|", "\\|\\sok")
	if err != nil {
		return IPMI{}, err
	}

	i.Voltage1, err = strToIntPtr(m, "Voltage 1")
	if err != nil {
		return IPMI{}, err
	}

	i.Voltage2, err = strToIntPtr(m, "Voltage 2")
	if err != nil {
		return IPMI{}, err
	}

	// Current
	m, err = parse.ExecRegexpMap("ipmitool", []string{"sdr", "type", "voltage"}, "\\|.*\\|", "\\|\\sok")
	if err != nil {
		return IPMI{}, err
	}

	i.Current1, err = strToFloat32Ptr(m, "Current 1")
	if err != nil {
		return IPMI{}, err
	}

	i.Current2, err = strToFloat32Ptr(m, "Current 2")
	if err != nil {
		return IPMI{}, err
	}

	i.PowerCons, err = strToIntPtr(m, "Pwr Consumption")
	if err != nil {
		return IPMI{}, err
	}

	// Temperature
	m, err = parse.ExecRegexpMap("ipmitool", []string{"sdr", "type", "temperature"}, "\\|.*\\|", "\\|\\sok")
	if err != nil {
		return IPMI{}, err
	}

	i.InletTempDegrC, err = strToIntPtr(m, "Inlet Temp")
	if err != nil {
		return IPMI{}, err
	}

	i.ExhaustTempDegrC, err = strToIntPtr(m, "Exhaust Temp")
	if err != nil {
		return IPMI{}, err
	}

	return i, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
