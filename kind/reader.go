package kind

type Reader interface{ Kind() (Code, bool) }

func Of(err error) Code {
	if implemented, ok := err.(Reader); ok {
		if code, exists := implemented.Kind(); exists {
			return code
		}
	}
	if implemented, ok := err.(interface{ Unwrap() error }); ok {
		return Of(implemented.Unwrap())
	}
	return Unknown
}
