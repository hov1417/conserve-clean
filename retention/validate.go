package retention

import (
	"github.com/gammazero/deque"
	"github.com/hov1417/conserve-clean/conserve"
	"sort"
	"time"
)

type RawBackup = conserve.RawBackup

type Backup struct {
	name         string
	relativeTime int64
}

func (b Backup) Name() string {
	return b.name
}

// SplitByPolicy Splits a list of backups into two lists, first that satisfy the policy and second that does not.
func SplitByPolicy(backups []RawBackup, policy Policy) ([]Backup, []Backup, error) {
	// sort from newest to oldest
	sort.Slice(backups, func(i, j int) bool {
		return backups[i].Date.Unix() > backups[j].Date.Unix()
	})
	var converted []Backup
	now := time.Now().Unix()
	for _, backup := range backups {
		converted = append(converted, Backup{
			name:         backup.Name,
			relativeTime: now - backup.Date.Unix(),
		})
	}

	return splitByPolicy(converted, policy)
}

func splitByPolicy(backups []Backup, policy Policy) ([]Backup, []Backup, error) {
	var previousEnd int64 = 0
	var generalRemove []Backup
	for _, policyItem := range policy {
		keep, remove := splitByInterval(
			backups,
			previousEnd,
			previousEnd+policyItem.durationSeconds,
			policyItem.intervalSeconds,
		)
		generalRemove = append(generalRemove, remove...)
		backups = keep
		previousEnd = previousEnd + policyItem.durationSeconds
	}

	return backups, generalRemove, nil
}

// splitByInterval Splits a list of backups into two lists,
// first that has at most one backup per interval in range [start, end],
// and second that has to be removed to satisfy previous condition.
//
// The list of backups must be sorted from newest to oldest.
func splitByInterval(backups []Backup, start, end, intervalSeconds int64) ([]Backup, []Backup) {
	//fmt.Println(backups, start, end, intervalSeconds)

	if intervalSeconds == 0 {
		return backups, []Backup{}
	}

	var count int64 = 0
	for _, b := range backups {
		if b.relativeTime >= start && b.relativeTime <= end {
			count++
		}
	}

	if end-start >= intervalSeconds*count {
		return backups, []Backup{}
	}

	var keep []Backup
	var remove []Backup
	var inRange deque.Deque[Backup]
	for _, backup := range backups {
		if backup.relativeTime < start || end < backup.relativeTime {
			keep = append(keep, backup)
		} else {
			inRange.PushBack(backup)
		}
	}
	intervalStart := start
	for intervalStart < end {
		if inRange.Len() == 0 {
			break
		}
		intervalEnd := min(end, intervalStart+intervalSeconds)
		if intervalStart <= inRange.Front().relativeTime &&
			inRange.Front().relativeTime < intervalEnd {

			keep = append(keep, inRange.PopFront())
		}
		for 0 < inRange.Len() && inRange.Front().relativeTime < intervalEnd {
			remove = append(remove, inRange.PopFront())
		}

		intervalStart += intervalSeconds
	}
	// keep might be reordered
	sort.Slice(keep, func(i, j int) bool {
		return keep[i].relativeTime < keep[j].relativeTime
	})

	return keep, remove
}
