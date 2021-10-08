package server

import "net/http"

type ILogger interface {
	Error(msg string)
}

type IDataProvider interface {
	Data() interface{}
}

type Server struct {
	log      ILogger
	provider IDataProvider
}

func New(l ILogger, d IDataProvider) *Server {
	return &Server{log: l, provider: d}
}

func (s *Server) HandleIndex(w http.ResponseWriter, _ *http.Request) {
	s.newJSONWriter(w).Write(s.provider.Data())
}

func (s *Server) newJSONWriter(w http.ResponseWriter) jsonWriter {
	return jsonWriter{w: w, log: s.log}
}
