package storage_test

import (
	"testing"

	"github.com/buunt/storage"
)

func TestTrie(t *testing.T) {
	// the value type is string and the length of the key is 3
	trie := storage.NewTrie[string](3)
	// testing Insert
	if ok := trie.Insert([]byte("key"), "value"); !ok {
		t.Errorf("Output expect 'true' instead of '%v'", ok)
	}
	// testing Find
	if ok, v := trie.Find([]byte("key")); !ok || v != "value" {
		t.Errorf("Output expect 'value' instead of '%v'", v)
	}
	// testing Remove
	if ok := trie.Remove([]byte("key")); !ok {
		t.Errorf("Output expect 'true' instead of '%v'", ok)
	}
}

func TestQueue(t *testing.T) {
	queue := new(storage.Queue[int])
	// testing Push
	queue.Push(100)
	if queue.Len() != 1 {
		t.Errorf("Output expect '1' instead of '%v'", queue.Len())
	}
	// testing Pop
	queue.Pop()
	if queue.Len() != 0 {
		t.Errorf("Output expect '1' instead of '%v'", queue.Len())
	}
}

func TestStack(t *testing.T) {
	stack := new(storage.Stack[int])
	// testing Push
	stack.Push(99)
	stack.Push(100)
	if stack.Len() != 2 {
		t.Errorf("Output expect '2' instead of '%v'", stack.Len())
	}
	// testing Pop
	v := stack.Pop()
	if v != 100 || stack.Len() != 1 {
		t.Errorf("Output expect '100' instead of '%v'", v)
	}
}
