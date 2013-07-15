package zfs

func Snapshot(dataset, name string) error {
	_, err := zfs("snapshot", dataset+"@"+name)
	return err
}

func RecursiveSnapshot(dataset, name string) error {
	_, err := zfs("snapshot", "-r", dataset+"@"+name)
	return err
}
