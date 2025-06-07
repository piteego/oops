package v1

type option interface{ mustBeImplementedOrEmbedErrorOptionStruct() }

type Option struct{}

func (Option) mustBeImplementedOrEmbedErrorOptionStruct() {}

type label struct {
	Option
	label Label
}

func Tag(custom Label) option {
	if custom == nil {
		return nil
	}
	return label{label: custom}
}

type cause struct {
	Option
	error error
}

type causes struct {
	Option
	errs []error
}

func Because(errs ...error) option {
	if len(errs) == 0 {
		return Option{}
	}
	n := 0
	for i := range errs {
		if errs[i] != nil {
			n++
		}
	}
	if n == 0 {
		return Option{}
	}
	if n == 1 {
		for i := range errs {
			if errs[i] != nil {
				return cause{error: errs[0]}
			}
		}
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

const (
	Low level = iota
	Medium
	High
	Critical
)

type level uint8

func (l level) mustBeImplementedOrEmbedErrorOptionStruct() {}

func (l level) String() string {
	switch l {
	case Low:
		return "Low"
	case Medium:
		return "Medium"
	case High:
		return "High"
	case Critical:
		return "Critical"
	default:
		return "Unknown"
	}
}

func (l level) Level() option {
	if l == 0 {
		return Option{}
	}
	return l
}

type Diagnosis string

func (Diagnosis) mustBeImplementedOrEmbedErrorOptionStruct() {}
