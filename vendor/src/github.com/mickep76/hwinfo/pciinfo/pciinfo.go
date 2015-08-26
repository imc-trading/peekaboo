package pciinfo

import (
	"fmt"
	"github.com/mickep76/hwinfo/common"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type PCI struct {
	BusNum      string `json:"bus_number"`
	DeviceNum   string `json:"device_number"`
	DeviceFunc  string `json:"device_function"`
	ClassID     string `json:"class_id"`
	ClassName   string `json:"class_name"`
	VendorID    string `json:"vendor_id"`
	DeviceID    string `json:"device_id"`
	SubVendorID string `json:"subsys_vendor_id"`
	SubDeviceID string `json:"subsys_device_id"`
	VendorName  string `json:"vendor_name"`
	DeviceName  string `json:"device_name"`
	SubsysName  string `json:"subsys_name,omitempty"`
}

// Info structure for information about a systems memory.
type Info struct {
	PCI []PCI `json:"pci"`
}

/*
# Syntax:
# vendor  vendor_name
#       device  device_name                             <-- single tab
#               subvendor subdevice  subsystem_name     <-- two tabs

0010  Allied Telesis, Inc (Wrong ID)
# This is a relabelled RTL-8139
        8139  AT-2500TX V3 Ethernet
001c  PEAK-System Technik GmbH
        0001  PCAN-PCI CAN-Bus controller
                001c 0004  2 Channel CAN Bus SJC1000
                001c 0005  2 Channel CAN Bus SJC1000 (Optically Isolated)
...
*/

// TODO: Cache PCI database as a map[string]string
func getPCIVendor(vendorID string, deviceID string, subsysVendorID string, subsysDeviceID string) (string, string, string, error) {

	vendorID = strings.Replace(vendorID, "0x", "", 1)
	deviceID = strings.Replace(deviceID, "0x", "", 1)
	subsysVendorID = strings.Replace(subsysVendorID, "0x", "", 1)
	subsysDeviceID = strings.Replace(subsysDeviceID, "0x", "", 1)

	fn := "/usr/share/hwdata/pci.ids"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return "", "", "", fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", "", "", err
	}

	vendor := ""
	device := ""
	subsysName := ""
	cols := 2
	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.SplitN(line, " ", cols)
		if len(vals) < 2 || strings.HasPrefix(line, "#") {
			continue
		}

		for i := 0; i < cols; i++ {
			vals[i] = strings.Trim(vals[i], " \t")
		}

		if strings.LastIndex(line, "\t") == -1 && vals[0] == vendorID {
			vendor = vals[1]
			continue
		}

		if vendor != "" && strings.LastIndex(line, "\t") == 0 && vals[0] == deviceID {
			device = vals[1]
			cols = 3
			continue
		}

		if vendor != "" && device != "" && strings.LastIndex(line, "\t") == 1 && vals[0] == subsysVendorID && vals[1] == subsysDeviceID {
			subsysName = vals[2]
			break
		}

		if vendor != "" && strings.LastIndex(line, "\t") == -1 {
			break
		}
	}

	return vendor, device, subsysName, nil
}

/*
# List of known device classes, subclasses and programming interfaces

# Syntax:
# C class       class_name
#       subclass        subclass_name           <-- single tab
#               prog-if  prog-if_name   <-- two tabs

C 01  Mass storage controller
        00  SCSI storage controller
        01  IDE interface
        02  Floppy disk controller
        03  IPI bus controller
        04  RAID bus controller
        05  ATA controller
                20  ADMA single stepping
                30  ADMA continuous operation
...
*/

// TODO: Cache PCI database as a map[string]string
func getPCIClass(classID string) (string, string, string, error) {
	if len(classID) < 8 {
		return "", "", "", fmt.Errorf("Class string is to short: %s", classID)
	}

	subClassID := classID[4:6]
	progIntfID := classID[6:8]
	classID = classID[2:4]

	fn := "/usr/share/hwdata/pci.ids"
	if _, err := os.Stat(fn); os.IsNotExist(err) {
		return "", "", "", fmt.Errorf("file doesn't exist: %s", fn)
	}

	o, err := ioutil.ReadFile(fn)
	if err != nil {
		return "", "", "", err
	}

	class := ""
	subClass := ""
	progIntf := ""
	begins := false
	for _, line := range strings.Split(string(o), "\n") {
		if !begins && strings.HasPrefix(line, "C 01  Mass storage controller") {
			begins = true
		} else if !begins || strings.HasPrefix(line, "#") {
			continue
		}

		vals := strings.SplitN(strings.Replace(line, "C ", "", 1), " ", 2)
		if len(vals) < 2 {
			continue
		}

		for i := 0; i < 2; i++ {
			vals[i] = strings.Trim(vals[i], " \t")
		}

		if strings.LastIndex(line, "\t") == -1 && vals[0] == classID {
			class = vals[1]
			continue
		}

		if class != "" && strings.LastIndex(line, "\t") == 0 && vals[0] == subClassID {
			subClass = vals[1]
			continue
		}

		if class != "" && subClass != "" && strings.LastIndex(line, "\t") == 1 && vals[0] == progIntfID {
			progIntf = vals[1]
			break
		}

		if class != "" && strings.LastIndex(line, "\t") == -1 {
			break
		}
	}

	return class, subClass, progIntf, nil
}

// GetInfo return information about PCI devices.
func GetInfo() (Info, error) {
	i := Info{}

	files, err := filepath.Glob("/sys/bus/pci/devices/*")
	if err != nil {
		return Info{}, err
	}

	for _, path := range files {
		pci := strings.Split(path, ":")
		dev := strings.Split(pci[2], ".")

		o, err := common.LoadFiles([]string{
			filepath.Join(path, "class"),
			filepath.Join(path, "vendor"),
			filepath.Join(path, "device"),
			filepath.Join(path, "subsystem_vendor"),
			filepath.Join(path, "subsystem_device"),
		})
		if err != nil {
			return Info{}, err
		}

		classID := o["class"]
		class, subClass, _, err := getPCIClass(classID)
		if err != nil {
			return Info{}, err
		}
		if subClass != "" {
			class = subClass
		}

		vendorID := o["vendor"]
		deviceID := o["device"]
		subVendorID := o["subsystem_vendor"]
		subDeviceID := o["subsystem_device"]
		vendor, device, subsysName, err := getPCIVendor(vendorID, deviceID, subVendorID, subDeviceID)
		if err != nil {
			return Info{}, err
		}

		i.PCI = append(i.PCI, PCI{
			BusNum:      pci[1],
			DeviceNum:   dev[0],
			DeviceFunc:  dev[1],
			ClassID:     classID,
			ClassName:   class,
			VendorID:    vendorID,
			DeviceID:    deviceID,
			SubDeviceID: subDeviceID,
			SubVendorID: subVendorID,
			VendorName:  vendor,
			DeviceName:  device,
			SubsysName:  subsysName,
		})
	}

	return i, nil
}
