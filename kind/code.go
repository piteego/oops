package kind

const (
	Unknown Code = 0
)

type Code int

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
