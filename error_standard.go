package oops

func Standard(kind Kind, severity Level, cause error) standard {
	return standard{
		kind: kind, level: severity, causedBy: cause,
	}
}

type standard struct {
	kind     Kind
	level    Level
	causedBy error
}

func (s standard) errorMetadata()               {}
func (s standard) Cause() error                 { return s.causedBy }
func (s standard) Severity(sanitize bool) Level { return s.level }
func (s standard) Kind(sanitize bool) Kind      { return s.kind }
