package conserve

import (
	"fmt"
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
	return execute(executable, "versions", dir)
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
