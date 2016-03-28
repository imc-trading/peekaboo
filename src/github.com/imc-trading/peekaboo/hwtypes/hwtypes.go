package hwtypes

import (
	"fmt"
	"strings"
)

func List() {
	fmt.Println(strings.Join(hwTypes, "\n"))
}
