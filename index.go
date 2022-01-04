package mpsl

// Index stores hashed key-entry pairs.
//
// Hashed key uses to reduce pointers in the package to follow pointers policy.
// Entry as uint64 value uses due to impossibility to take a pointer of map value.
type index map[uint64]entry

// Save new entry.
func (i *index) set(key uint64, lo, hi uint32, icann bool, typ uint8) {
	var (
		e entry
		f uint8
	)
	if icann {
		f = 1
	}
	e.encode(lo, hi, f, typ)
	(*i)[key] = e
}

// Get entry by given key.
func (i index) get(key uint64) entry {
	if e, ok := i[key]; ok {
		return e
	}
	return 0
}

// Remove all keys from index.
func (i *index) reset() {
	for h := range *i {
		delete(*i, h)
	}
}
