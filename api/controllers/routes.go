package controllers

import "github.com/peterwade153/ivents/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.Use(middlewares.SetContentTypeMiddleware)

	// User routes
	s.Router.HandleFunc("/api/users", s.GetUsers).Methods("GET")
}
