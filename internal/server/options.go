package server

func (srv *Server) WithInterceptors(interceptors ...Interceptor) *Server {
	srv.interceptorsChain = chainInterceptors(interceptors)
	return srv
}

func (srv *Server) WithHandler(handler Handler) *Server {
	srv.handler = handler
	return srv
}
