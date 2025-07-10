package kind

const (
	// Unknown represents an unspecified or invalid [Code].
	Unknown Code = iota - 1
	// Internal represents and internal server error
	Internal
	// Unprocessable represents
	Unprocessable
	Validation
	Unauthorized
	Forbidden
	NotFound
	Duplication

	start, end = Internal, Duplication
)

type Code int8

// Kind implements [Reader] interface which helps clients to read kind from errors
// using [Of] just by embedding [Code] into error metadata (See [example.MetaKindAndSeverity] or [oops.Standard])
func (c Code) Kind(sanitize bool) Code {
	if !sanitize {
		return c
	}
	return c.Sanitize()
}

// Sanitize returns a valid [Code] level or [Unknown] if the receiver is outside the valid range.
func (c Code) Sanitize() Code {
	if c < start || c > end {
		return Unknown
	}
	return c
}
