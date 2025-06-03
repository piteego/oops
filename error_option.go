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

func CausedBy(errs ...error) ErrorOption {
	return func(err *Error) {
		if len(errs) == 0 {
			return
		}
		for i := range errs {
			if errs[i] != nil {
				err.causes = append(err.causes, errs[i])
			}
		}
		if len(err.causes) < cap(err.causes) {
			// clip the slice
			err.causes = err.causes[:len(err.causes):len(err.causes)]
		}
	}
}
