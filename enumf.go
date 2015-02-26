// Package enumf provides simple wrapper functions around a "enumeration flag" type
package enumf

// E is an alias around an int that provides setting and getting
// flag statuses.
type E int

// HasFlag returns true if a given flag value is set.
func (e *E) HasFlag(f E) bool {
	return *e&f != 0
}

// SetFlag sets a flag in a given E instance
func (e *E) SetFlag(f E) {
	*e |= f
}
