package zfs_test

import (
	"github.com/calmh/zfs"
	"testing"
)

func TestSnapshotNonexistant(t *testing.T) {
	err := zfs.Snapshot("nonexistant", "foo")
	if err == nil {
		t.Error("Unexpected success for snapshotting nonexistant dataset")
	}
}

func TestSnapshotOk(t *testing.T) {
	err := zfs.Snapshot("zones", "foo")
	if err != nil {
		t.Error("Unexpected error", err)
	}
}
