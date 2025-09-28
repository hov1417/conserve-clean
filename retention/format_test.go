package retention

import (
	"testing"
)

func Test1Parse(t *testing.T) {
	res, err := Parse("7D:0s")
	if err != nil {
		t.Errorf("Error parsing policy: %s", err)
	}
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
	if res[0].durationSeconds != 7*24*60*60 {
		t.Errorf("Expected 7 days, got %d", res[0].durationSeconds)
	}
	if res[0].intervalSeconds != 0 {
		t.Errorf("Expected 0 seconds, got %d", res[0].intervalSeconds)
	}
}

func Test2Parse(t *testing.T) {
	res, err := Parse("7D:0s,3M:1D")
	if err != nil {
		t.Errorf("Error parsing policy: %s", err)
	}
	if len(res) != 2 {
		t.Errorf("Expected 2 results, got %d", len(res))
	}
	if res[0].durationSeconds != 7*24*60*60 {
		t.Errorf("Expected 7 days, got %d", res[0].durationSeconds)
	}
	if res[0].intervalSeconds != 0 {
		t.Errorf("Expected 0 seconds, got %d", res[0].intervalSeconds)
	}
	if res[1].durationSeconds != 3*30*24*60*60 {
		t.Errorf("Expected 7 days, got %d", res[0].durationSeconds)
	}
	if res[1].intervalSeconds != 24*60*60 {
		t.Errorf("Expected 0 seconds, got %d", res[0].intervalSeconds)
	}
}

func TestParse(t *testing.T) {
	res, err := Parse("7D:0s,3M:1D,10Y:2M")
	if err != nil {
		t.Errorf("Error parsing policy: %s", err)
	}
	if len(res) != 3 {
		t.Errorf("Expected 3 results, got %d", len(res))
	}
	if res[0].durationSeconds != 7*24*60*60 {
		t.Errorf("Expected 7 days, got %d", res[0].durationSeconds)
	}
	if res[0].intervalSeconds != 0 {
		t.Errorf("Expected 0 seconds, got %d", res[0].intervalSeconds)
	}
	if res[1].durationSeconds != 3*30*24*60*60 {
		t.Errorf("Expected 7 days, got %d", res[0].durationSeconds)
	}
	if res[1].intervalSeconds != 24*60*60 {
		t.Errorf("Expected 24*60*60 seconds, got %d", res[0].intervalSeconds)
	}
	if res[2].durationSeconds != 10*365*24*60*60 {
		t.Errorf("Expected 7 days, got %d", res[0].durationSeconds)
	}
	if res[2].intervalSeconds != 2*30*24*60*60 {
		t.Errorf("Expected 2*30*24*60*60 seconds, got %d", res[0].intervalSeconds)
	}
}

func TestHigherIntervalThanDuration(t *testing.T) {
	res, err := Parse("7D:8D")
	if err != nil {
		t.Errorf("Error parsing policy: %s", err)
	}
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
	if res[0].durationSeconds != 7*24*60*60 {
		t.Errorf("Expected 7 days, got %d", res[0].durationSeconds)
	}
	if res[0].intervalSeconds != 8*24*60*60 {
		t.Errorf("Expected 8 days, got %d", res[0].intervalSeconds)
	}
}

func TestEqualIntervalAndDuration(t *testing.T) {
	res, err := Parse("7D:7D")
	if err != nil {
		t.Errorf("Error parsing policy: %s", err)
	}
	if len(res) != 1 {
		t.Errorf("Expected 1 result, got %d", len(res))
	}
	if res[0].durationSeconds != 7*24*60*60 {
		t.Errorf("Expected 7 days, got %d", res[0].durationSeconds)
	}
	if res[0].intervalSeconds != 7*24*60*60 {
		t.Errorf("Expected 7 days, got %d", res[0].intervalSeconds)
	}
}
