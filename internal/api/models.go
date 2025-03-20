package api

// PointerValueWithDefault dereferences the given pointer if it is not null.
// Otherwise, it returns the default value of type this pointer refers to.
func PointerValueWithDefault[T any](p *T, def T) T {
	if p != nil {
		return *p
	}
	return def
}
