package network

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
	}
}
