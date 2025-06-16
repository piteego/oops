package oops

type Standard struct {
	Kind error
	Diag diag
}

func (s Standard) Unwrap() error { return s.Kind }

func With(kind error, diag *diag) func(*Standard) {
	return func(s *Standard) {
		if kind != nil {
			s.Kind = kind
		}
		if diag != nil {
			s.Diag.cause = diag.cause
			s.Diag.severity = diag.severity
			s.Diag.note = diag.note
		}
	}
}

const (
	// Low severity level indicates a minor issue with low impact or urgency.
	Low level = iota + 1
	// Medium severity level indicates a moderate issue that needs attention but isn't critical.
	Medium
	// High severity level indicates a significant issue requiring immediate attention.
	High
	// Critical severity level indicates a severe, system-impacting issue requiring urgent resolution.
	Critical
)

// level is used as a severity level in the [Diagnosis] error option,
// indicating the importance or urgency of a particular error.
// It defines predefined levels ranging from [Low] to [Critical],
// with [undefinedSeverity] as the zero value.
type level uint8

// String returns the string representation of the severity level
func (l level) String() string {
	switch l {
	case Low: // 1
		return "Low"
	case Medium: // 2
		return "Medium"
	case High: // 3
		return "High"
	case Critical: // 4
		return "Critical"
	default:
		return "Unknown"
	}
}

// Diag creates a new Diag with the specified note and severity level.
func (l level) Diag(note string) *diag { return &diag{note: note, severity: l} }

// diag is a type which could be attached to a [Metadata] option for creating errors with the [New] function.
// It allows you to attach a detailed note and a specific severity level
// to an error, providing deeper insight into its nature and urgency.
// Create a Diag as follows:
//
// - [Low].Diag("note...") for minor issues,
//
// - [Medium].Diag("note...") for moderate issues,
//
// - [High].Diag("note...") for significant issues, or
//
// - [Critical].Diag("note...") for severe, urgent issues.
type diag struct {
	cause    error
	note     string // A detailed explanation or specific diagnostic message for the error.
	severity level  // The severity level of the error, indicating its importance or urgency.
}

func (d *diag) CausedBy(err error) *diag {
	d.cause = err
	return d
}

func (d *diag) Cause() error    { return d.cause }
func (d *diag) Note() string    { return d.note }
func (d *diag) Severity() level { return d.severity }
