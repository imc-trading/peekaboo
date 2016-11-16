// +build linux

package load

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Get() (Load, error) {
	l := Load{}
	loadFileName := "/proc/loadavg"

	if _, err := os.Stat(loadFileName); os.IsNotExist(err) {
		return Load{}, errors.New("file doesn't exist: " + loadFileName)
	}

	loadString, err := ioutil.ReadFile(loadFileName)
	if err != nil {
		return Load{}, err
	}

	tokens := strings.Split(strings.TrimRight(string(loadString), "\n"), " ")

	avg1, _ := strconv.ParseFloat(tokens[0], 32)
	avg5, _ := strconv.ParseFloat(tokens[1], 32)
	avg15, _ := strconv.ParseFloat(tokens[2], 32)
	l.Avg1 = float32(avg1)
	l.Avg5 = float32(avg5)
	l.Avg15 = float32(avg15)

	procTokens := strings.Split(tokens[3], "/")
	l.Running, _ = strconv.Atoi(procTokens[0])
	l.Total, _ = strconv.Atoi(procTokens[1])

	l.LastPid, _ = strconv.Atoi(tokens[4])

	return l, nil
}
