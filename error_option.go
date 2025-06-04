package oops

// ErrorOption is a function that modifies an [Error] instance.
// It is used to set options like Tagging the error with a [Label] or
// adding a stack trace with [Because], etc.
type ErrorOption func(*Error)

// Tag sets a custom [Label] for the *[Error]. If the error already has a non-nil [Label], it will not overwrite it.
func Tag(custom Label) ErrorOption {
	return func(err *Error) {
		if err.Label != nil {
			// Do not overwrite existing label
			return
		}
		err.Label = custom
	}
}

// Because append stack errors to the *[Error] stack.
func Because(stack ...error) ErrorOption {
	return func(err *Error) {
		if len(stack) == 0 {
			return
		}
		// TODO: improve allocation by adding new capacity to err.stack based on the number of non-nil stack errors.
		for i := range stack {
			if stack[i] != nil {
				err.stack = append(err.stack, stack[i])
			}
		}
		if len(err.stack) < cap(err.stack) {
			// clip the slice
			err.stack = err.stack[:len(err.stack):len(err.stack)]
		}
	}
}
