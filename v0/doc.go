// Package v0 offers a straightforward and structured approach to error handling in Go applications.
//
// It enables you to create, categorize, and manage errors effectively using a system of labels and handlers.
//
// Key Features:
//
// - Categorize Errors with [Label]: Define custom error categories using [Label] error.
// This allows you to classify application errors consistently.
// Examples demonstrating how to define these custom categories can be found in the example package.
//
// - Create Labeled Errors: Use the [New] function to create new errors and associate them
// with your predefined labels. For instance:
//
// err := oops.New("failed to process", oops.Tag(example.Duplicated.Error))
//
// - Flexible Error Options:
//
// [ErrorOption] is a function that modifies an [Error] instance, allowing you to set options like
// tagging the error with a [Label] or adding a stack trace with [Because].
//
// - Stack Traces: Use [Because] in [New] function to append stack traces to your errors, providing valuable context for debugging.
//
// - Structured Error Handling:
//
// -- [errors.Is]: Handle errors in higher layers of your application using [errors.Is] to check against specific labeled errors:
//
//	if errors.Is(err, example.Duplicated.Error) { // You will be here! }
//
// -- [Map] Type: The [Map] type provides a structured way to handle builtin errors.
// Just define a map of errors to their corresponding *[Error] instances, and use the Map.Handle method to process errors.
// The [Handle] method will append the original error to the stack of the returned *[Error].
//
// -- Custom Handlers: Define complex [Handler] functions in different application layers.
// The [Handle] function can then be used to process a given error by invoking a series of these custom handlers.
package v0
