package oops

type Map map[error]*Error

func (m Map) Handle(err error) error {
	if oopsErr, exists := m[err]; exists {
		CausedBy(err)(oopsErr)
		return oopsErr
	}
	return err
}
