package models

type Server struct {
	DomainName string
	Port       int
}

func (server *Server) New(domainName string, port int) (*Server, error) {
	return &Server{
		DomainName: domainName,
		Port:       port,
	}, nil
}
