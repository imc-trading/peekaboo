// +build linux

package memory

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/imc-trading/peekaboo/parse"
)

type Memory struct {
	TotalKB             *int     `json:"totalKB,omitempty"`
	TotalGB             *float32 `json:"totalGB,omitempty"`
	FreeKB              *int     `json:"freeKB,omitempty"`
	FreeGB              *float32 `json:"freeGB,omitempty"`
	AvailableKB         *int     `json:"availableKB,omitempty"`
	AvailableGB         *float32 `json:"availableGB,omitempty"`
	UsedKB              *int     `json:"usedKB,omitempty"`
	UsedGB              *float32 `json:"usedGB,omitempty"`
	BuffersKB           *int     `json:"buffersKB,omitempty"`
	BuffersGB           *float32 `json:"buffersGB,omitempty"`
	CachedKB            *int     `json:"cachedKB,omitempty"`
	CachedGB            *float32 `json:"cachedGB,omitempty"`
	SwapCachedKB        *int     `json:"swapCachedKB,omitempty"`
	SwapCachedGB        *float32 `json:"swapCachedGB,omitempty"`
	ActiveKB            *int     `json:"activeKB,omitempty"`
	ActiveGB            *float32 `json:"activeGB,omitempty"`
	InactiveKB          *int     `json:"inactiveKB,omitempty"`
	InactiveGB          *float32 `json:"inactiveGB,omitempty"`
	ActiveAnonKB        *int     `json:"activeAnonKB,omitempty"`
	ActiveAnonGB        *float32 `json:"activeAnonKB,omitempty"`
	InactiveAnonKB      *int     `json:"inactiveAnonBuffersKB,omitempty"`
	InactiveAnonGB      *float32 `json:"inactiveAnonBuffersKB,omitempty"`
	ActiveFileKB        *int     `json:"activeFileKB,omitempty"`
	ActiveFileGB        *float32 `json:"activeFileGB,omitempty"`
	InactiveFileKB      *int     `json:"inactiveFileKB,omitempty"`
	InactiveFileGB      *float32 `json:"inactiveFileGB,omitempty"`
	UnevictableKB       *int     `json:"unevictableKB,omitempty"`
	UnevictableGB       *float32 `json:"unevictableGB,omitempty"`
	MLockedKB           *int     `json:"mLockedKB,omitempty"`
	MLockedGB           *float32 `json:"mLockedGB,omitempty"`
	SwapTotalKB         *int     `json:"swapTotalKB,omitempty"`
	SwapTotalGB         *float32 `json:"swapTotalGB,omitempty"`
	SwapFreeKB          *int     `json:"swapFreeKB,omitempty"`
	SwapFreeGB          *float32 `json:"swapFreeGB,omitempty"`
	DirtyKB             *int     `json:"dirtyKB,omitempty"`
	DirtyGB             *float32 `json:"dirtyGB,omitempty"`
	WritebackKB         *int     `json:"writebackKB,omitempty"`
	WritebackGB         *float32 `json:"writebackGB,omitempty"`
	AnonPagesKB         *int     `json:"anonPagesKB,omitempty"`
	AnonPagesGB         *float32 `json:"anonPagesGB,omitempty"`
	MappedKB            *int     `json:"mappedKB,omitempty"`
	MappedGB            *float32 `json:"mappedGB,omitempty"`
	ShmemKB             *int     `json:"shmemKB,omitempty"`
	ShmemGB             *float32 `json:"shmemGB,omitempty"`
	SlabKB              *int     `json:"slabKB,omitempty"`
	SlabGB              *float32 `json:"slabGB,omitempty"`
	SReclaimableKB      *int     `json:"sReclaimableKB,omitempty"`
	SReclaimableGB      *float32 `json:"sReclaimableGB,omitempty"`
	SUnreclaimKB        *int     `json:"sUnreclaimKB,omitempty"`
	SUnreclaimGB        *float32 `json:"sUnreclaimGB,omitempty"`
	KernelStackKB       *int     `json:"kernelStackKB,omitempty"`
	KernelStackGB       *float32 `json:"kernelStackGB,omitempty"`
	PageTablesKB        *int     `json:"pageTablesKB,omitempty"`
	PageTablesGB        *float32 `json:"pageTablesGB,omitempty"`
	NFSUnstableKB       *int     `json:"nfsUnstableKB,omitempty"`
	NFSUnstableGB       *float32 `json:"nfsUnstableGB,omitempty"`
	BounceKB            *int     `json:"bounceKB,omitempty"`
	BounceGB            *float32 `json:"bounceGB,omitempty"`
	WritebackTmpKB      *int     `json:"writebackTmpKB,omitempty"`
	WritebackTmpGB      *float32 `json:"writebackTmpGB,omitempty"`
	CommitLimitKB       *int     `json:"commitLimitKB,omitempty"`
	CommitLimitGB       *float32 `json:"commitLimitGB,omitempty"`
	CommittedASKB       *int     `json:"committedASKB,omitempty"`
	CommittedASGB       *float32 `json:"committedASGB,omitempty"`
	VmallocTotalKB      *int     `json:"vmallocTotalKB,omitempty"`
	VmallocTotalGB      *float32 `json:"vmallocTotalGB,omitempty"`
	VmallocUsedKB       *int     `json:"vmallocUsedKB,omitempty"`
	VmallocUsedGB       *float32 `json:"vmallocUsedGB,omitempty"`
	VmallocChunkKB      *int     `json:"vmallocChunkKB,omitempty"`
	VmallocChunkGB      *float32 `json:"vmallocChunkGB,omitempty"`
	HardwareCorruptedKB *int     `json:"hardwareCorruptedKB,omitempty"`
	HardwareCorruptedGB *float32 `json:"hardwareCorruptedGB,omitempty"`
	AnonHugePagesKB     *int     `json:"anonHugePagesKB,omitempty"`
	AnonHugePagesGB     *float32 `json:"anonHugePagesGB,omitempty"`
	HugePagesTotal      *int     `json:"hugePagesTotal,omitempty"`
	HugePagesFree       *int     `json:"hugePagesFree,omitempty"`
	HugePagesRsvd       *int     `json:"hugePagesRsvd,omitempty"`
	HugePagesSurp       *int     `json:"hugePagesSurp,omitempty"`
	HugePageSizeKB      *int     `json:"hugePagesSizeKB,omitempty"`
	DirectMap4kKB       *int     `json:"directMap4kKB,omitempty"`
	DirectMap2MKB       *int     `json:"directMap2mKB,omitempty"`
	DirectMap1GKB       *int     `json:"directMap1gKB,omitempty"`
	CapacityFreeKB      *int     `json:"capacityFreeKB,omitempty"`
	CapacityUsedKB      *int     `json:"capacityUsedKB,omitempty"`
	CapacityFreeGB      *float32 `json:"capacityFreeGB,omitempty"`
	CapacityUsedGB      *float32 `json:"capacityUsedGB,omitempty"`
}

