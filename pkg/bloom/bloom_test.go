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
	urls := []string{"https://evil.com", "https://malicious.com", "https://malware.co", "https://twittter.co", "https://facebok.com"}
	for _, url := range urls {
		b.Add([]byte(url))
	}
	for _, url := range urls {
		if !b.Contains([]byte(url)) {
			t.Error()
		}
	}

	if b.Contains([]byte("https://google.com")) {
		t.Error()
	}
}
