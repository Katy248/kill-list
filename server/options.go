package server

func WithName(name string) OptionsFunc {
	return func(o Options) {
		o.Name = name
	}
}
