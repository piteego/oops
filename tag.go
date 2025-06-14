package oops

type Tag struct {
	Kind  error
	Cause error
}

func (*Tag) errorOption() {}

func (tag *Tag) newError(msg string, meta data) error {
	switch {
	case tag == nil && meta == nil:
		return &msgErr{msg: msg}

	case tag == nil && meta != nil:
		return &metaErr{msg: msg, meta: meta}

	default: // tag is not nil
		switch {
		case tag.Cause == nil && tag.Kind == nil && meta == nil: // basic error
			return &msgErr{msg: msg}

		case tag.Cause == nil && tag.Kind == nil && meta != nil: // meta error
			return &metaErr{msg: msg, meta: meta}

		case tag.Cause == nil && tag.Kind != nil && meta == nil: // kind error
			return &kindErr{msg: msg, kind: tag.Kind}

		case tag.Cause == nil && tag.Kind != nil && meta != nil: // kind & meta error
			return &kindMetaErr{msg: msg, kind: tag.Kind, meta: meta}

		case tag.Cause != nil && tag.Kind == nil && meta == nil: // cause error
			return &causeErr{msg: msg, cause: tag.Cause}

		case tag.Cause != nil && tag.Kind == nil && meta != nil: // cause & meta error
			return &causeMetaErr{msg: msg, cause: tag.Cause, meta: meta}

		case tag.Cause != nil && tag.Kind != nil && meta == nil: // standard error
			return &stdErr{msg: msg, standard: standard{kind: tag.Kind, cause: tag.Cause}}

		case tag.Cause != nil && tag.Kind != nil && meta != nil: // rich error
			return &richErr{msg: msg, standard: standard{kind: tag.Kind, cause: tag.Cause}, meta: meta}

		default:
			// This case should never happen, but we handle it gracefully.
			return &msgErr{msg: msg}
		}
	}
}
