package v1

type (
	// Tag is an error option which adds a label to the error for general identification and categorization purposes.
	Tag struct{ Label Label }
	// Because is an error option which adds a cause to the error, indicating the reason for the error.
	// The non-nil provided error will be wrapped in the [CustomDiagnosticError] type and could be compared later using [errors.Is].
	Because struct{ Error error }
	// Diagnosis is an error option which adds a diagnosis not and severity level to the error,
	// providing additional context or information about the error.
	Diagnosis struct {
		Note     string
		Severity level
	}

	// Metadata is a generic error option that can be used to add custom data to the error.
	// Clients can embed this struct in their own error types to provide
	// additional context or information to the [CustomDiagnosticError].
	Metadata struct{}
)

func (Tag) errorOption()       {}
func (Because) errorOption()   {}
func (Diagnosis) errorOption() {}

func (Metadata) errorOption()    {}
func (Metadata) Unwrap() []error { return nil }

const (
	Low level = iota
	Medium
	High
	Critical
)

// level is used as a severity level in [Diagnosis] error option,
// indicating the importance or urgency of the error.
// Includes four predefined levels: [Low], [Medium], [High], and [Critical].
type level uint8

func (l level) String() string {
	switch l {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	case Critical:
		return "Critical"
	default:
		return "Unknown"
	}
}
