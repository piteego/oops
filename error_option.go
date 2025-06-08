package oops

type (
	// Tag is an option for creating errors with the [New] function.
	// It embeds a [Label] to categorize the error, aiding in general identification
	// and classification.
	//
	// When a Tag option is provided to [oops.New], the resulting error will be
	// either a [StandardError] or a [RichError], depending on other options used
	// (e.g., if custom metadata is also provided).
	//
	// If [oops.New] is called a Tag option, the label will be wrapped within the returned
	// [StandardError] or [RichError].
	// This allows for inspecting the error chain and comparison using [errors.Is].
	Tag struct{ Label Label }

	// Because is an option for creating errors with the [New] function.
	// It allows you to specify an underlying cause for the error, indicating
	// the reason for its occurrence.
	//
	// When a non-nil error is provided via the Because option to [New],
	// that original error will be wrapped within the returned [StandardError] or [RichError].
	// This wrapping enables inspecting the error chain and comparison using [errors.Is].
	Because struct{ Error error }

	// Diagnosis is an option for creating errors with the [New] function.
	// It allows you to attach a detailed note and a specific Severity level
	// to an error, providing deeper insight into its nature and urgency.
	Diagnosis struct {
		Note     string // A detailed explanation or specific diagnostic message for the error.
		Severity level  // The severity level of the error, indicating its importance or urgency.
	}

	// Metadata is an option that signals the intent to include custom, client-defined
	// data when creating an error with the [New] function.
	//
	// When provided to [New], this option enables the resulting [MetaError]
	// or [RichError] to store arbitrary additional context or information.
	// The actual custom data itself is typically passed alongside this option
	// during error creation (e.g., as another argument or via a dedicated setter).
	//
	// While this struct is empty and does not hold data directly, it acts as a
	// type marker for this capability within the package's error construction.
	Metadata struct{}
)

func (Tag) errorOption()       {}
func (Because) errorOption()   {}
func (Diagnosis) errorOption() {}

func (Metadata) errorOption() {}

const (
	// UndefinedSeverity indicates an uninitialized or unknown severity level.
	// This is the zero value for the level type.
	UndefinedSeverity level = iota
	// Low indicates a minor issue with low impact or urgency.
	Low
	// Medium indicates a moderate issue that needs attention but isn't critical.
	Medium
	// High indicates a significant issue requiring immediate attention.
	High
	// Critical indicates a severe, system-impacting issue requiring urgent resolution.
	Critical
)

// level is used as a severity level in the [Diagnosis] error option,
// indicating the importance or urgency of a particular error.
// It defines predefined levels ranging from [Low] to [Critical],
// with [UndefinedSeverity] as the zero value.
type level uint8

// String returns the string representation of the severity level.
func (l level) String() string {
	switch l {
	case UndefinedSeverity: // 0
		return "Undefined"
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
