package oops

type option interface{ errorOption() }

func New(msg string, options ...option) error {
	var (
		tag  *Tag
		meta option
	)
	for i := range options {
		if options[i] == nil {
			continue // skip nil options
		}
		switch opt := options[i].(type) {
		case *Tag:
			if opt == nil || tag != nil {
				continue // skip nil tags, and do not overwrite the first
			}
			tag = opt

		case Metadata, *Metadata:
			continue // skip empty Metadata options

		default: // a client custom metadata struct
			if opt == nil || meta != nil {
				continue // skip nil custom metadata
			}
			meta = opt // only the first non-nil custom metadata is stored
		}
	}
	return tag.newError(msg, meta)
}

type (
	msgErr  struct{ msg string }
	data    = option
	metaErr struct {
		msg  string
		meta data
	}
	kindErr struct {
		msg  string
		kind error
	}
	kindMetaErr struct {
		msg  string
		kind error
		meta data
	}
	causeErr struct {
		msg   string
		cause error
	}
	causeMetaErr struct {
		msg   string
		cause error
		meta  data
	}
	standard struct {
		kind  error
		cause error
	}
	stdErr struct {
		msg string
		standard
	}
	richErr struct {
		msg      string
		standard standard
		meta     data
	}
)

func (err *msgErr) Error() string       { return err.msg }
func (err *metaErr) Error() string      { return err.msg }
func (err *kindErr) Error() string      { return err.msg }
func (err *kindMetaErr) Error() string  { return err.msg }
func (err *causeErr) Error() string     { return err.msg }
func (err *causeMetaErr) Error() string { return err.msg }
func (err *stdErr) Error() string       { return err.msg }
func (err *richErr) Error() string      { return err.msg }

// Unwrap returns the wrapped kind error, allowing interoperability
// of a kind error with [errors.Is], [errors.As], and [errors.Unwrap].
func (err *kindErr) Unwrap() error { return err.kind }

// Unwrap returns the wrapped kind error, allowing interoperability
// of a kind-meta error with [errors.Is], [errors.As], and [errors.Unwrap].
func (err *kindMetaErr) Unwrap() error { return err.kind }

// Unwrap returns the wrapped kind error, allowing interoperability
// of a standard error with [errors.Is], [errors.As], and [errors.Unwrap].
func (err *stdErr) Unwrap() error { return err.kind }

// Unwrap returns the wrapped kind error, allowing interoperability
// of a rich error with [errors.Is], [errors.As], and [errors.Unwrap].
func (err *richErr) Unwrap() error { return err.standard.kind }
