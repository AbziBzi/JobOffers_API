package controllers

import "github.com/abzibzi/jobOffers_API/api/middlewares"

func (s *Server) initializeRoutes() {

	// User routes
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
}
