package v1

type (
	ErrorOption func() option
	option      interface{ oops() }
	//Option can be embedded in client structs to provide additional context or metadata for [Error]
	Option struct{}
)

func (Option) oops() {}

type label struct {
	Option
	label Label
}

func Tag(custom Label) ErrorOption {
	return func() option {
		if custom == nil {
			return nil
		}
		return label{label: custom}
	}
}

type cause struct {
	Option
	error
}

type causes struct {
	Option
	errs []error
}

func Because(errs ...error) ErrorOption {
	return func() option {
		if len(errs) == 0 {
			return nil
		}
		n := 0
		for i := range errs {
			if errs[i] != nil {
				n++
			}
		}
		if n == 0 {
			return nil
		}
		if n == 1 {
			return cause{error: errs[0]}
		}
		result := causes{
			errs: make([]error, 0, n),
		}
		for i := range errs {
			if errs[i] != nil {
				result.errs = append(result.errs, errs[i])
			}
		}
		return result
	}
}

type level uint8

const (
	Unknown level = iota // 0
	Low
	Medium
	High
	Critical
)

func (level) oops() {}

func Level(s level) ErrorOption {
	return func() option {
		if s < 0 {
			return nil
		}
		return s
	}
}

type diagnosis string

func (diagnosis) oops() {}

func Diagnosis(note string) ErrorOption {
	return func() option {
		if note == "" {
			return nil
		}
		return diagnosis(note)
	}
}