func strToIntPtr(m map[string]string, f string) (*int, *float32, error) {
	if v, ok := m[f]; ok {
		kb, err := strconv.Atoi(strings.TrimRight(v, " kB"))
		if err != nil {
			return nil, nil, fmt.Errorf("failed parsing field: %s error: %s", f, err.Error())
		}

		gb := float32(kb) / 1024 / 1024
		return &kb, &gb, nil
	}
	return nil, nil, nil
}

func Get() (Memory, error) {
	m := Memory{}

	o, err2 := parse.FileRegexpMap("/proc/meminfo", ":", "\\S+:\\s+\\S+")
	if err2 != nil {
		return Memory{}, err2
	}

	var err error

	// MemTotal
	m.TotalKB, m.TotalGB, err = strToIntPtr(o, "MemTotal")
	if err != nil {
		return Memory{}, err
	}

	// MemFree
	m.FreeKB, m.FreeGB, err = strToIntPtr(o, "MemFree")
	if err != nil {
		return Memory{}, err
	}

	// Buffers
	m.BuffersKB, m.BuffersGB, err = strToIntPtr(o, "Buffers")
	if err != nil {
		return Memory{}, err
	}

	// Cached
	m.CachedKB, m.CachedGB, err = strToIntPtr(o, "Cached")
	if err != nil {
		return Memory{}, err
	}

	// SwapCached
	m.SwapCachedKB, m.SwapCachedGB, err = strToIntPtr(o, "SwapCached")
	if err != nil {
		return Memory{}, err
	}

	// Active
	m.ActiveKB, m.ActiveGB, err = strToIntPtr(o, "Active")
	if err != nil {
		return Memory{}, err
	}

	// Inactive
	m.InactiveKB, m.InactiveGB, err = strToIntPtr(o, "Inactive")
	if err != nil {
		return Memory{}, err
	}

	// Active(anon)
	m.ActiveAnonKB, m.ActiveAnonGB, err = strToIntPtr(o, "Active(anon)")
	if err != nil {
		return Memory{}, err
	}

	// Inactive(anon)
	m.InactiveAnonKB, m.InactiveAnonGB, err = strToIntPtr(o, "Inactive(anon)")
	if err != nil {
		return Memory{}, err
	}

	// Active(file)
	m.ActiveFileKB, m.ActiveFileGB, err = strToIntPtr(o, "Active(file)")
	if err != nil {
		return Memory{}, err
	}

	// Inactive(file)
	m.InactiveFileKB, m.InactiveFileGB, err = strToIntPtr(o, "Inactive(file)")
	if err != nil {
		return Memory{}, err
	}

	// Unevictable
	m.UnevictableKB, m.UnevictableGB, err = strToIntPtr(o, "Unevictable")
	if err != nil {
		return Memory{}, err
	}

	// Mlocked
	m.MLockedKB, m.MLockedGB, err = strToIntPtr(o, "Mlocked")
	if err != nil {
		return Memory{}, err
	}

	// SwapTotal
	m.SwapTotalKB, m.SwapTotalGB, err = strToIntPtr(o, "SwapTotal")
	if err != nil {
		return Memory{}, err
	}

	// SwapFree
	m.SwapFreeKB, m.SwapFreeGB, err = strToIntPtr(o, "SwapFree")
	if err != nil {
		return Memory{}, err
	}

	// Dirty
	m.DirtyKB, m.DirtyGB, err = strToIntPtr(o, "Dirty")
	if err != nil {
		return Memory{}, err
	}

	// Writeback
	m.WritebackKB, m.WritebackGB, err = strToIntPtr(o, "Writeback")
	if err != nil {
		return Memory{}, err
	}

	// AnonPages
	m.AnonPagesKB, m.AnonPagesGB, err = strToIntPtr(o, "AnonPages")
	if err != nil {
		return Memory{}, err
	}

	// Mapped
	m.MappedKB, m.MappedGB, err = strToIntPtr(o, "Mapped")
	if err != nil {
		return Memory{}, err
	}

	// Shmem
	m.ShmemKB, m.ShmemGB, err = strToIntPtr(o, "Shmem")
	if err != nil {
		return Memory{}, err
	}

	// Slab
	m.SlabKB, m.SlabGB, err = strToIntPtr(o, "Slab")
	if err != nil {
		return Memory{}, err
	}

	// SReclaimable
	m.SReclaimableKB, m.SReclaimableGB, err = strToIntPtr(o, "SReclaimable")
	if err != nil {
		return Memory{}, err
	}

	// SUnreclaim
	m.SUnreclaimKB, m.SUnreclaimGB, err = strToIntPtr(o, "SUnreclaim")
	if err != nil {
		return Memory{}, err
	}

	// KernelStack
	m.KernelStackKB, m.KernelStackGB, err = strToIntPtr(o, "KernelStack")
	if err != nil {
		return Memory{}, err
	}

	// PageTables
	m.PageTablesKB, m.PageTablesGB, err = strToIntPtr(o, "PageTables")
	if err != nil {
		return Memory{}, err
	}

	// NFS_Unstable
	m.NFSUnstableKB, m.NFSUnstableGB, err = strToIntPtr(o, "NFS_Unstable")
	if err != nil {
		return Memory{}, err
	}

	// Bounce
	m.BounceKB, m.BounceGB, err = strToIntPtr(o, "Bounce")
	if err != nil {
		return Memory{}, err
	}

	// WritebackTmp
	m.WritebackTmpKB, m.WritebackTmpGB, err = strToIntPtr(o, "WritebackTmp")
	if err != nil {
		return Memory{}, err
	}

	// CommitLimit
	m.CommitLimitKB, m.CommitLimitGB, err = strToIntPtr(o, "CommitLimit")
	if err != nil {
		return Memory{}, err
	}

	// Committed_AS
	m.CommittedASKB, m.CommittedASGB, err = strToIntPtr(o, "Committed_AS")
	if err != nil {
		return Memory{}, err
	}

	// VmallocTotal
	m.VmallocTotalKB, m.VmallocTotalGB, err = strToIntPtr(o, "VmallocTotal")
	if err != nil {
		return Memory{}, err
	}

	// VmallocUsed
	m.VmallocUsedKB, m.VmallocUsedGB, err = strToIntPtr(o, "VmalllocUsed")
	if err != nil {
		return Memory{}, err
	}

	// VmallocChunk
	m.VmallocChunkKB, m.VmallocChunkGB, err = strToIntPtr(o, "VmallocChunk")
	if err != nil {
		return Memory{}, err
	}

	// HardwareCorrupted
	m.HardwareCorruptedKB, m.HardwareCorruptedGB, err = strToIntPtr(o, "HardwareCorrupted")
	if err != nil {
		return Memory{}, err
	}

	// AnonHugePages
	m.AnonHugePagesKB, m.AnonHugePagesGB, err = strToIntPtr(o, "AnonHugePages")
	if err != nil {
		return Memory{}, err
	}

	// HugePages_Total
	m.HugePagesTotal, _, err = strToIntPtr(o, "HugePages_Total")
	if err != nil {
		return Memory{}, err
	}

	// HugePages_Free
	m.HugePagesFree, _, err = strToIntPtr(o, "HugePages_Free")
	if err != nil {
		return Memory{}, err
	}

	// HugePages_Rsvd
	m.HugePagesRsvd, _, err = strToIntPtr(o, "HugePages_Rsvd")
	if err != nil {
		return Memory{}, err
	}

	// HugePages_Surp
	m.HugePagesSurp, _, err = strToIntPtr(o, "HugePages_Surp")
	if err != nil {
		return Memory{}, err
	}

	// Hugepagesize
	m.HugePageSizeKB, _, err = strToIntPtr(o, "Hugepagesize")
	if err != nil {
		return Memory{}, err
	}

	// DirectMap4k
	m.DirectMap4kKB, _, err = strToIntPtr(o, "DirectMap4k")
	if err != nil {
		return Memory{}, err
	}

	// DirectMap2M
	m.DirectMap2MKB, _, err = strToIntPtr(o, "DirectMap2M")
	if err != nil {
		return Memory{}, err
	}

	// DirectMap1G
	m.DirectMap1GKB, _, err = strToIntPtr(o, "DirectMap1G")
	if err != nil {
		return Memory{}, err
	}

	if m.CommittedASKB != nil && m.HugePageSizeKB != nil && m.HugePagesTotal != nil && m.TotalKB != nil {
		var usedKB, freeKB int
		var usedGB, freeGB float32

		usedKB = *m.CommittedASKB + *m.HugePageSizeKB**m.HugePagesTotal
		usedGB = float32(usedKB) / 1024 / 1024
		m.CapacityUsedKB = &usedKB
		m.CapacityUsedGB = &usedGB

		freeKB = *m.TotalKB - *m.CapacityUsedKB

		if freeKB < 0 {
			freeKB = 0
		}

		freeGB = float32(freeKB) / 1024 / 1024
		m.CapacityFreeKB = &freeKB
		m.CapacityFreeGB = &freeGB
	}

	// MemAvailable
	// Not always available
	/*
	   m.AvailableKB, m.AvailableGB, err = strToIntPtr(o, "MemAvailable")
	   if err != nil {
	       return Memory{}, err
	   }
	*/
	availableKB := *m.FreeKB + *m.BuffersKB + *m.CachedKB
	usedKB := *m.TotalKB - availableKB

	m.AvailableKB = &availableKB
	m.UsedKB = &usedKB

	availableGB := float32(availableKB) / 1024 / 1024
	usedGB := float32(usedKB) / 1024 / 1024

	m.AvailableGB = &availableGB
	m.UsedGB = &usedGB

	return m, nil
}
