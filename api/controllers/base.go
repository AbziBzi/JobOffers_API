package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abzibzi/jobOfferts_API/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	server.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		fmt.Printf("\n Cannot connect to database %s", DbName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database %s", DbName)
	}

	server.DB.Debug().AutoMigrate(&models.User{}) //database migration

	server.Router = mux.NewRouter().StrictSlash(true)
	server.initializeRoutes()
}

// RunServer runs server on port :3030
func (server *Server) RunServer() {
	log.Printf("\nServer starting on port 3030")
	log.Fatal(http.ListenAndServe(":3030", server.Router))
}
