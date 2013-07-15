package zfs

import (
	"bufio"
	"io"
	"os/exec"
)

func zfs(args ...string) (lines []string, err error) {
	cmd := exec.Command("zfs", args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	r := bufio.NewReader(stdout)
	cmd.Start()

	lines = make([]string, 0, 16)
	for {
		var line string
		line, err = r.ReadString('\n')
		if err == io.EOF {
			return lines, nil
		}
		if err != nil {
			return
		}
		lines = append(lines, line)
	}
	return
}
