package v0

// Map is a type that maps errors to *[Error] instances.
type Map map[error]*Error

// Handle processes an error using the Map, returning the corresponding *[Error] if it exists.
// If the error is not found in the Map, it returns the original error.
// It also appends the original error to the stack of the returned *[Error] using [Because].
func (m Map) Handle(err error) error {
	if oopsErr, exists := m[err]; exists {
		Because(err)(oopsErr)
		return oopsErr
	}
	return err
}
