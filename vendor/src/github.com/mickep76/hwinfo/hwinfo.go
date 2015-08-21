package hwinfo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func cpuInfo() (map[string]string, error) {
	d := make(map[string]string)
	logical := int64(0)

	b, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return map[string]string{}, fmt.Errorf("can't read file: %s", err)
	}

	cpuID := int64(-1)
	cpuIDs := make(map[int64]int64)
	cores := int64(0)
	for _, line := range strings.Split(string(b), "\n") {
		values := strings.Split(line, ":")
		if len(values) < 1 {
			continue
		} else if _, ok := d["cpu_model"]; !ok && strings.HasPrefix(line, "model name") {
			d["cpu_model"] = strings.Trim(strings.Join(values[1:], " "), " ")
		} else if _, ok := d["cpu_flags"]; !ok && strings.HasPrefix(line, "flags") {
			d["cpu_flags"] = strings.Trim(strings.Join(values[1:], " "), " ")
		} else if _, ok := d["cpu_cores"]; !ok && strings.HasPrefix(line, "cpu cores") {
			cores, _ = strconv.ParseInt(strings.Trim(strings.Join(values[1:], " "), " "), 10, 0)
		} else if strings.HasPrefix(line, "processor") {
			logical++
		} else if strings.HasPrefix(line, "physical id") {
			cpuID, _ = strconv.ParseInt(strings.Trim(strings.Join(values[1:], " "), " "), 10, 0)
			cpuIDs[cpuID] = cpuIDs[cpuID] + 1
		}
	}
	d["cpu_logical"] = strconv.FormatInt(logical, 10)
	sockets := int64(len(cpuIDs))
	d["cpu_sockets"] = strconv.FormatInt(sockets, 10)
	d["cpu_cores_per_socket"] = strconv.FormatInt(cores, 10)
	physical := int64(len(cpuIDs)) * cores
	d["cpu_physical"] = strconv.FormatInt(physical, 10)
	t := logical / sockets / cores
	d["cpu_threads_per_core"] = strconv.FormatInt(t, 10)

	return d, nil
}

func loadFiles(files map[string]string) (map[string]string, error) {
	d := make(map[string]string)

	for k, v := range files {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			return map[string]string{}, errors.New("file doesn't exist")
		}

		b, err := ioutil.ReadFile(v)
		if err != nil {
			return map[string]string{}, fmt.Errorf("can't read file: %s", err)
		}

		d[k] = strings.Trim(string(b), "\n")
	}

	return d, nil
}

func loadFile(file string, del string, fields map[string]string) (map[string]string, error) {
	d := make(map[string]string)

	out, err := ioutil.ReadFile(file)
	if err != nil {
		return map[string]string{}, fmt.Errorf("can't read file: %s", err)
	}

	for _, line := range strings.Split(string(out), "\n") {
		values := strings.Split(line, del)
		if len(values) < 1 {
			continue
		}

		for k, v := range fields {
			if strings.HasPrefix(line, v) {
				d[k] = strings.Trim(strings.Join(values[1:], " "), " \t")
			}
		}
	}

	return d, nil
}

func execCmd(cmd string, args []string, del string, fields map[string]string) (map[string]string, error) {
	d := make(map[string]string)

	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return map[string]string{}, err
	}

	for _, line := range strings.Split(string(out), "\n") {
		values := strings.Split(line, del)
		if len(values) < 1 {
			continue
		}

		for k, v := range fields {
			if strings.HasPrefix(strings.Trim(line, " "), v) {
				d[k] = strings.Trim(strings.Join(values[1:], " "), " \t")
			}
		}
	}

	return d, nil
}

func merge(a map[string]string, b map[string]string) {
	for k, v := range b {
		a[k] = v
	}
}

