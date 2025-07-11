package oops

const (
	// UnknownSeverityLevel represents an unspecified or invalid severity [Level].
	UnknownSeverityLevel Level = iota
	// Critical (SEV1, Severity 1):
	//
	// represents a major incident with very high impact, potentially causing significant business disruption or harm
	Critical
	// High (SEV2, Severity 2):
	//
	// A major incident with significant impact, potentially affecting a large number of users or key functionalities
	High
	// Medium (SEV3, Severity 3):
	//
	// A moderate incident with partial loss of functionality or impact on a small subset of users
	Medium
	// Low (SEV4, Severity 4):
	//
	//A minor incident with little to no impact on functionality or users
	Low
	// Informational (SEV5, Severity 5):
	//
	//An event that is logged for informational purposes, often with no immediate
	Informational

	levelStart, levelEnd = Critical, Informational
)

// Level of severity help determine the extent of damage or disruption caused by an incident or vulnerability.
// They guide teams in prioritizing which issues to address first, ensuring critical problems are handled promptly.
// They also provide a common language for communicating the severity of an incident.
//
// Valid values:
//
// - [Critical]
//
// - [High]
//
// - [Medium]
//
// - [Low]
//
// - [Informational]
//
// An invalid Severity will be represented as [Unknown].
type Level uint8

func (Level) errorMetadata() {}

func (l Level) Severity(sanitize bool) Level {
	if !sanitize {
		return l
	}
	return l.Sanitize()
}

// Sanitize returns a valid [Level] level or unknown severity level if the receiver is outside the valid range.
func (l Level) Sanitize() Level {
	if l < levelStart || l > levelEnd {
		return UnknownSeverityLevel
	}
	return l
}

// Valid reports whether the [Level] level is within the defined range.
// Returns true for values between [Critical] and [Informational] (inclusive).
func (l Level) Valid() bool {
	if l < levelStart || l > levelEnd {
		return false
	}
	return true
}

type levelReader interface{ Severity(sanitize bool) Level }

func LevelOf(err error) Level {
	if err == nil {
		return UnknownSeverityLevel
	}
	if implemented, ok := err.(levelReader); ok {
		if level := implemented.Severity(true); level != UnknownSeverityLevel {
			return level
		}
	}
	if implemented, ok := err.(interface{ Unwrap() error }); ok {
		return LevelOf(implemented.Unwrap())
	}
	return UnknownSeverityLevel
}
