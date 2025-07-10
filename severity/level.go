package severity

var _ Reader = Unknown

const Unknown Level = -1

type Level int

func (l Level) Severity() Level { return l }

func Is(level, target Level, or ...Level) bool {
	if level == target {
		return true
	}
	for i := range or {
		if level == or[i] {
			return true
		}
	}
	return false
}
