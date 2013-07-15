package zfs

import (
	"fmt"
	"strconv"
	"strings"
)

type ListEntry struct {
	// Name of the filesystem, volume or clone in standard pool/fs format.
	Name       string
	// NUmber of bytes used.
	Used       uint64
	// NUmber of bytes available.
	Avail      uint64
	// NUmber of bytes referred to.
	Refer      uint64
	// File system mountpoint or "-".
	Mountpoint string
	// "filesystem", "volume" or "clone"
	Type       string
}

type SnapshotEntry struct {
	Dataset  string
	Snapshot string
	Used     uint64
	Refer    uint64
	Creation int32
}

// ListDatasets lists regular ZFS datasets, i.e. filesystems, volumes and
// clones. Snapshots are not included, similarly to how they are not included
// in "zfs list" by default.
func ListDatasets() ([]ListEntry, error) {
	lines, err := zfs("list", "-pHo", "name,used,avail,refer,mountpoint,type")
	if err != nil {
		return nil, err
	}

	entries := make([]ListEntry, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Fields(string(line))
		if len(fields) != 6 {
			return nil, fmt.Errorf("Unparseable line: %#v", line)
		}

		name := fields[0]
		used, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			return nil, err
		}
		avail, err := strconv.ParseUint(fields[2], 10, 64)
		if err != nil {
			return nil, err
		}
		refer, err := strconv.ParseUint(fields[3], 10, 64)
		if err != nil {
			return nil, err
		}
		mountpoint := fields[4]
		fstype := fields[5]

		e := ListEntry{name, used, avail, refer, mountpoint, fstype}
		entries = append(entries, e)
	}
	return entries, nil
}

// ListSnapshots lists all ZFS snapshots.
func ListSnapshots() ([]SnapshotEntry, error) {
	lines, err := zfs("list", "-pHo", "name,used,refer,creation", "-t", "snapshot")
	if err != nil {
		return nil, err
	}

	entries := make([]SnapshotEntry, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)

		name := fields[0]
		nameFields := strings.SplitN(name, "@", 2)
		used, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			return nil, err
		}
		refer, err := strconv.ParseUint(fields[2], 10, 64)
		if err != nil {
			return nil, err
		}
		creation, err := strconv.ParseInt(fields[3], 10, 32)
		if err != nil {
			return nil, err
		}

		e := SnapshotEntry{nameFields[0], nameFields[1], used, refer, int32(creation)}
		entries = append(entries, e)
	}
	return entries, nil
}
