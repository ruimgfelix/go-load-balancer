package models

type Server struct {
	DomainName string
	Port       int
	Weight     int
}

func (server *Server) New(domainName string, port int, weight int) (*Server, error) {
	return &Server{
		DomainName: domainName,
		Port:       port,
		Weight:     weight,
	}, nil
}
