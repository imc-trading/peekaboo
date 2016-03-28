package images

import (
	"regexp"
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type Images []Image

type Image struct {
	ID       string `json:"id"`
	Repo     string `json:"repo"`
	Tag      string `json:"tag"`
	Created  string `json:"created"`
	VirtSize string `json:"virtSize"`
}

func Get() (Images, error) {
	images := Images{}

	o, err := parse.Exec("docker", []string{"images", "--no-trunc=true"})
	if err != nil {
		return nil, err
	}

	for c, line := range strings.Split(string(o), "\n") {
		re := regexp.MustCompile(`\s{2,}`)
		v := re.Split(line, -1)

		if c < 1 || len(v) < 5 {
			continue
		}

		i := Image{}

		if v[0] == "<none>" {
			i.Repo = ""
		} else {
			i.Repo = v[0]
		}

		if v[1] == "<none>" {
			i.Tag = ""
		} else {
			i.Tag = v[1]
		}

		i.ID = strings.TrimLeft(v[2], "sha256:")
		i.Created = v[3]
		i.VirtSize = v[4]

		images = append(images, i)
	}

	return images, nil
}

func GetInterface() (interface{}, error) {
	return Get()
}
