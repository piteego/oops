package kind

type Reader interface{ Kind(sanitize bool) Code }

func Of(err error) Code {
	if err == nil {
		return Unknown
	}
	if implemented, ok := err.(Reader); ok {
		if level := implemented.Kind(true); level != Unknown {
			return level
		}
	}
	if implemented, ok := err.(interface{ Unwrap() error }); ok {
		return Of(implemented.Unwrap())
	}
	return Unknown
}
