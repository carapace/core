package core

// Core is the godlevel construct, which fulfilles the interface defined in api/proto/v1/services.proto
type Core struct {
	Conf *Config
}

// New is the constructor for Core
func New(conf *Config, opts ...Option) *Core {
	c := &Core{
		Conf: conf,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// An Option is a function which sets Core configurations
type Option func(core *Core)

// Config holds the dependencies of Core
type Config struct {
	In ConfigIngress
}
