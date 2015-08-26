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
	Slot      string `json:"slot"`
	ClassID   string `json:"class_id"`
	Class     string `json:"class"`
	VendorID  string `json:"vendor_id"`
	DeviceID  string `json:"device_id"`
	Vendor    string `json:"vendor"`
	Device    string `json:"device"`
	SVendorID string `json:"svendor_id"`
	SDeviceID string `json:"sdevice_id"`
	SName     string `json:"sname,omiempty"`
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
func getPCIVendor(vendorID string, deviceID string, sVendorID string, sDeviceID string) (string, string, string, error) {
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
	sName := ""
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

		if vendor != "" && device != "" && strings.LastIndex(line, "\t") == 1 && vals[0] == sVendorID && vals[1] == sDeviceID {
			sName = vals[2]
			break
		}

		if vendor != "" && strings.LastIndex(line, "\t") == -1 {
			break
		}
	}

	return vendor, device, sName, nil
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
	if len(classID) < 6 {
		return "", "", "", fmt.Errorf("Class string is to short: %s", classID)
	}

	subClassID := classID[2:4]
	progIntfID := classID[4:6]
	classID = classID[0:2]

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
		slot := strings.SplitN(path, ":", 2)

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

		classID := o["class"][2:]
		class, subClass, _, err := getPCIClass(classID)
		if err != nil {
			return Info{}, err
		}
		if subClass != "" {
			class = subClass
		}

		vendorID := o["vendor"][2:]
		deviceID := o["device"][2:]
		sVendorID := o["subsystem_vendor"][2:]
		sDeviceID := o["subsystem_device"][2:]
		vendor, device, sName, err := getPCIVendor(vendorID, deviceID, sVendorID, sDeviceID)
		if err != nil {
			return Info{}, err
		}

		i.PCI = append(i.PCI, PCI{
			Slot:      slot[1],
			ClassID:   classID,
			Class:     class,
			VendorID:  vendorID,
			DeviceID:  deviceID,
			SDeviceID: sDeviceID,
			SVendorID: sVendorID,
			Vendor:    vendor,
			Device:    device,
			SName:     sName,
		})
	}

	return i, nil
}
