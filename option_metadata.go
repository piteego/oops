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
