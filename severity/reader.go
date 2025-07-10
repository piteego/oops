package severity

type Reader interface{ Severity() Level }

func Of(err error) Level {
	if implemented, ok := err.(Reader); ok {
		if implemented != nil {
			return implemented.Severity()
		}
	}
	if implemented, ok := err.(interface{ Unwrap() error }); ok {
		return Of(implemented.Unwrap())
	}
	return Unknown
}
