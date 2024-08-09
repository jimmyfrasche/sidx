// Package sidx includes a number of helpers to extend slice indexing to negative indices.
//
// [Index] calculates positive and negative indices
// allowing 0 for the first element as usual or -1 for the last
// (and -2 for the second to last, etc.).
// [N] is a more convenient form, for the common case of a slice.
// Everything else in this package is a wrapper around N for common cases.
//
// [At](s, i) is s[N(s, i)] and [End] is short for At(s, -1).
// [AtOk] is a the comma-ok version of At.
// [AtOr] is similar to AtOk but it returns a default value when the index is invalid.
//
// [Slice] is a 2 value slice where the start and end are calculated by N.
//
// [Pop](s) returns, clears, and slices off At(s, -1).
package sidx

// Index computes a positive or negative index wrt the length N.
//
// An invalid index always returns -1.
//
// This is only useful for strings and arrays.
// If you have a slice you should use [N].
func Index(N, i int) int {
	if N < 0 {
		return -1
	}
	// flip a negative index
	if i < 0 {
		i = N + i
	}
	// bounds check either case
	if i < 0 || i >= N {
		i = -1
	}
	return i
}

// N returns [Index] called with len(s).
func N[E any, S ~[]E](s S, i int) int {
	return Index(len(s), i)
}

// AtOk is like [At] except that it return false if i is an invalid index.
func AtOk[E any, S ~[]E](s S, i int) (E, bool) {
	var v E
	j := N(s, i)
	ok := j >= 0
	if ok {
		v = s[j]
	}
	return v, ok
}

// AtOr is like [At] except that it returns defaultValue if i is an invalid index.
func AtOr[E any, S ~[]E](s S, i int, defaultValue E) E {
	if j := N(s, i); j >= 0 {
		return s[j]
	}
	return defaultValue
}

// At returns s[[N](s, i)].
func At[E any, S ~[]E](s S, i int) E {
	return s[N(s, i)]
}

// End returns [At](s, -1).
func End[E any, S ~[]E](s S) E {
	return At(s, -1)
}

// Slice s using [N] for start and end.
func Slice[E any, S ~[]E](s S, start, end int) S {
	return s[N(s, start):N(s, end)]
}

// Pop returns the last element of s and a new slice without the last item.
// The last item is cleared before returning.
func Pop[E any, S ~[]E](s S) (E, S) {
	i := N(s, -1)
	last := s[i]
	clear(s[i:])
	return last, s[:i]
}
