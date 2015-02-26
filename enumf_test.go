package enumf

import "testing"

const (
	// Permissions
	None E = 1 << iota
	Read
	Write
	Admin = 32768
)

func TestHasFlagReturnsTrue(t *testing.T) {
	e := Read | Write
	if !e.HasFlag(Read) || !e.HasFlag(Write) {
		t.Fatal("Failed")
	}

	if e.HasFlag(Admin) {
		t.Fatal("Failed")
	}
}

func TestSetFlag(t *testing.T) {
	e := E(0)

	e.SetFlag(Write)

	if !e.HasFlag(Write) {
		t.Fatal("Failed")
	}

	if e.HasFlag(Read) || e.HasFlag(Admin) {
		t.Fatal("Failed")
	}
}

func TestMultiSetFlag(t *testing.T) {
	e := E(0)

	e.SetFlag(Write)
	e.SetFlag(Read)

	if !e.HasFlag(Write) || !e.HasFlag(Read) {
		t.Fatal("Failed")
	}

	if e.HasFlag(Admin) {
		t.Fatal("Failed")
	}
}

func TestIntegerComparison(t *testing.T) {
	all := 32771

	e := E(Read | Write | Admin)

	if !e.HasFlag(E(all)) {
		t.Fatal("failed")
	}
}
