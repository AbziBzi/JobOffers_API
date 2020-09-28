package controllers

import (
	"net/http"

	"github.com/abzibzi/jobOfferts_API/api/models"
	"github.com/abzibzi/jobOfferts_API/api/responses"
)

// GetUsers return all Users from DB
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}
