package severity

type Reader interface{ Severity(sanitize bool) Level }

func Of(err error) Level {
	if err == nil {
		return Unknown
	}
	if implemented, ok := err.(Reader); ok {
		if level := implemented.Severity(true); level != Unknown {
			return level
		}
	}
	if implemented, ok := err.(interface{ Unwrap() error }); ok {
		return Of(implemented.Unwrap())
	}
	return Unknown
}
