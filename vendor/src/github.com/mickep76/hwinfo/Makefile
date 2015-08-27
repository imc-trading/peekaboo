all: readme

readme:
	godoc2md github.com/mickep76/hwinfo | grep -v Generated >README.md
	godoc2md github.com/mickep76/hwinfo/cpuinfo | grep -v Generated >>README.md
	godoc2md github.com/mickep76/hwinfo/meminfo | grep -v Generated >>README.md
	godoc2md github.com/mickep76/hwinfo/osinfo | grep -v Generated >>README.md
	godoc2md github.com/mickep76/hwinfo/sysinfo | grep -v Generated >>README.md
	godoc2md github.com/mickep76/hwinfo/pciinfo | grep -v Generated >>README.md
	godoc2md github.com/mickep76/hwinfo/diskinfo | grep -v Generated >>README.md
