// Package oops provides a simple error handling mechanism that allows you to
// create, tag, and handle errors in a structured way.
//
// You can start using it defining your own error categories using [Label] errors.
// There are some examples in the [example] package that show an example of how to define custom error categories.
//
// Having a well-defined labels allows you to categorize your application errors using [New] function as follows:
//
//	err := oops.New("failed to process", oops.Tag(example.Duplicated.Error))
//
// [ErrorOption] is a function that modifies an [Error] instance, allowing you to set options like
// tagging the error with a [Label] or adding a stack trace with [CausedBy].
//
// This will allow you to handle errors in upper layers of your application either by [errors.Is] as follows:
//
//	if errors.Is(err, example.Duplicated.Error) {// You will be here!}
//
// or by using the [Map] type to handle builtin errors in a more structured way.
//
// Another option is to define more complicated [Handler](s) in different layers of your application and
// use the [Handle] function to process a given error calling a series of handlers.
package oops
