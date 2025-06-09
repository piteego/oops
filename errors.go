package oops

type option interface{ errorOption() }

func New(msg string, options ...option) error {
	var (
		tagOpt, metaOpt option
	)
	for i := range options {
		if options[i] == nil {
			continue // skip nil options
		}
		switch opt := options[i].(type) {
		case *Tag:
			if opt == nil || tagOpt != nil {
				continue // skip nil tags, and do not overwrite the first
			}
			tagOpt = opt

		case Metadata, *Metadata:
			continue // skip empty Metadata options

		default: // a client custom metadata option
			if opt == nil || metaOpt != nil {
				continue // skip nil custom metadata
			}
			metaOpt = opt // only the first non-nil custom metadata is stored
		}
	}
	switch {
	case tagOpt == nil && metaOpt == nil: // Simple Message Error
		return &msgErr{msg: msg}

	case tagOpt == nil && metaOpt != nil: // Meta Error
		return &metaErr{msg: msg, data: metaOpt}

	case tagOpt != nil && metaOpt == nil: // Standard Error TODO: break it down based on non-nil properties
		tag, _ := tagOpt.(*Tag)
		return &stdErr{msg, standard{label: tag.Label, cause: tag.Cause, diag: tag.Diag}}

	case tagOpt != nil && metaOpt != nil: // Rich Error
		tag, _ := tagOpt.(*Tag)
		return &richErr{msg, standard{label: tag.Label, cause: tag.Cause, diag: tag.Diag}, metaOpt}

	default:
		// This case should never happen, but we handle it gracefully.
		return &msgErr{msg: msg}
	}
}

type (
	msgErr  struct{ msg string }
	metaErr struct {
		msg  string
		data option
	}
	standard struct {
		label error
		cause error
		diag  diag
	}
	stdErr struct {
		msg string
		standard
	}
	richErr struct {
		msg      string
		standard standard
		data     option
	}
	wrapErr struct {
		msg     string
		wrapped wrapped
	}
	// diagErr is a struct that represents a diagnostic note with a severity level.
	diagErr struct {
		msg      string
		note     string
		severity level
	}
)

func (err *msgErr) Error() string  { return err.msg }
func (err *metaErr) Error() string { return err.msg }
func (err *stdErr) Error() string  { return err.msg }
func (err *richErr) Error() string { return err.msg }

func (err *wrapErr) Error() string { return err.msg }
func (err *diagErr) Error() string { return err.msg }

// Unwrap returns the wrapped error, allowing interoperability with [errors.Is], [errors.As], and [errors.Unwrap].
func (err *wrapErr) Unwrap() error { return err.wrapped.err }
