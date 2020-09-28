package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abzibzi/jobOfferts_API/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
)

// Server structure with references to the
// router and database used by the application.
type Server struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize connect to the database and wire up routes
func (server *Server) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("\nCannot connect to database %s:%s ", DbName, DbPort)
		log.Fatal("\nThis is the error:", err)
	} else {
		fmt.Printf("We are connected to the database %s ", DbName)
	}

	server.DB.Debug().AutoMigrate(&models.User{}) //database migration

	server.Router = mux.NewRouter().StrictSlash(true)
	server.initializeRoutes()
}

// RunServer runs server on given port
func (server *Server) RunServer(port string) {
	log.Printf("\nServer starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, server.Router))
}
