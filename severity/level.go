package severity

const Unknown Level = 0

type Level int

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
