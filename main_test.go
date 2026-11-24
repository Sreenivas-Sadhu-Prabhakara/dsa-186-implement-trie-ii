package main

import "testing"

func TestCountingTrie(t *testing.T) {
	tr := NewTrie()
	tr.Insert("apple")
	tr.Insert("apple")
	tr.Insert("apps")

	if got := tr.CountWordsEqualTo("apple"); got != 2 {
		t.Errorf("CountWordsEqualTo(apple) = %d, want 2", got)
	}
	if got := tr.CountWordsStartingWith("app"); got != 3 {
		t.Errorf("CountWordsStartingWith(app) = %d, want 3", got)
	}
	if got := tr.CountWordsEqualTo("app"); got != 0 {
		t.Errorf("app was never inserted as a word, got %d", got)
	}

	if !tr.Erase("apple") {
		t.Error("erasing an existing word must succeed")
	}
	if got := tr.CountWordsEqualTo("apple"); got != 1 {
		t.Errorf("after erase: %d, want 1", got)
	}
	if got := tr.CountWordsStartingWith("app"); got != 2 {
		t.Errorf("after erase prefix count: %d, want 2", got)
	}

	if tr.Erase("banana") {
		t.Error("erasing an absent word must fail without corrupting counts")
	}
	if got := tr.CountWordsStartingWith("app"); got != 2 {
		t.Errorf("counts corrupted by failed erase: %d", got)
	}
}
