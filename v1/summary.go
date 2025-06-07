package v1

type Summary struct {
	Message   string    `json:"message,omitempty"`
	Label     Label     `json:"label,omitempty"`
	Cause     error     `json:"cause,omitempty"`
	Diagnosis Diagnosis `json:"diagnosis,omitempty"`

	Metadata any `json:"metadata,omitempty"`
}

func Describe(target error) *Summary {
	if target == nil {
		return nil
	}

	result := &Summary{
		Label:   Untagged,
		Message: target.Error(),
	}
	switch err := target.(type) {
	case *CustomDiagnosticError:
		if err.label != nil {
			result.Label = err.label
		}
		if err.cause != nil {
			result.Cause = err.cause
		}
		result.Diagnosis = err.diagnosis
		if err.custom != nil {
			result.Metadata = err.custom
		}
		return result

	case *CustomError:
		if err.data != nil {
			result.Metadata = err.data
		}
		return result

	default:
		return result
	}
}
