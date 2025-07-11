package oops

const (
	// UnknownKind represents an unspecified or invalid [Code].
	UnknownKind Kind = iota
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

type Kind int8

func (Kind) errorMetadata() {}

// Kind implements [Reader] interface which helps clients to read kind from errors
// using [Of] just by embedding [Code] into error metadata (See [example.MetaKindAndSeverity] or [oops.Standard])
func (k Kind) Kind(sanitize bool) Kind {
	if !sanitize {
		return k
	}
	return k.Sanitize()
}

// Sanitize returns a valid [Code] level or [Unknown] if the receiver is outside the valid range.
func (k Kind) Sanitize() Kind {
	if k < start || k > end {
		return UnknownKind
	}
	return k
}

type kindReader interface{ Kind(sanitize bool) Kind }

func KindOf(err error) Kind {
	if err == nil {
		return UnknownKind
	}
	if implemented, ok := err.(kindReader); ok {
		if level := implemented.Kind(true); level != UnknownKind {
			return level
		}
	}
	if implemented, ok := err.(interface{ Unwrap() error }); ok {
		return KindOf(implemented.Unwrap())
	}
	return UnknownKind
}
