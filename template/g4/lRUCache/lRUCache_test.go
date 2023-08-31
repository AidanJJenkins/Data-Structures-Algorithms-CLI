package lrucache

import (
	"testing"
)

func TestLRU(t *testing.T) {
	lru := NewLRU(3)

	if value := lru.Get("foo"); value != nil {
		t.Errorf("Expected value: nil, Got: %v", value)
	}

	lru.Update("foo", 70)
	if value := lru.Get("foo"); value != 70 {
		t.Errorf("Expected value: 70, Got: %v", value)
	}

	lru.Update("bar", 400)
	if value := lru.Get("bar"); value != 400 {
		t.Errorf("Expected value: 400, Got: %v", value)
	}

	lru.Update("baz", 1337)
	if value := lru.Get("baz"); value != 1337 {
		t.Errorf("Expected value: 1337, Got: %v", value)
	}

	lru.Update("ball", 70400)
	if value := lru.Get("ball"); value != 70400 {
		t.Errorf("Expected value: 70400, Got: %v", value)
	}
	if value := lru.Get("foo"); value != nil {
		t.Errorf("Expected value: nil, Got: %v", value)
	}
	if value := lru.Get("bar"); value != 400 {
		t.Errorf("Expected value: 400, Got: %v", value)
	}

	lru.Update("foo", 70)
	if value := lru.Get("bar"); value != 400 {
		t.Errorf("Expected value: 400, Got: %v", value)
	}
	if value := lru.Get("foo"); value != 70 {
		t.Errorf("Expected value: 70, Got: %v", value)
	}

	// Shouldn't have been deleted, but since bar was accessed,
	// bar was added to the front of the list, so baz became the end
	if value := lru.Get("baz"); value != nil {
		t.Errorf("Expected value: nil, Got: %v", value)
	}
}
