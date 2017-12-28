package bloom

import (
	"testing"
)

func TestCardinality(t *testing.T) {
	b := New(10, 0.01)
	if b.EstimateSize() != 0 {
		t.Error()
	}

	b.Add([]byte("https://devinmcgloin.com"))
	if b.EstimateSize() == 0 {
		t.Error()
	}

}

func TestMembership(t *testing.T) {
	b := New(6000000, 0.01)
	urls := []string{"https://google.com", "https://devinmcgloin.com", "https://fok.al", "https://twitter.com", "https://facebook.com"}
	for _, url := range urls {
		b.Add([]byte(url))
	}
	for _, url := range urls {
		if !b.Contains([]byte(url)) {
			t.Error()
		}
	}

	if b.Contains([]byte("https://bad.com")) {
		t.Error()
	}
}