// HWInfo returns a map[string]string with hardware info about the current system.
func HWInfo() (map[string]string, error) {
	sysFiles := map[string]string{
		// check for perm. to read it
		//		"serial_number":   "/sys/devices/virtual/dmi/id/product_serial",
		"manufacturer":    "/sys/devices/virtual/dmi/id/chassis_vendor",
		"product_version": "/sys/devices/virtual/dmi/id/product_version",
		"product":         "/sys/devices/virtual/dmi/id/product_name",
		"bios_date":       "/sys/devices/virtual/dmi/id/bios_date",
		"bios_vendor":     "/sys/devices/virtual/dmi/id/bios_vendor",
		"bios_version":    "/sys/devices/virtual/dmi/id/bios_version",
	}

	sysctlFields := map[string]string{
		"mem_total_b":          "hw.memsize",
		"cpu_cores_per_socket": "machdep.cpu.core_count",
		"cpu_physical":         "hw.physicalcpu_max",
		"cpu_logical":          "hw.logicalcpu_max",
		"cpu_model":            "machdep.cpu.brand_string",
		"cpu_flags":            "machdep.cpu.features",
	}

	swVersFields := map[string]string{
		"os_name":    "ProductName",
		"os_version": "ProductVersion",
	}

	lsbReleaseFields := map[string]string{
		"os_name":    "Distributor ID",
		"os_version": "Release",
	}

	meminfoFields := map[string]string{
		"mem_total_kb": "MemTotal",
	}

	systemProfilerFields := map[string]string{
		"model_name":        "Model Name",
		"model_id":          "Model Identifier",
		"boot_room_version": "Boot ROM Version",
		"smc_version":       "SMC Version",
		"serial_number":     "Serial Number",
	}

	sys := make(map[string]string)

	sys["os_kernel"] = runtime.GOOS

	h, err := os.Hostname()
	if err != nil {
		return map[string]string{}, err
	}
	sys["fqdn"] = h

	addrs, _ := net.LookupIP(h)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			sys["fqdn_ip"] = ipv4.String()
		}
	}

	switch runtime.GOOS {
	case "darwin":
		o, err := execCmd("/usr/sbin/sysctl", []string{"-a"}, ":", sysctlFields)
		if err != nil {
			return map[string]string{}, err
		}

		merge(sys, o)

		b, err := strconv.ParseUint(sys["mem_total_b"], 10, 64)
		if err != nil {
			return map[string]string{}, err
		}

		kb := b / 1024
		mb := kb / 1024
		gb := mb / 1024
		sys["mem_total_kb"] = strconv.FormatUint(kb, 10)
		sys["mem_total_mb"] = strconv.FormatUint(mb, 10)
		sys["mem_total_gb"] = strconv.FormatUint(gb, 10)

		c, _ := strconv.ParseUint(sys["cpu_cores_per_socket"], 10, 64)
		p, _ := strconv.ParseUint(sys["cpu_physical"], 10, 64)
		l, _ := strconv.ParseUint(sys["cpu_logical"], 10, 64)
		s := p / c
		sys["cpu_sockets"] = strconv.FormatUint(s, 10)
		t := l / s / c
		sys["cpu_threads_per_core"] = strconv.FormatUint(t, 10)

		sys["cpu_flags"] = strings.ToLower(sys["cpu_flags"])

		o2, err2 := execCmd("/usr/bin/sw_vers", []string{}, ":", swVersFields)
		if err2 != nil {
			return map[string]string{}, err2
		}
		merge(sys, o2)

		o3, err3 := execCmd("/usr/sbin/system_profiler", []string{"SPHardwareDataType"}, ":", systemProfilerFields)
		if err3 != nil {
			return map[string]string{}, err3
		}
		merge(sys, o3)

	case "linux":
		o, err := loadFiles(sysFiles)
		if err != nil {
			return map[string]string{}, err
		}
		merge(sys, o)

		if strings.Contains(sys["product_version"], "amazon") {
			sys["virtual"] = "Amazon EC2"
		}

		o2, err2 := execCmd("/usr/bin/lsb_release", []string{"-a"}, ":", lsbReleaseFields)
		if err2 != nil {
			return map[string]string{}, err2
		}
		merge(sys, o2)

		o3, err3 := cpuInfo()
		if err3 != nil {
			return map[string]string{}, err3
		}
		merge(sys, o3)

		o4, err4 := loadFile("/proc/meminfo", ":", meminfoFields)
		if err4 != nil {
			return map[string]string{}, err4
		}
		merge(sys, o4)

		sys["mem_total_kb"] = strings.Trim(sys["mem_total_kb"], " kB")

		kb, err := strconv.ParseUint(sys["mem_total_kb"], 10, 64)
		if err != nil {
			return map[string]string{}, err
		}

		b := kb * 1024
		mb := kb / 1024
		gb := mb / 1024
		sys["mem_total_b"] = strconv.FormatUint(b, 10)
		sys["mem_total_mb"] = strconv.FormatUint(mb, 10)
		sys["mem_total_gb"] = strconv.FormatUint(gb, 10)

	default:
		return map[string]string{}, fmt.Errorf("unsupported plattform (%s), needs to be either linux or darwin", runtime.GOOS)
	}

	return sys, nil
}
