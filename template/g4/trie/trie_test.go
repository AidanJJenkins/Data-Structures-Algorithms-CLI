package trie

import (
	"sort"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("foo")
	trie.Insert("fool")
	trie.Insert("foolish")
	trie.Insert("bar")

	found := trie.Find("fo")
	sort.Strings(found)

	expected := []string{"foo", "fool", "foolish"}

	if !isEqual(found, expected) {
		t.Errorf("Expected %v, but got %v", expected, found)
	}

	trie.Delete("fool")

	foundAfterDeletion := trie.Find("fo")
	sort.Strings(foundAfterDeletion)

	expectedAfterDeletion := []string{"foo", "foolish"}

	if !isEqual(foundAfterDeletion, expectedAfterDeletion) {
		t.Errorf("Expected %v, but got %v after deletion", expectedAfterDeletion, foundAfterDeletion)
	}
}

func isEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
