package oops

import (
	"errors"
	"fmt"
)

type option interface{ errorOption() }

func New(msg string, options ...option) error {
	var n struct{ internal, external int }
	// count internal and external options
	for i := range options {
		switch opt := options[i].(type) {
		case Tag:
			if opt.Label != nil {
				n.internal++
			}
		case Because:
			if opt.Error != nil {
				n.internal++
			}
		case Diagnosis:
			n.internal++

		default:
			n.external++
		}
	}
	// determine the type of error to return based on the options provided
	switch {
	case n.internal+n.external == 0: // errors.New...
		return errors.New(msg)

	case n.internal > 0 && n.external == 0: //  StandardError
		return newStandardError(msg, options...)

	case n.internal > 0 && n.external > 0: // RichError
		return newRichErr(msg, options...)

	case n.external > 0: // MetaError
		return newMetaError(msg, options...)

	default: // Unreachable code!
		return nil
	}
}

// MetaError is an error type that implements the standard [error] interface.
// It allows for attaching arbitrary, user-defined [Metadata] to an error,
// providing additional context relevant to the client's specific use case.
type MetaError struct {
	msg      string
	metadata option
}

// Error implements the error interface for [MetaError].
func (err *MetaError) Error() string { return err.msg }

// Unwrap returns the client's custom wrapped errors, to allow interoperability
// with [errors.Is], [errors.As].
func (err *MetaError) Unwrap() []error {
	if err.metadata == nil {
		return nil
	}
	if implemented, ok := err.metadata.(interface{ Unwrap() error }); ok {
		if implemented != nil {
			if unwrapped := implemented.Unwrap(); unwrapped != nil {
				errs := make([]error, 1)
				errs[0] = unwrapped
				return errs
			}
		}
		return nil
	}
	if implemented, ok := err.metadata.(interface{ Unwrap() []error }); ok {
		if implemented != nil {
			return implemented.Unwrap()
		}
		return nil
	}
	return nil
}

// StandardError is an error type that implements the standard [error] interface.
// It represents the 'oops' package's default structured error, providing core
// diagnostic details. These include a categorical [Label], the underlying [Because] error,
// and a specific [Diagnosis] that offers notes and a severity level.
type StandardError struct {
	msg       string
	label     Label
	cause     error
	diagnosis Diagnosis
}

// Error implements the error interface for [StandardError].
func (err *StandardError) Error() string { return err.msg }

// Unwrap returns the [Label], and [Because] errors to allow interoperability
// with [errors.Is], [errors.As].
func (err *StandardError) Unwrap() []error {
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

// RichError is an error type that implements the standard [error] interface.
// It serves as the most comprehensive error, combining the structured diagnostic
// capabilities of a [StandardError] with the flexibility to attach arbitrary,
// user-defined [Metadata] for additional context.
type RichError struct {
	StandardError
	metadata option
}

// Unwrap returns the [Label] error, [Because] error, and client's custom wrapped errors,
// to allow interoperability with [errors.Is], [errors.As]
func (err *RichError) Unwrap() []error {
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
	if err.metadata != nil {
		if implemented, ok := err.metadata.(interface{ Unwrap() []error }); ok {
			if unwrapped := implemented.Unwrap(); unwrapped != nil {
				errs = append(errs, unwrapped...)
			}
		}
		if implemented, ok := err.metadata.(interface{ Unwrap() error }); ok {
			if unwrapped := implemented.Unwrap(); unwrapped != nil {
				errs = append(errs, unwrapped)
			}
		}
	}
	return errs[:len(errs):len(errs)]
}

func newStandardError(msg string, options ...option) error {
	err := &StandardError{msg: msg}
	for i := range options {
		switch opt := options[i].(type) {
		case Tag:
			if err.label != nil {
				// if the error already has a label, do not overwrite it
				continue
			}
			err.label = opt.Label

		case Because:
			if opt.Error == nil {
				// skip nil errors in Because option
				continue
			}
			if err.cause != nil { // just append the new one to the existing cause
				err.cause = fmt.Errorf("%w: %v", err.cause, opt.Error)
			} else { // no cause yet, set it
				err.cause = opt.Error
			}

		case Diagnosis:
			err.diagnosis = opt

		default: // a client custom option or an unimplemented internal option!
			continue
		}
	}
	// If no label is set, use the default untagged label.
	if err.label == nil {
		err.label = Untagged
	}
	return err
}

func newRichErr(msg string, options ...option) error {
	err := new(RichError)
	err.msg = msg
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
			if err.metadata == nil && opt != nil {
				err.metadata = opt
			}
		}
	}
	// If no label is set, use the default Untagged label.
	if err.label == nil {
		err.label = Untagged
	}
	return err
}

func newMetaError(msg string, options ...option) error {
	err := &MetaError{msg: msg}
	for i := range options {
		switch opt := options[i].(type) {
		case Tag, Because, Diagnosis:
			// unreachable code, because we already checked for internal options
			continue

		default: // a client custom option
			if err.metadata == nil && opt != nil {
				// only the first non-nil custom option is stored
				err.metadata = opt
				break
			}
		}
	}
	return err
}
