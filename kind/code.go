package kind

var _ Reader = Unknown

const Unknown Code = -1

type Code int8

func (c Code) Kind() Code { return c }

func Is(code, target Code, or ...Code) bool {
	if code == target {
		return true
	}
	for i := range or {
		if code == or[i] {
			return true
		}
	}
	return false
}
