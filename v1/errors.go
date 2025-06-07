package v1

func New(msg string, options ...option) error {
	n := 0
	for i := range options {
		if options[i] != nil {
			n++
		}
	}
	if n == 0 {
		return &basic{msg: msg}
	}
	err := Error{msg: msg, data: make(map[string]option, n)}
	for i := range options {
		var key string
		switch options[i].(type) {
		case label:
			key = "_label"
		case cause:
			key = "_cause"
		case causes:
			key = "_causes"
		case Diagnosis:
			key = "_diagnosis"
		case level:
			key = "_level"
		default:
			key = "_custom"
		}

		if _, exist := err.data[key]; exist {
			// skip duplicate keys
			continue
		}
		if options[i] == nil {
			// skip options with nil value
			continue
		}
		err.data[key] = options[i]
	}
	return &err
}

// basic is a simple error type that implements the error interface.
// It is used as a base for other error types, such as [Labeled] and [Diagnostic].
// [New] will return a [basic] error if no options are provided.
type basic struct{ msg string }

// Error implements the error interface for a basic errors
func (b *basic) Error() string { return b.msg }

type Error struct {
	msg string
	//options  []option
	data map[string]option
}

func (e *Error) Error() string { return e.msg }

// Unwrap implements the error interface for [Error].
func (e *Error) Unwrap() []error {
	n := 0
	var (
		labeled, causedBy error
	)
	if opt, exist := e.data["_causes"]; exist && opt != nil {
		n += len(opt.(causes).errs)
	}

	if opt, exist := e.data["_label"]; exist && opt != nil {
		n++
		labeled = opt.(label).label
	}
	if opt, exist := e.data["_cause"]; exist && opt != nil {
		n++
		causedBy = opt.(cause).error
	}
	if n == 0 {
		return nil
	}
	errs := make([]error, 0, n)
	if opt, exist := e.data["_causes"]; exist && opt != nil {
		errs = append(errs, opt.(causes).errs...)
	}
	if causedBy != nil {
		errs = append(errs, causedBy)
	}
	if labeled != nil {
		errs = append(errs, labeled)
	}
	return errs
}
