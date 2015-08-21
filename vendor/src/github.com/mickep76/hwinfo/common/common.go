package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// LoadFileFields returns fields from file.
func LoadFileFields(fn string, del string, fields []string) (map[string]string, error) {
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return map[string]string{}, fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return map[string]string{}, err
	}

	r, err := parseFields(string(o), del, fields)
	if err != nil {
		return map[string]string{}, err
	}

	return r, nil
}

// ExecCmdFields returns fields from command output.
func ExecCmdFields(cmd string, args []string, del string, fields []string) (map[string]string, error) {
	o, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return map[string]string{}, err
	}

	r, err := parseFields(string(o), del, fields)
	if err != nil {
		return map[string]string{}, err
	}

	return r, nil
}

func parseFields(o string, del string, fields []string) (map[string]string, error) {
	r := make(map[string]string)

	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.Split(line, del)
		if len(vals) < 1 {
			continue
		}

		for _, f := range fields {
			if strings.HasPrefix(line, f) {
				r[f] = strings.Trim(strings.Join(vals[1:], " "), " \t")
			}
		}
	}

	return r, nil
}
