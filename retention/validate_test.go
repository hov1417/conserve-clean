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

func TestSplitBy_1W0s(t *testing.T) {
	backups := []Backup{
		{name: "backup1", relativeTime: 1 * dayInSeconds},
		{name: "backup2", relativeTime: 2 * dayInSeconds},
		{name: "backup3", relativeTime: 30 * dayInSeconds},
		{name: "backup4", relativeTime: (31 + 7) * dayInSeconds},
	}
	policy, err := Parse("1W:0s")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	if err != nil {
		t.Fatal(err)
	}
	if len(keep) != 4 {
		t.Fatalf("expected 4 backups to be kept, got %d", len(keep))
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

func TestSplitBy_1W1W_1(t *testing.T) {
	backups := mockData
	policy, err := Parse("1W:1W")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

func TestSplitBy_1W1W_2(t *testing.T) {
	backups := shiftBySeconds(mockData, 1*dayInSeconds)
	policy, err := Parse("1W:1W")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

func TestSplitBy_1W1W_3(t *testing.T) {
	backups := shiftBySeconds(mockData, 2*dayInSeconds)
	policy, err := Parse("1W:1W")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

func TestSplitBy_1W1W_4(t *testing.T) {
	backups := shiftBySeconds(mockData, 3*dayInSeconds)
	policy, err := Parse("1W:1W")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

func TestSplitBy_1W1W_5(t *testing.T) {
	backups := shiftBySeconds(mockData, 4*dayInSeconds)
	policy, err := Parse("1W:1W")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

func TestSplitBy_1W1W_6(t *testing.T) {
	backups := shiftBySeconds(mockData, 5*dayInSeconds)
	policy, err := Parse("1W:1W")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

func TestSplitBy_1W1W_7(t *testing.T) {
	backups := shiftBySeconds(mockData, 6*dayInSeconds)
	policy, err := Parse("1W:1W")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

var mockData2 = []Backup{
	{name: "backup1", relativeTime: 1 * dayInSeconds},
	{name: "backup1-1", relativeTime: 1*dayInSeconds + 1},
	{name: "backup1-2", relativeTime: 1*dayInSeconds + 2},
	{name: "backup1-3", relativeTime: 1*dayInSeconds + 3},
	{name: "backup1-4", relativeTime: 1*dayInSeconds + 4},
	{name: "backup2", relativeTime: 2 * dayInSeconds},
	{name: "backup2-1", relativeTime: 2*dayInSeconds + 1},
	{name: "backup2-2", relativeTime: 2*dayInSeconds + 2},
	{name: "backup3", relativeTime: 3 * dayInSeconds},
	{name: "backup4", relativeTime: 4 * dayInSeconds},
	{name: "backup5", relativeTime: 5 * dayInSeconds},
	{name: "backup6", relativeTime: 15 * dayInSeconds},
	{name: "backup6-1", relativeTime: 15*dayInSeconds + 1},
	{name: "backup6-2", relativeTime: 15*dayInSeconds + 2},
	{name: "backup7", relativeTime: 16 * dayInSeconds},
}

func TestSplitBy_1W1D(t *testing.T) {
	backups := mockData2

	policy, err := Parse("1W:1D")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}

func TestSplitBy_1W1D_1M1M(t *testing.T) {
	backups := mockData2

	policy, err := Parse("1W:1D,1M:1M")
	if err != nil {
		t.Fatal(err)
	}
	keep, remove, err := splitByPolicy(backups, policy)
	s.MatchSnapshot(t, keep, remove, err)
}
