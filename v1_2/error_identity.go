package v1_2

func WithIdentity(kind error, code int) func(*identity) {
	return func(i *identity) {
		i.Code = code
		i.Kind = kind
	}
}

type identity struct {
	Code int
	Kind error
}

func (identity) errData()        {}
func (i identity) Unwrap() error { return i.Kind }
