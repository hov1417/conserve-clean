package conserve

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type RawBackup struct {
	Name string
	Date time.Time
}

func Versions(executable string, dir string) ([]RawBackup, error) {
	outTable, err := executeVersions(executable, dir)
	if err != nil {
		return []RawBackup{}, err
	}
	return parseVersionsOutput(outTable)
}

func executeVersions(executable string, dir string) (string, error) {
	cmd := exec.Command(executable, "versions", dir)
	if cmd.Stdout != nil {
		return "", fmt.Errorf("exec: Stdout already set")
	}
	if cmd.Stderr != nil {
		return "", fmt.Errorf("exec: Stderr already set")
	}
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Error on conserve versions execution\n%s\n%s", err, stderrBuf.String())
	}
	return stdoutBuf.String(), nil
}

func parseVersionsOutput(outTable string) ([]RawBackup, error) {
	rows := strings.Split(outTable, "\n")
	var backups []RawBackup
	for _, row := range rows {
		if row == "" {
			continue
		}
		fields := strings.Fields(row)
		if len(fields) <= 2 {
			return []RawBackup{}, fmt.Errorf("invalid output format\n %s", row)
		}
		name := fields[0]
		date, err := time.Parse(time.RFC3339, fields[1])
		if err != nil {
			return []RawBackup{}, err
		}
		backups = append(backups, RawBackup{
			Name: name,
			Date: date,
		})
	}
	return backups, nil
}
