package v0

// Handler is a function type that takes an error and returns an *[Error].
type Handler func(error) *Error

// Handle processes an error using a list of handlers and returns an *[Error] as builtin error interface.
// It returns nil if the error is nil.
// It returns the original error if it is already an *[Error], or if no handlers return an non-nil *[Error].
func Handle(err error, handlers ...Handler) error {
	if err == nil {
		return nil
	}
	switch err.(type) {
	case *Error:
		return err

	default:
		for i := range handlers {
			if handlers[i] != nil {
				if oopsErr := handlers[i](err); oopsErr != nil {
					Because(err)(oopsErr)
					return oopsErr
				}
			}
		}
		return err
	}
}
