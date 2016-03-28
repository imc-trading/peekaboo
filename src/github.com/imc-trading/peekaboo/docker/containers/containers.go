package containers

import (
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type Containers []Container

type Container struct {
	ID         string `json:"id"`
	Image      string `json:"image"`
	Command    string `json:"command"`
	CreatedAt  string `json:"created_at"`
	RunningFor string `json:"running_for"`
	Ports      string `json:"ports"`
	Status     string `json:"status"`
	Size       string `json:"size"`
	Names      string `json:"names"`
	Labels     string `json:"labels"`
}

func Get() (Containers, error) {
	cnts := Containers{}

	o, err := parse.Exec("docker", []string{"ps", "-a", "--no-trunc=true",
		"--format={{.ID}}!{{.Image}}!{{.Command}}!{{.CreatedAt}}!{{.RunningFor}}!{{.Ports}}!{{.Status}}!{{.Size}}!{{.Names}}!{{.Labels}}"})
	if err != nil {
		return Containers{}, err
	}

	for _, line := range strings.Split(o, "\n") {
		v := strings.Split(line, "!")
		if len(v) < 10 {
			continue
		}

		c := Container{}

		c.ID = v[0]
		c.Image = v[1]
		c.Command = v[2]
		c.CreatedAt = v[3]
		c.RunningFor = v[4]
		c.Ports = v[5]
		c.Status = v[6]
		c.Size = v[7]
		c.Names = v[8]
		c.Labels = v[9]

		cnts = append(cnts, c)
	}

	return cnts, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
