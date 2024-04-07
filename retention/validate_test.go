package retention

import (
	"fmt"
	"github.com/gkampitakis/go-snaps/snaps"
	"reflect"
	"strconv"
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

func TestRunningIntervalShouldNotAffectResult0(t *testing.T) {
	policy, err := Parse("1W:1D,1M:3D,1Y:2W,10Y:2M")
	if err != nil {
		t.Fatal(err)
	}

	// run every 3 days
	backups1, err := runEveryNDays(3, 10, policy)
	if err != nil {
		t.Fatal(err)
	}
	// run once at the end
	backups2, err := runEveryNDays(9, 10, policy)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(backups1, backups2) {
		fmt.Println("backups1", backups1)
		fmt.Println("backups2", backups2)
		t.Fatal("backups1 != backups2")
	}
}

func TestRunningIntervalShouldNotAffectResult1(t *testing.T) {
	policy, err := Parse("2D:1D,6D:2D")
	if err != nil {
		t.Fatal(err)
	}

	runningInterval3n(t, 4, policy)
	runningInterval3n(t, 5, policy)
	runningInterval3n(t, 6, policy)
	runningInterval3n(t, 7, policy)
	runningInterval3n(t, 8, policy)
	runningInterval3n(t, 9, policy)
	runningInterval3n(t, 10, policy)
	runningInterval3n(t, 11, policy)
	runningInterval3n(t, 12, policy)
	runningInterval3n(t, 13, policy)
	runningInterval3n(t, 14, policy)
	runningInterval3n(t, 15, policy)
	runningInterval3n(t, 16, policy)
	runningInterval3n(t, 17, policy)
}

func runningInterval3n(t *testing.T, n int, policy Policy) {
	//fmt.Println("runningInterval3n", n)
	// run every 3 days
	backups1, err := runEveryNDays(3, n, policy)
	if err != nil {
		t.Fatal(err)
	}
	// run once at the end
	backups2, err := runEveryNDays(n-1, n, policy)
	if err != nil {
		t.Fatal(err)
	}
	if len(backups1) < len(backups2) {
		fmt.Println("backups1", backups1)
		fmt.Println("backups2", backups2)
		t.Fatal("backups1 != backups2")
	}
}

func TestRunningIntervalShouldNotAffectResult2(t *testing.T) {
	policy, err := Parse("1W:1D,1M:3D,1Y:2W,10Y:2M")
	if err != nil {
		t.Fatal(err)
	}

	// run every 3 days
	backups1, err := runEveryNDays(3, 30, policy)
	if err != nil {
		t.Fatal(err)
	}
	// run once at the end
	backups2, err := runEveryNDays(29, 30, policy)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(backups1, backups2) {
		fmt.Println("backups1", backups1)
		fmt.Println("backups2", backups2)
		t.Fatal("backups1 != backups2")
	}
}

func TestRunningIntervalShouldNotAffectResult3(t *testing.T) {
	policy, err := Parse("1W:1D,1M:3D,1Y:2W,10Y:2M")
	if err != nil {
		t.Fatal(err)
	}

	// run every 3 days
	backups1, err := runEveryNDays(3, 365, policy)
	if err != nil {
		t.Fatal(err)
	}
	// run every month
	backups2, err := runEveryNDays(30, 365, policy)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(backups1, backups2) {
		fmt.Println("backups1", backups1)
		fmt.Println("backups2", backups2)
		t.Fatal("backups1 != backups2")
	}
}

func runEveryNDays(n, numberOfDays int, policy Policy) ([]Backup, error) {
	var backups []Backup
	for i := 0; i < numberOfDays; i++ {
		backups = AddADayToAll(backups)
		backups = append([]Backup{{
			name:         "backup" + strconv.Itoa(i),
			relativeTime: 0,
		}}, backups...)
		//fmt.Println(backups)
		if i%n == 0 {
			keep, _, err := splitByPolicy(backups, policy)
			if err != nil {
				return nil, err
			}
			backups = keep
		}
	}

	return backups, nil
}

func AddADayToAll(backups []Backup) []Backup {
	var shifted []Backup
	for _, backup := range backups {
		shifted = append(shifted, Backup{
			name:         backup.name,
			relativeTime: backup.relativeTime + dayInSeconds,
		})
	}
	return shifted
}
