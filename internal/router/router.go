package router

// Router is a middleware used in EE to connect to the master
type Router interface {
	Pass(from, to Peer)
}

type Peer struct {
	Host string
	Port string
}
