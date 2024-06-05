package server

type Options struct {
	Name string
}

func DefaultOptions() Options {
	return Options{
		Name: "default",
	}
}

type Server struct {
	Options
}

type OptionsFunc func(Options)

func NewServer(opts ...OptionsFunc) Server {
	options := DefaultOptions()
	for _, fn := range opts {
		fn(options)
	}
	return Server{Options: options}
}
