package conserve

import (
	"github.com/gkampitakis/go-snaps/snaps"
	"testing"
)

var s = snaps.WithConfig(
	snaps.Dir("snaps"),
	snaps.Filename("ConserveOutputParse"),
	snaps.Update(false),
)

func TestVersions(t *testing.T) {
	parsed, err := parseVersionsOutput(`b0009                2023-12-17T21:32:21+04:00       0:00
b0010                2023-12-17T21:38:14+04:00       0:00
b0011                2023-12-23T17:57:06+04:00       0:01
b0012                2023-12-23T18:23:16+04:00       0:01
b0013                2023-12-27T14:09:25+04:00       0:01
b0014                2023-12-27T14:10:27+04:00       0:00
b0015                2023-12-28T14:00:01+04:00       0:01
b0016                2023-12-29T00:00:01+04:00       0:02
b0017                2023-12-30T19:00:01+04:00       0:01
b0018                2023-12-31T00:00:01+04:00       0:02
b0019                2024-01-01T12:00:01+04:00       0:01
b0020                2024-01-02T02:00:01+04:00       0:01
b0021                2024-01-03T11:00:01+04:00       0:02
b0022                2024-01-04T00:00:01+04:00       0:00
b0023                2024-01-05T02:00:01+04:00       0:01
b0024                2024-01-06T00:00:01+04:00       0:02`)
	if err != nil {
		t.Fatal(err)
	}
	s.MatchSnapshot(t, parsed)
}
