package hwinfo

import (
	"github.com/mickep76/hwinfo/cpu"
	"github.com/mickep76/hwinfo/mem"
	"github.com/mickep76/hwinfo/os"
)

// Info structure for information a system.
type Info struct {
	CPU *cpu.Info `json:"cpu"`
	Mem *mem.Info `json:"mem"`
	OS  *os.Info  `json:"os"`
}

// GetInfo return information about a system.
func GetInfo() (Info, error) {
	h := Info{}

	c, err := cpu.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.CPU = &c

	m, err := mem.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.Mem = &m

	o, err := os.GetInfo()
	if err != nil {
		return Info{}, err
	}
	h.OS = &o

	return h, nil
}
