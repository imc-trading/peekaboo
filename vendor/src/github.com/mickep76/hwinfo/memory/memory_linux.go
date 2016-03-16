// +build linux

package memory

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mickep76/hwinfo/common"
)

type data struct {
	TotalKB             int `json:"total_kb"`
	TotalGB             int `json:"total_gb"`
	FreeKB              int `json:"free_kb"`
	FreeGB              int `json:"free_gb"`
	AvailableKB         int `json:"available_kb"`
	AvailableGB         int `json:"available_gb"`
	BuffersKB           int `json:"buffers_kb"`
	BuffersGB           int `json:"buffers_gb"`
	CachedKB            int `json:"cached_kb"`
	CachedGB            int `json:"cached_gb"`
	SwapCachedKB        int `json:"swap_cached_kb"`
	SwapCachedGB        int `json:"swap_cached_gb"`
	ActiveKB            int `json:"active_kb"`
	ActiveGB            int `json:"active_gb"`
	InactiveKB          int `json:"inactive_kb"`
	InactiveGB          int `json:"inactive_gb"`
	ActiveAnonKB        int `json:"active_anon_kb"`
	ActiveAnonGB        int `json:"active_anon_kb"`
	InactiveAnonKB      int `json:"inactive_anon_buffers_kb"`
	InactiveAnonGB      int `json:"inactive_anon_buffers_kb"`
	ActiveFileKB        int `json:"active_file_kb"`
	ActiveFileGB        int `json:"active_file_gb"`
	InactiveFileKB      int `json:"inactive_file_kb"`
	InactiveFileGB      int `json:"inactive_file_gb"`
	UnevictableKB       int `json:"unevictable_kb"`
	UnevictableGB       int `json:"unevictable_gb"`
	MLockedKB           int `json:"m_locked_kb"`
	MLockedGB           int `json:"m_locked_gb"`
	SwapTotalKB         int `json:"swap_total_kb"`
	SwapTotalGB         int `json:"swap_total_gb"`
	SwapFreeKB          int `json:"swap_free_kb"`
	SwapFreeGB          int `json:"swap_free_gb"`
	DirtyKB             int `json:"dirty_kb"`
	DirtyGB             int `json:"dirty_gb"`
	WritebackKB         int `json:"writeback_kb"`
	WritebackGB         int `json:"writeback_gb"`
	AnonPagesKB         int `json:"anon_pages_kb"`
	AnonPagesGB         int `json:"anon_pages_gb"`
	MappedKB            int `json:"mapped_kb"`
	MappedGB            int `json:"mapped_gb"`
	ShmemKB             int `json:"shmem_kb"`
	ShmemGB             int `json:"shmem_gb"`
	SlabKB              int `json:"slab_kb"`
	SlabGB              int `json:"slab_gb"`
	SReclaimableKB      int `json:"s_reclaimable_kb"`
	SReclaimableGB      int `json:"s_reclaimable_gb"`
	SUnreclaimKB        int `json:"s_unreclaim_kb"`
	SUnreclaimGB        int `json:"s_unreclaim_gb"`
	KernelStackKB       int `json:"kernel_stack_kb"`
	KernelStackGB       int `json:"kernel_stack_gb"`
	PageTablesKB        int `json:"page_tables_kb"`
	PageTablesGB        int `json:"page_tables_gb"`
	NFSUnstableKB       int `json:"nfs_unstable_kb"`
	NFSUnstableGB       int `json:"nfs_unstable_gb"`
	BounceKB            int `json:"bounce_kb"`
	BounceGB            int `json:"bounce_gb"`
	WritebackTmpKB      int `json:"writeback_tmp_kb"`
	WritebackTmpGB      int `json:"writeback_tmp_gb"`
	CommitLimitKB       int `json:"commit_limit_kb"`
	CommitLimitGB       int `json:"commit_limit_gb"`
	CommittedASKB       int `json:"committed_a_s_kb"`
	CommittedASGB       int `json:"committed_a_s_gb"`
	VmallocTotalKB      int `json:"vmalloc_total_kb"`
	VmallocTotalGB      int `json:"vmalloc_total_gb"`
	VmallocUsedKB       int `json:"vmalloc_used_kb"`
	VmallocUsedGB       int `json:"vmalloc_used_gb"`
	VmallocChunkKB      int `json:"vmalloc_chunk_kb"`
	VmallocChunkGB      int `json:"vmalloc_chunk_gb"`
	HardwareCorruptedKB int `json:"hardware_corrupted_kb"`
	HardwareCorruptedGB int `json:"hardware_corrupted_gb"`
	AnonHugePagesKB     int `json:"anon_huge_pages_kb"`
	AnonHugePagesGB     int `json:"anon_huge_pages_gb"`
	HugePagesTot        int `json:"huge_pages_tot"`
	HugePagesFree       int `json:"huge_pages_free"`
	HugePagesRsvd       int `json:"huge_pages_rsvd"`
	HugePagesSurp       int `json:"huge_pages_surp"`
	HugePageSizeKB      int `json:"huge_pages_size_kb"`
	DirectMap4kKB       int `json:"direct_map_4k_kb"`
	DirectMap2MKB       int `json:"direct_map_2m_kb"`
	DirectMap1GKB       int `json:"direct_map_1g_kb"`
}

