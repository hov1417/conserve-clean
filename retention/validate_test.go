package retention

import (
	"github.com/gkampitakis/go-snaps/snaps"
	"testing"
)

var dayInSeconds int64 = 24 * 60 * 60

var s = snaps.WithConfig(
	snaps.Dir("snaps"),
	snaps.Filename("SplitBackup"),
	snaps.Update(false),
)

func TestSplitBy_7D0s(t *testing.T) {
	backups := []Backup{
		{name: "backup1", relativeTime: 1 * dayInSeconds},
		{name: "backup2", relativeTime: 30 * dayInSeconds},
		{name: "backup3", relativeTime: (31 + 7) * dayInSeconds},
	}
	policy, err := Parse("7D:0s")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	if err != nil {
		t.Fatal(err)
	}
	if len(keep) != 3 {
		t.Fatalf("expected 3 backups to be kept, got %d", len(keep))
	}
	if len(remove) != 0 {
		t.Fatalf("expected 0 backups to be removed, got %d", len(remove))
	}
}

var mockData = []Backup{
	{name: "backup1", relativeTime: 1 * dayInSeconds},
	{name: "backup2", relativeTime: 2 * dayInSeconds},
	{name: "backup3", relativeTime: 3 * dayInSeconds},
	{name: "backup4", relativeTime: 4 * dayInSeconds},
	{name: "backup5", relativeTime: 5 * dayInSeconds},
	{name: "backup6", relativeTime: 15 * dayInSeconds},
	{name: "backup7", relativeTime: 16 * dayInSeconds},
}

func shiftBySeconds(backups []Backup, seconds int64) []Backup {
	var shifted []Backup
	for _, backup := range backups {
		shifted = append(shifted, Backup{
			name:         backup.name,
			relativeTime: backup.relativeTime + seconds,
		})
	}
	return shifted
}

func TestSplitBy_7D7D(t *testing.T) {
	backups := mockData
	policy, err := Parse("7D:7D")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

func TestSplitBy_7D7D_2(t *testing.T) {
	backups := shiftBySeconds(mockData, 1*dayInSeconds)
	policy, err := Parse("7D:7D")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

func TestSplitBy_7D7D_3(t *testing.T) {
	backups := shiftBySeconds(mockData, 2*dayInSeconds)
	policy, err := Parse("7D:7D")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

func TestSplitBy_7D7D_4(t *testing.T) {
	backups := shiftBySeconds(mockData, 3*dayInSeconds)
	policy, err := Parse("7D:7D")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

func TestSplitBy_7D7D(t *testing.T) {
	backups := []Backup{
		{name: "backup1", relativeTime: 1 * dayInSeconds},
		{name: "backup2", relativeTime: 2 * dayInSeconds},
		{name: "backup3", relativeTime: 3 * dayInSeconds},
		{name: "backup4", relativeTime: 4 * dayInSeconds},
		{name: "backup5", relativeTime: 5 * dayInSeconds},
		{name: "backup6", relativeTime: 15 * dayInSeconds},
		{name: "backup7", relativeTime: 16 * dayInSeconds},
	}
	policy, err := Parse("7D:7D")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}
