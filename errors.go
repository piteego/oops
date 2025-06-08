package oops

import (
	"errors"
	"fmt"
)

type (
	option interface{ errorOption() }
)

func New(msg string, options ...option) error {
	// count valid standard and metadata options
	standard, metadata := countValidOptions(options...)
	// determine the type of error to return based on the options provided
	switch {
	case standard == 0 && metadata == 0:
		return errors.New(msg)

	case standard == 0 && metadata > 0: // MetaError
		err := &MetaError{msg: msg}
		for i := range options {
			switch opt := options[i].(type) {
			case Tag, Because, Diagnosis:
				// Skip possible invalid standard options
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

	case standard > 0 && metadata == 0: // StandardError
		err := &StandardError{msg: msg, label: Untagged}
		for i := range options {
			applyStandardOption(err, options[i])
		}
		return err

	case standard > 0 && metadata > 0: // RichError
		err := &RichError{StandardError: StandardError{msg: msg, label: Untagged}}
		for i := range options {
			switch opt := options[i].(type) {
			case Tag, Because, Diagnosis:
				applyStandardOption(&err.StandardError, opt)

			default: // a client custom option
				if err.metadata == nil && opt != nil { // only the first non-nil custom option is stored
					err.metadata = opt
				}
			}
		}
		return err

	default: // Unreachable code!
		return nil
	}
}

func countValidOptions(options ...option) (standard, metadata int) {
	if len(options) == 0 {
		return 0, 0 // no options provided
	}
	for i := range options {
		switch opt := options[i].(type) {
		case Tag:
			if opt.Label == nil || opt.Label == Untagged {
				continue // skip nil and Untagged labels
			}
			standard++

		case Because:
			if opt.Error == nil {
				continue // skip nil cause error
			}
			standard++

		case Diagnosis: // is always valid (see Diag method )
			standard++

		default:
			if opt == nil {
				continue // skip nil options
			}
			metadata++
		}
	}
	return
}

func applyStandardOption(err *StandardError, input option) {
	if input == nil { // skip nil options
		return
	}
	switch opt := input.(type) {
	case Tag:
		if opt.Label == nil || opt.Label == Untagged {
			// skip nil and Untagged labels
			return
		}
		if err.label == nil || err.label == Untagged {
			// if no label is set, or the error already labeled as Untagged, set the new tagged label
			err.label = opt.Label
		}

	case Because:
		if opt.Error == nil { // skip nil cause error in Because option
			return
		}
		if err.cause != nil { // just append the new one to the existing cause
			err.cause = fmt.Errorf("%w: %v", err.cause, opt.Error)
		} else { // no cause yet, set it
			err.cause = opt.Error
		}

	case Diagnosis:
		err.diagnosis = opt

	default:
		// an unimplemented internal option!
		// or an invalid metadata option.
		return
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

// Unwrap returns the client's custom wrapped errors, to allow interoperability with [errors.Is], [errors.As].
func (err *MetaError) Unwrap() []error {
	if err.metadata == nil {
		return nil
	}
	switch obj := err.metadata.(type) {
	case nil: // skip nil metadata object
		return nil

	case interface{ Unwrap() error }:
		wErr := obj.Unwrap()
		if wErr == nil {
			return nil
		}
		return []error{wErr}

	case interface{ Unwrap() []error }:
		return obj.Unwrap()

	default:
		return nil
	}
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

// Unwrap returns the [Label], and [Because] errors to allow interoperability with [errors.Is], [errors.As].
func (err *StandardError) Unwrap() []error {
	n := 1 // always at least one error (the label)
	if err.cause != nil {
		n++
	}
	errs := make([]error, 0, n)
	if err.cause != nil {
		errs = append(errs, err.cause)
	}
	if err.label != nil {
		errs = append(errs, err.label)
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

// Unwrap returns the [Label] error, [Because] error, and client's custom wrapped errors
// to allow interoperability with [errors.Is], [errors.As]
func (err *RichError) Unwrap() []error {
	errs := err.StandardError.Unwrap() //  an allocated slice of non-nil errors (the label and cause)
	if err.metadata == nil {
		return errs
	}
	switch obj := err.metadata.(type) {
	case nil:
	case interface{ Unwrap() error }:
		if wErr := obj.Unwrap(); wErr != nil {
			errs = append(errs, wErr)
		}

	case interface{ Unwrap() []error }:
		if wErrs := obj.Unwrap(); wErrs != nil {
			for i := range wErrs {
				if wErrs[i] == nil {
					continue // skip nil errors
				}
				errs = append(errs, wErrs[i])
			}
		}
	}
	return errs[:len(errs):len(errs)] // clip the slice to the actual length
}
