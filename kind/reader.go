package kind

type Reader interface{ Kind() Code }

func Of(err error) Code {
	if implemented, ok := err.(Reader); ok {
		if implemented != nil {
			return implemented.Kind()
		}
	}
	if implemented, ok := err.(interface{ Unwrap() error }); ok {
		return Of(implemented.Unwrap())
	}
	return Unknown
}
