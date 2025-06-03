package oops

type Handler func(error) *Error

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
					CausedBy(err)(oopsErr)
					return oopsErr
				}
			}
		}
		return err
	}
}