func (m *memory) ForceUpdate() error {
	m.cache.LastUpdated = time.Now()
	m.cache.FromCache = false

	o, err := common.LoadFileFields("/proc/meminfo", ":", []string{
		"MemTotal",
		"MemFree",
		"MemAvailable",
		"Buffers",
		"Cached",
		"SwapCached",
		"Active",
		"Inactive",
		"Active(anon)",
		"Inactive(anon)",
		"Active(file)",
		"Inactive(file)",
		"Unevictable",
		"Mlocked",
		"SwapTotal",
		"SwapFree",
		"Dirty",
		"Writeback",
		"AnonPages",
		"Mapped",
		"Shmem",
		"Slab",
		"SReclaimable",
		"SUnreclaim",
		"KernelStack",
		"PageTables",
		"NFS_Unstable",
		"Bounce",
		"WritebackTmp",
		"CommitLimit",
		"Committed_AS",
		"VmallocTotal",
		"VmallocUsed",
		"VmallocChunk",
		"HardwareCorrupted",
		"AnonHugePages",
		"HugePages_Total",
		"HugePages_Free",
		"HugePages_Rsvd",
		"HugePages_Surp",
		"Hugepagesize",
		"DirectMap4k",
		"DirectMap2M",
		"DirectMap1G",
	})
	if err != nil {
		return err
	}

	// MemTotal
	m.data.TotalKB, err = strconv.Atoi(strings.TrimRight(o["MemTotal"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "MemTotal", err.Error())
	}
	m.data.TotalGB = m.data.TotalKB / 1024 / 1024

	// MemFree
	m.data.FreeKB, err = strconv.Atoi(strings.TrimRight(o["MemFree"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "MemFree", err.Error())
	}
	m.data.FreeGB = m.data.FreeKB / 1024 / 1024

	// MemAvailable
	m.data.AvailableKB, err = strconv.Atoi(strings.TrimRight(o["MemAvailable"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "MemAvailable", err.Error())
	}
	m.data.AvailableGB = m.data.AvailableKB / 1024 / 1024

	// Buffers
	m.data.BuffersKB, err = strconv.Atoi(strings.TrimRight(o["Buffers"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Buffers", err.Error())
	}
	m.data.BuffersGB = m.data.BuffersKB / 1024 / 1024

	// Cached
	m.data.CachedKB, err = strconv.Atoi(strings.TrimRight(o["Cached"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Cached", err.Error())
	}
	m.data.CachedGB = m.data.CachedKB / 1024 / 1024

	// SwapCached
	m.data.SwapCachedKB, err = strconv.Atoi(strings.TrimRight(o["SwapCached"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "SwapCached", err.Error())
	}
	m.data.SwapCachedGB = m.data.SwapCachedKB / 1024 / 1024

	// Active
	m.data.ActiveKB, err = strconv.Atoi(strings.TrimRight(o["Active"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Active", err.Error())
	}
	m.data.ActiveGB = m.data.ActiveKB / 1024 / 1024

	// Inactive
	m.data.InactiveKB, err = strconv.Atoi(strings.TrimRight(o["Inactive"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Inactive", err.Error())
	}
	m.data.InactiveGB = m.data.InactiveKB / 1024 / 1024

	// Active(anon)
	m.data.ActiveAnonKB, err = strconv.Atoi(strings.TrimRight(o["Active(anon)"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Active(anon)", err.Error())
	}
	m.data.ActiveAnonGB = m.data.ActiveAnonKB / 1024 / 1024

	// Inactive(anon)
	m.data.InactiveAnonKB, err = strconv.Atoi(strings.TrimRight(o["Inactive(anon)"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Inactive(anon)", err.Error())
	}
	m.data.InactiveAnonGB = m.data.InactiveAnonKB / 1024 / 1024

	// Active(file)
	m.data.ActiveFileKB, err = strconv.Atoi(strings.TrimRight(o["Active(file)"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Active(file)", err.Error())
	}
	m.data.ActiveFileGB = m.data.ActiveFileKB / 1024 / 1024

	// Inactive(file)
	m.data.InactiveFileKB, err = strconv.Atoi(strings.TrimRight(o["Inactive(file)"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Inactive(file)", err.Error())
	}
	m.data.InactiveFileGB = m.data.InactiveFileKB / 1024 / 1024

	// Unevictable
	m.data.UnevictableKB, err = strconv.Atoi(strings.TrimRight(o["Unevictable"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Unevictable", err.Error())
	}
	m.data.UnevictableGB = m.data.UnevictableKB / 1024 / 1024

	// Mlocked
	m.data.MLockedKB, err = strconv.Atoi(strings.TrimRight(o["Mlocked"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Mlocked", err.Error())
	}
	m.data.MLockedGB = m.data.MLockedKB / 1024 / 1024

	// SwapTotal
	m.data.SwapTotalKB, err = strconv.Atoi(strings.TrimRight(o["SwapTotal"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "SwapTotal", err.Error())
	}
	m.data.SwapTotalGB = m.data.SwapTotalKB / 1024 / 1024

	// SwapFree
	m.data.SwapFreeKB, err = strconv.Atoi(strings.TrimRight(o["SwapFree"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "SwapFree", err.Error())
	}
	m.data.SwapFreeGB = m.data.SwapFreeKB / 1024 / 1024

	// Dirty
	m.data.DirtyKB, err = strconv.Atoi(strings.TrimRight(o["Dirty"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Dirty", err.Error())
	}
	m.data.DirtyGB = m.data.DirtyKB / 1024 / 1024

	// Writeback
	m.data.WritebackKB, err = strconv.Atoi(strings.TrimRight(o["Writeback"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Writeback", err.Error())
	}
	m.data.WritebackGB = m.data.WritebackKB / 1024 / 1024

	// AnonPages
	m.data.AnonPagesKB, err = strconv.Atoi(strings.TrimRight(o["AnonPages"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "AnonPages", err.Error())
	}
	m.data.AnonPagesGB = m.data.AnonPagesKB / 1024 / 1024

	// Mapped
	m.data.MappedKB, err = strconv.Atoi(strings.TrimRight(o["Mapped"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Mapped", err.Error())
	}
	m.data.MappedGB = m.data.MappedKB / 1024 / 1024

	// Shmem
	m.data.ShmemKB, err = strconv.Atoi(strings.TrimRight(o["Shmem"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Shmem", err.Error())
	}
	m.data.ShmemGB = m.data.ShmemKB / 1024 / 1024

	// Slab
	m.data.SlabKB, err = strconv.Atoi(strings.TrimRight(o["Slab"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Slab", err.Error())
	}
	m.data.SlabGB = m.data.SlabKB / 1024 / 1024

	// SReclaimable
	m.data.SReclaimableKB, err = strconv.Atoi(strings.TrimRight(o["SReclaimable"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "SReclaimable", err.Error())
	}
	m.data.SReclaimableGB = m.data.SReclaimableKB / 1024 / 1024

	// SUnreclaim
	m.data.SUnreclaimKB, err = strconv.Atoi(strings.TrimRight(o["SUnreclaim"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "SUnreclaim", err.Error())
	}
	m.data.SUnreclaimGB = m.data.SUnreclaimKB / 1024 / 1024

	// KernelStack
	m.data.KernelStackKB, err = strconv.Atoi(strings.TrimRight(o["KernelStack"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "KernelStack", err.Error())
	}
	m.data.KernelStackGB = m.data.KernelStackKB / 1024 / 1024

	// PageTables
	m.data.PageTablesKB, err = strconv.Atoi(strings.TrimRight(o["PageTables"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "PageTables", err.Error())
	}
	m.data.PageTablesGB = m.data.PageTablesKB / 1024 / 1024

	// NFS_Unstable
	m.data.NFSUnstableKB, err = strconv.Atoi(strings.TrimRight(o["NFS_Unstable"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "NFS_Unstable", err.Error())
	}
	m.data.NFSUnstableGB = m.data.NFSUnstableKB / 1024 / 1024

	// Bounce
	m.data.BounceKB, err = strconv.Atoi(strings.TrimRight(o["Bounce"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Bounce", err.Error())
	}
	m.data.BounceGB = m.data.BounceKB / 1024 / 1024

	// WritebackTmp
	m.data.WritebackTmpKB, err = strconv.Atoi(strings.TrimRight(o["WritebackTmp"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "WritebackTmp", err.Error())
	}
	m.data.WritebackTmpGB = m.data.WritebackTmpKB / 1024 / 1024

	// CommitLimit
	m.data.CommitLimitKB, err = strconv.Atoi(strings.TrimRight(o["CommitLimit"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "CommitLimit", err.Error())
	}
	m.data.CommitLimitGB = m.data.CommitLimitKB / 1024 / 1024

	// Committed_AS
	m.data.CommittedASKB, err = strconv.Atoi(strings.TrimRight(o["Committed_AS"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Committed_AS", err.Error())
	}
	m.data.CommittedASGB = m.data.CommittedASKB / 1024 / 1024

	// VmallocTotal
	m.data.VmallocTotalKB, err = strconv.Atoi(strings.TrimRight(o["VmallocTotal"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "VmallocTotal", err.Error())
	}
	m.data.VmallocTotalGB = m.data.VmallocTotalKB / 1024 / 1024

	// VmallocUsed
	m.data.VmallocUsedKB, err = strconv.Atoi(strings.TrimRight(o["VmallocUsed"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "VmallocUsed", err.Error())
	}
	m.data.VmallocUsedGB = m.data.VmallocUsedKB / 1024 / 1024

	// VmallocChunk
	m.data.VmallocChunkKB, err = strconv.Atoi(strings.TrimRight(o["VmallocChunk"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "VmallocChunk", err.Error())
	}
	m.data.VmallocChunkGB = m.data.VmallocChunkKB / 1024 / 1024

	// HardwareCorrupted
	m.data.HardwareCorruptedKB, err = strconv.Atoi(strings.TrimRight(o["HardwareCorrupted"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "HardwareCorrupted", err.Error())
	}
	m.data.HardwareCorruptedGB = m.data.HardwareCorruptedKB / 1024 / 1024

	// AnonHugePages
	m.data.AnonHugePagesKB, err = strconv.Atoi(strings.TrimRight(o["AnonHugePages"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "AnonHugePages", err.Error())
	}
	m.data.AnonHugePagesGB = m.data.AnonHugePagesKB / 1024 / 1024

	// HugePages_Total
	m.data.HugePagesTot, err = strconv.Atoi(o["HugePages_Total"])
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "HugePages_Total", err.Error())
	}

	// HugePages_Free
	m.data.HugePagesFree, err = strconv.Atoi(o["HugePages_Free"])
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "HugePages_Free", err.Error())
	}

	// HugePages_Rsvd
	m.data.HugePagesRsvd, err = strconv.Atoi(o["HugePages_Rsvd"])
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "HugePages_Rsvd", err.Error())
	}

	// Hugepagesize
	m.data.HugePageSizeKB, err = strconv.Atoi(strings.TrimRight(o["Hugepagesize"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "Hugepagesize", err.Error())
	}

	// DirectMap4k
	m.data.DirectMap4kKB, err = strconv.Atoi(strings.TrimRight(o["DirectMap4k"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "DirectMap4k", err.Error())
	}

	// DirectMap2M
	m.data.DirectMap2MKB, err = strconv.Atoi(strings.TrimRight(o["DirectMap2M"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "DirectMap2M", err.Error())
	}

	// DirectMap1G
	m.data.DirectMap1GKB, err = strconv.Atoi(strings.TrimRight(o["DirectMap1G"], " kB"))
	if err != nil {
		return fmt.Errorf("failed parsing field: %s error: %s", "DirectMap1G", err.Error())
	}

	return nil
}
