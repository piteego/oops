package v1

import "fmt"

type metadata interface{ errorOption() }

func New(msg string, options ...metadata) error {
	var num struct{ internalOptions, externalOptions int }
	for i := range options {
		switch options[i].(type) {
		case Tag, Because, Diagnosis:
			num.internalOptions++
		default:
			num.externalOptions++
		}
	}
	switch {
	case num.internalOptions+num.externalOptions == 0: // no options provided
		return &basic{msg: msg}

	case num.internalOptions > 0: // at least one internal option is provided
		err := &CustomDiagnosticError{msg: msg}
		for i := range options {
			switch opt := options[i].(type) {
			case Tag:
				if err.label != nil {
					// if the error already has a label, do not overwrite it
					continue
				}
				err.label = opt.Label
			case Because:
				if err.cause != nil && opt.Error != nil {
					err.cause = fmt.Errorf("%w: %v", err.cause, opt.Error)
				}
				err.cause = opt.Error

			case Diagnosis:
				err.diagnosis = opt

			default: // a client custom option
				if err.custom == nil {
					err.custom = opt
				}
			}
		}
		return err

	case num.externalOptions > 0: // no internal options provided && at least one external option is provided
		err := &CustomError{msg: msg}
		for i := range options {
			switch opt := options[i].(type) {
			case Tag, Because, Diagnosis:
				// unreachable code, because we already checked for internal options
				continue

			default: // a client custom option
				if err.data == nil && opt != nil {
					// only the first non-nil custom option is stored
					err.data = opt
					break
				}
			}
		}
		return err

	default: // Unreachable code, because we already checked for internal and external options
		return nil
	}
}

// basic is a simple error type that implements the error interface.
// It is used as a base for other error types, such as [Labeled] and [Diagnostic].
// [New] will return a [basic] error if no options are provided.
type basic struct{ msg string }

// CustomDiagnosticError implements the error interface for a basic errors
func (b *basic) Error() string { return b.msg }

// CustomDiagnosticError is an error type that implements the builtin error interface.
// It can be used to create errors with additional context, such as a [Label], cause, [Diagnosis], and client's custom [Metadata].
type CustomDiagnosticError struct {
	msg       string
	label     Label
	cause     error
	diagnosis Diagnosis

	custom metadata
}

// Error implements the error interface for [CustomDiagnosticError].
func (err *CustomDiagnosticError) Error() string { return err.msg }

// Unwrap returns the [Label] error, [Because] error, and client's custom wrapped errors,
// to allow interoperability with [errors.Is], [errors.As]
func (err *CustomDiagnosticError) Unwrap() []error {
	n := 0
	if err.label != nil {
		n++
	}
	if err.cause != nil {
		n++
	}
	errs := make([]error, 0, n)
	if err.label != nil {
		errs = append(errs, err.label)
	}
	if err.cause != nil {
		errs = append(errs, err.cause)
	}
	return errs
}

// CustomError is an error type that implements the builtin error interface.
// It can be used to create errors with additional context, such as client's custom [Metadata].
type CustomError struct {
	msg  string
	data metadata
}

// Error implements the error interface for [CustomError].
func (c *CustomError) Error() string { return c.msg }

// Unwrap returns the client's custom wrapped errors, to allow interoperability
// with [errors.Is], [errors.As].
func (c *CustomError) Unwrap() []error {
	if c.data == nil {
		return nil
	}
	if implemented, ok := c.data.(interface{ Unwrap() error }); ok {
		if implemented != nil {
			errs := make([]error, 1)
			errs[0] = implemented.Unwrap()
			return errs
		}
		return nil
	}
	if implemented, ok := c.data.(interface{ Unwrap() []error }); ok {
		if implemented != nil {
			return implemented.Unwrap()
		}
		return nil
	}
	return nil
}
