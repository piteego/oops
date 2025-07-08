package severity

type Reader interface{ Severity() (Level, bool) }

func Of(err error) Level {
	if implemented, ok := err.(Reader); ok {
		if code, exists := implemented.Severity(); exists {
			return code
		}
	}
	if implemented, ok := err.(interface{ Unwrap() error }); ok {
		return Of(implemented.Unwrap())
	}
	return Unknown
}
