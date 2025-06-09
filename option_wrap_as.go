package oops

const (
	AsLabel wrapAs = iota
	AsCause
)

type wrapAs int

// Wrap returns an option for [New] that wraps the error [AsLabel] or [AsCause].
func (as wrapAs) Wrap(err error) option {
	switch as {
	case AsLabel, AsCause:
		return wrapped{as: as, err: err}
	default:
		return nil
	}
}

type wrapped struct {
	as  wrapAs
	err error
}

func (w wrapped) errorOption() {}
