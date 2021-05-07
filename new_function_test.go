package cache

import (
	"testing"
	"time"
)

func TestCache_Pop(t *testing.T) {
	tc := New(DefaultExpiration, 0)
	tc.Set("name1", "Tom", DefaultExpiration)
	tc.Set("name2", "Bob", time.Second)

	_, found := tc.Pop("name1")
	if !found {
		t.Errorf("name1 %s not found", "Tom")
	}

	_, found = tc.Get("name1")
	if found {
		t.Errorf("name1 %s not removed", "Tom")
	}

	time.Sleep(time.Second)
	_, found = tc.Pop("name2")
	if found {
		t.Errorf("name2 %s not removed", "Bob")
	}
}
