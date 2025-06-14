package oops

// Metadata is an option that signals the intent to include custom, client-defined
// data when creating an error with the [New] function.
//
// When provided to [New], this option enables the resulting [MetaErr]
// or [RichErr] to store arbitrary additional context or information.
// The actual custom data itself is typically passed alongside this option
// during error creation (e.g., as another argument or via a dedicated setter).
//
// While this struct is empty and does not hold data directly, it acts as a
// type marker for this capability within the package's error construction.
type Metadata struct{}

func (Metadata) errorOption() {}

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

// Diag creates a new diag with the specified note and severity level.
func (l level) Diag(note string) diag { return diag{note: note, severity: l} }

// diag is a type which could be attached to [Tag] option for creating errors with the [New] function.
// It allows you to attach a detailed note and a specific severity level
// to an error, providing deeper insight into its nature and urgency.
// Provide a diag option to [New] using as follows:
//
// - [Low].Diag("note...") for minor issues,
//
// - [Medium].Diag("note...") for moderate issues,
//
// - [High].Diag("note...") for significant issues, or
//
// - [Critical].Diag("note...") for severe, urgent issues.
type diag struct {
	note     string // A detailed explanation or specific diagnostic message for the error.
	severity level  // The severity level of the error, indicating its importance or urgency.
}
