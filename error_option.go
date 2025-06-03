package oops

type ErrorOption func(*Error)

func Tag(custom Label) ErrorOption {
	return func(err *Error) {
		if err.Label.Error != nil {
			// Do not overwrite existing category
			return
		}
		err.Label = custom
	}
}

func CausedBy(stack ...error) ErrorOption {
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
