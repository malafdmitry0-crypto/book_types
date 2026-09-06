package containers

import "testing"

func TestPopClearsReference(t *testing.T) {
	value := 42
	var s Stack[*int]
	s.Push(&value)
	storage := s.items
	got, ok := s.Pop()
	if !ok || got != &value || s.Len() != 0 {
		t.Fatalf("Pop() = %v, %v; Len() = %d", got, ok, s.Len())
	}
	if storage[0] != nil {
		t.Fatal("popped element is still retained in the backing array")
	}
	if got, ok := s.Pop(); got != nil || ok {
		t.Fatalf("empty Pop() = %v, %v", got, ok)
	}
}
