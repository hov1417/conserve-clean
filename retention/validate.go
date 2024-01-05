package retention

import (
	"github.com/gammazero/deque"
	"sort"
	"time"
)

type RawBackup struct {
	name string
	date time.Time
}

type Backup struct {
	name         string
	relativeTime int64
}

// SplitByPolicy Splits a list of backups into two lists, first that satisfy the policy and second that does not.
func SplitByPolicy(backups []RawBackup, policy Policy) ([]Backup, []Backup, error) {
	// sort from newest to oldest
	sort.Slice(backups, func(i, j int) bool {
		return backups[i].date.Unix() > backups[j].date.Unix()
	})
	var converted []Backup
	now := time.Now().Unix()
	for _, backup := range backups {
		converted = append(converted, Backup{
			name:         backup.name,
			relativeTime: now - backup.date.Unix(),
		})
	}

	return splitByPolicy(converted, policy)
}

func splitByPolicy(backups []Backup, policy Policy) ([]Backup, []Backup, error) {
	var previousEnd int64 = 0
	var generalRemove []Backup
	for _, policyItem := range policy {
		keep, remove := splitByInterval2(
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

// splitByInterval2 Splits a list of backups into two lists,
// first that has at most one backup per interval in range [start, end],
// and second that has to be removed to satisfy previous condition.
//
// The list of backups must be sorted from newest to oldest.
func splitByInterval2(backups []Backup, start, end, intervalSeconds int64) ([]Backup, []Backup) {
	if intervalSeconds == 0 {
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

//func splitByInterval(backups []RawBackup, start, end, intervalSeconds int64) ([]RawBackup, []RawBackup) {
//	var keep []RawBackup
//	var remove []RawBackup
//	oneInTheInterval := false
//	for _, backup := range backups {
//		if backup.date.Unix() < start && end < backup.date.Unix() {
//			keep = append(keep, backup)
//		} else {
//			if oneInTheInterval {
//				lastInInterval := keep[len(keep)-1]
//				if lastInInterval.date.Unix()-backup.date.Unix() < intervalSeconds {
//					remove = append(remove, backup)
//					oneInTheInterval = false
//				} else {
//					keep = append(keep, backup)
//				}
//			} else {
//				keep = append(keep, backup)
//				oneInTheInterval = true
//			}
//		}
//	}
//	return keep, remove
//}
