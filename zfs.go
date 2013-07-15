package zfs

import (
	"os/exec"
	"strings"
)

func zfs(args ...string) (lines []string, err error) {
	cmd := exec.Command("zfs", args...)
	bytes, err := cmd.Output()
	if err != nil {
		return
	}

	tmpLines := strings.Split(string(bytes), "\n")
	lines = make([]string, 0, len(lines))
	for _, line := range tmpLines {
		if line != "" {
			lines = append(lines, line)
		}
	}
	return
}
