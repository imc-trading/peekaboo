package parse

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func Exists(cmd string) error {
	// Check command exist's.
	_, err := exec.LookPath(cmd)
	if err != nil {
		return fmt.Errorf("command doesn't exist: %s", cmd)
	}
	return nil
}

// ExecRegexpMap execute command and split on regexp returning key/values as a map.
func ExecRegexpMap(cmd string, args []string, delim string, match string) (map[string]string, error) {
	// Check command exist's.
	_, err := exec.LookPath(cmd)
	if err != nil {
		return nil, fmt.Errorf("command doesn't exist: %s", cmd)
	}

	// Execute command.
	c := exec.Command(cmd, args...)
	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	if err := c.Run(); err != nil {
		return nil, fmt.Errorf("failed to exec: \"%s %s\" error: %s stderr: %s", cmd, strings.Join(args, " "), err.Error(), stderr.String())
	}

	// Compile regexp.
	reDelim := regexp.MustCompile(delim)
	reMatch := regexp.MustCompile(match)

	// Parse output and create key/value map.
	m := make(map[string]string)
	for _, l := range strings.Split(stdout.String(), "\n") {
		if !reMatch.MatchString(l) {
			continue
		}

		v := reDelim.Split(l, -1)
		if len(v) < 1 {
			continue
		} else if len (v) > 1 {
			v[1] = strings.TrimSpace(strings.Join(v[1:], delim))
		}

		m[strings.TrimSpace(v[0])] = strings.TrimSpace(v[1])
	}

	return m, nil
}

// FileRegexpMap execute command and split on regexp returning key/values as a map.
func FileRegexpMap(fn string, delim string, match string) (map[string]string, error) {
	// Check file exist's.
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return nil, fmt.Errorf("file doesn't exist: %s", fn)
	}

	r, err := ioutil.ReadFile(fn)
	if err != nil {
		return map[string]string{}, err
	}

	// Compile regexp.
	reDelim := regexp.MustCompile(delim)
	reMatch := regexp.MustCompile(match)

	// Parse output and create key/value map.
	m := make(map[string]string)
	for _, l := range strings.Split(string(r), "\n") {
		if !reMatch.MatchString(l) {
			continue
		}

		v := reDelim.Split(l, -1)
		if len(v) < 1 {
			continue
		}

		m[strings.TrimSpace(v[0])] = strings.TrimSpace(v[1])
	}

	return m, nil
}

// Exec returns output.
func Exec(cmd string, args []string) (string, error) {
	// Check command exist's.
	_, err := exec.LookPath(cmd)
	if err != nil {
		return "", fmt.Errorf("command doesn't exist: %s", cmd)
	}

	// Execute command.
	c := exec.Command(cmd, args...)
	var stdout, stderr bytes.Buffer
	c.Stdout = &stdout
	c.Stderr = &stderr
	if err := c.Run(); err != nil {
		return "", fmt.Errorf("failed to exec: \"%s %s\" error: %s stderr: %s", cmd, strings.Join(args, " "), err.Error(), stderr.String())
	}

	return strings.TrimRight(stdout.String(), "\n"), err
}

// LoadFiles returns info from multiple files.
func LoadFiles(files []string) (map[string]string, error) {
	r := make(map[string]string)

	for _, fn := range files {
		if _, err := os.Stat(fn); os.IsNotExist(err) {
			return map[string]string{}, fmt.Errorf("file doesn't exist: %s", fn)
		}

		o, err := ioutil.ReadFile(fn)
		if err != nil {
			return map[string]string{}, err
		}

		r[path.Base(fn)] = strings.TrimRight(string(o), "\n")
	}

	return r, nil
}

// StrToInt convert string to integer.
func StrToInt(m map[string]string, k string) (int, error) {
	v, err := strconv.Atoi(m[k])
	if err != nil {
		return 0, fmt.Errorf("failed to convert key: \"%s\" with value: %s to an integer, error: %s", k, m[k], err.Error())
	}
	return v, nil
}
