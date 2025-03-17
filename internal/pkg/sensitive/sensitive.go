package sensitive

// Sensitive data
type Sensitive string

func (s Sensitive) String() string {
	return "<REDACTED>"
}
