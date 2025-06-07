package v1

import "errors"

// Untagged label serves as a default for errors created with the [New] function
// that don't have a specific [Label] assigned.
// It helps identify errors that haven't been categorized or labeled,
// ensuring a safe, non-nil [Label] value within the [CustomDiagnosticError] or [Summary] struct.
var Untagged Label = errors.New("untagged oops error label")

// Label is a type alias for builtin error, used to categorize application errors.
type Label error
