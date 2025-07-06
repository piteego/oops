package v1_2

const (
	// SeverityLow level indicates a minor issue with low impact or urgency.
	SeverityLow level = iota + 1
	// SeverityMedium level indicates a moderate issue that needs attention but isn't critical.
	SeverityMedium
	// SeverityHigh level indicates a significant issue requiring immediate attention.
	SeverityHigh
	// SeverityCritical level indicates a severe, system-impacting issue requiring urgent resolution.
	SeverityCritical
)

// level is used as a severity level in the [Diagnosis] error option,
// indicating the importance or urgency of a particular error.
// It defines predefined levels ranging from [Low] to [Critical],
// with [undefinedSeverity] as the zero value.
type level uint8

// String returns the string representation of the severity level
func (l level) String() string {
	switch l {
	case SeverityLow: // 1
		return "Low"
	case SeverityMedium: // 2
		return "Medium"
	case SeverityHigh: // 3
		return "High"
	case SeverityCritical: // 4
		return "Critical"
	default:
		return "Unknown"
	}
}

// Diag creates a new Diag with the specified note and severity level.
func (l level) Diag(causedBy error, note string) func(*diagnosis) {
	return func(d *diagnosis) {
		d.Cause = causedBy
		d.Note = note
		d.Severity = l
	}
}

// diagnosis is a type which could be attached to a [Metadata] option for creating errors with the [New] function.
// It allows you to attach a detailed note and a specific severity level
// to an error, providing deeper insight into its nature and urgency.
// Create a Diag as follows:
//
// - [SeverityLow].Diag(causeErr, "note...") for minor issues,
//
// - [SeverityMedium].Diag(causeErr, "note...") for moderate issues,
//
// - [SeverityHigh].Diag(causeErr, "note...") for significant issues, or
//
// - [SeverityCritical].Diag(causeErr, "note...") for severe, urgent issues.
type diagnosis struct {
	Cause    error
	Note     string // A detailed explanation or specific diagnosis message for the error.
	Severity level  // The severity level of the error, indicating its importance or urgency.
}

func (diagnosis) errData()        {}
func (d diagnosis) Unwrap() error { return d.Cause }
