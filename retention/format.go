package retention

import (
	"fmt"
	"sort"
	"strings"
)

type Unit byte

const (
	Second Unit = 's'
	Minute Unit = 'm'
	Hour   Unit = 'h'
	Day    Unit = 'D'
	Week   Unit = 'W'
	Month  Unit = 'M'
	Year   Unit = 'Y'
)

func (u Unit) asSeconds() int64 {
	switch u {
	case Second:
		return 1
	case Minute:
		return 60 * Second.asSeconds()
	case Hour:
		return 60 * Minute.asSeconds()
	case Day:
		return 24 * Hour.asSeconds()
	case Week:
		return 7 * Day.asSeconds()
	case Month:
		return 30 * Day.asSeconds()
	case Year:
		return 365 * Day.asSeconds()
	}
	panic("invalid unit " + string(u))
}

type PolicyItem struct {
	durationSeconds int64
	intervalSeconds int64
}

func (p PolicyItem) Less(other PolicyItem) bool {
	return p.durationSeconds < other.durationSeconds
}

type Policy []PolicyItem

func (p Policy) Len() int {
	return len(p)
}

func (p Policy) Less(i, j int) bool {
	return p[i].Less(p[j])
}

func (p Policy) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func parseItem(str string) (PolicyItem, error) {
	parts := strings.Split(str, ":")
	if len(parts) != 2 {
		return PolicyItem{}, fmt.Errorf("invalid policy item format")
	}
	var duration int64
	var unit Unit
	var interval int64
	var intervalUnit Unit
	if _, err := fmt.Sscanf(parts[0], "%d%c", &duration, &unit); err != nil {
		return PolicyItem{}, fmt.Errorf("invalid duration format: %s", err)
	}
	if _, err := fmt.Sscanf(parts[1], "%d%c", &interval, &intervalUnit); err != nil {
		return PolicyItem{}, fmt.Errorf("invalid interval format: %s", err)
	}
	return PolicyItem{
		durationSeconds: duration * unit.asSeconds(),
		intervalSeconds: interval * intervalUnit.asSeconds(),
	}, nil
}

func Parse(str string) (Policy, error) {
	parts := strings.Split(str, ",")
	var policy Policy
	for _, part := range parts {
		item, err := parseItem(part)
		if err != nil {
			return nil, err
		}
		policy = append(policy, item)
	}
	sort.Sort(policy)
	return policy, nil
}
