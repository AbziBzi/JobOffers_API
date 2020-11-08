package controllers

import "github.com/abzibzi/jobOffers_API/api/middlewares"

func (s *Server) initializeRoutes() {
	// Login route
	s.Router.HandleFunc("/api/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	// User routes
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteUser))).Methods("DELETE")

	// Company routes
	s.Router.HandleFunc("/api/companies/{id}", middlewares.SetMiddlewareJSON(s.GetCompany)).Methods("GET")
	s.Router.HandleFunc("/api/companies", middlewares.SetMiddlewareJSON(s.GetCompanies)).Methods("GET")
	s.Router.HandleFunc("/api/companies", middlewares.SetMiddlewareJSON(s.CreateCompany)).Methods("POST")
	s.Router.HandleFunc("/api/companies/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateCompany))).Methods("PUT")
	s.Router.HandleFunc("/api/companies/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteCompany))).Methods("DELETE")

	// Office routes
	s.Router.HandleFunc("/api/offices/{id}", middlewares.SetMiddlewareJSON(s.GetOffice)).Methods("GET")
	s.Router.HandleFunc("/api/offices", middlewares.SetMiddlewareJSON(s.GetOffices)).Methods("GET")
	s.Router.HandleFunc("/api/offices", middlewares.SetMiddlewareJSON(s.CreateOffice)).Methods("POST")
	s.Router.HandleFunc("/api/offices/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateOffice))).Methods("PUT")
	s.Router.HandleFunc("/api/offices/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteOffice))).Methods("DELETE")

	// Job offers routes
	s.Router.HandleFunc("/api/jobOffers/{id}", middlewares.SetMiddlewareJSON(s.GetJobOffert)).Methods("GET")
	s.Router.HandleFunc("/api/jobOffers", middlewares.SetMiddlewareJSON(s.GetJobOffers)).Methods("GET")
	s.Router.HandleFunc("/api/jobOffers", middlewares.SetMiddlewareJSON(s.CreateJobOffert)).Methods("POST")
	s.Router.HandleFunc("/api/jobOffers/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateJobOffert))).Methods("PUT")
	s.Router.HandleFunc("/api/jobOffers/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteJobOffert))).Methods("DELETE")

}
