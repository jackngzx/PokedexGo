package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	if cache.entries == nil {
		t.Error("Expected cache entries to be initialized")
	}
	if cache.interval != time.Minute*5 {
		t.Errorf("Expected interval to be %v, got %v", time.Minute*5, cache.interval)
	}
}

func TestAdd(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	key := "test"
	val := []byte("value")
	cache.Add(key, val)

	entry, exists := cache.entries[key]
	if !exists {
		t.Error("Expected entry to exist after Add")
	}
	if string(entry.val) != "value" {
		t.Errorf("Expected val to be 'value', got %s", string(entry.val))
	}
}

func TestGet(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	key := "test"
	val := []byte("value")
	cache.Add(key, val)

	retrieved, ok := cache.Get(key)
	if !ok {
		t.Error("Expected to get value for existing key")
	}
	if string(retrieved) != "value" {
		t.Errorf("Expected retrieved value to be 'value', got %s", string(retrieved))
	}
}

func TestGetNonExisting(t *testing.T) {
	cache := NewCache(time.Minute * 5)
	_, ok := cache.Get("nonexisting")
	if ok {
		t.Error("Expected false for non-existing key")
	}
}

func TestReapLoop(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	key := "test"
	val := []byte("value")
	cache.Add(key, val)

	// Wait for reaping
	time.Sleep(interval * 2)

	_, ok := cache.Get(key)
	if ok {
		t.Error("Expected entry to be reaped")
	}
}
