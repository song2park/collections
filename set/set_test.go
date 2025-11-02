package set

import (
	"testing"
)

func TestContains(t *testing.T) {
	s := NewSet("a", "b", "c")

	if !s.Contains("b") {
		t.Errorf("Expected set to contain 'b'")
	}

	if s.Contains("d") {
		t.Errorf("Expected set to not contain 'd'")
	}
}

func TestSize(t *testing.T) {
	s := NewSet(1, 2, 3, 4, 5)

	if s.Size() != 5 {
		t.Errorf("Expected size 5, got %d", s.Size())
	}

	s.Remove(3)

	if s.Size() != 4 {
		t.Errorf("Expected size 4 after removal, got %d", s.Size())
	}
}

func TestAdd(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(2) // Adding duplicate

	if s.Size() != 2 {
		t.Errorf("Expected size 2, got %d", s.Size())
	}
}

func TestAddAll(t *testing.T) {
	s := NewSet(1)
	s.AddAll(2, 3, 4)

	if s.Size() != 4 {
		t.Errorf("Expected size 4 after AddAll, got %d", s.Size())
	}

	for i := 1; i <= 4; i++ {
		if !s.Contains(i) {
			t.Errorf("Expected set to contain %d", i)
		}
	}
}

func TestRemove(t *testing.T) {
	s := NewSet(1, 2, 3)
	s.Remove(2)

	if s.Size() != 2 {
		t.Errorf("Expected size 2 after removal, got %d", s.Size())
	}

	if s.Contains(2) {
		t.Errorf("Expected set to not contain 2 after removal")
	}
}

func TestRemoveAll(t *testing.T) {
	s := NewSet("x", "y", "z", "w")
	s.RemoveAll("y", "w")

	if s.Size() != 2 {
		t.Errorf("Expected size 2 after RemoveAll, got %d", s.Size())
	}

	if s.Contains("y") || s.Contains("w") {
		t.Errorf("Expected set to not contain 'y' or 'w' after RemoveAll")
	}
}

func TestToSlice(t *testing.T) {
	s := NewSet(10, 20, 30)
	slice := s.ToSlice()

	if len(slice) != 3 {
		t.Errorf("Expected slice length 3, got %d", len(slice))
	}

	elementMap := make(map[int]bool)
	for _, v := range slice {
		elementMap[v] = true
	}

	for _, v := range []int{10, 20, 30} {
		if !elementMap[v] {
			t.Errorf("Expected slice to contain %d", v)
		}
	}
}

func TestClear(t *testing.T) {
	s := NewSet(1, 2, 3)
	s.Clear()

	if s.Size() != 0 {
		t.Errorf("Expected size 0 after Clear, got %d", s.Size())
	}

	if s.Contains(1) || s.Contains(2) || s.Contains(3) {
		t.Errorf("Expected set to be empty after Clear")
	}
}