package oops

const (
	UnknownCode     Code  = 0
	UnknownSeverity Level = 0
)

type (
	Code  int
	Level int
	Cause error
)

func (c Code) Is(target Code, or ...Code) bool {
	if c == target {
		return true
	}
	for i := range or {
		if c == or[i] {
			return true
		}
	}
	return false
}

func (l Level) Is(target Level, or ...Level) bool {
	if l == target {
		return true
	}
	for i := range or {
		if l == or[i] {
			return true
		}
	}
	return false
}
