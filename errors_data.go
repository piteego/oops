package oops

// Diagnosis represents a detailed explanation of an error
type Diagnosis struct {
	Note     string // A detailed explanation or specific diagnostic message for the error.
	Severity level  // The severity level of the error, indicating its importance or urgency.
}

// Data returns the client's custom metadata associated with a [MetaError].
func (err *MetaError) Data() any {
	if err == nil {
		return nil
	}
	return err.metadata
}

func (err *StandardError) Label() Label {
	if err == nil {
		return nil
	}
	return err.label
}

// CausedBy returns the underlying cause of a [StandardError], if any.
func (err *StandardError) CausedBy() error {
	if err == nil {
		return nil
	}
	return err.cause
}

// Diag returns the diagnosis of the error, which includes a note and severity level.
func (err *StandardError) Diag() Diagnosis {
	if err == nil {
		return Diagnosis{}
	}
	return Diagnosis{
		Note:     err.diagnosis.note,
		Severity: err.diagnosis.severity,
	}
}

// Data returns the client's custom metadata associated with a [RichError].
func (err *RichError) Data() any {
	if err == nil {
		return nil
	}
	return err.metadata
}
