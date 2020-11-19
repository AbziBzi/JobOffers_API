package controllers

import (
	"net/http"
	"strconv"

	"github.com/abzibzi/jobOffers_API/api/models/enums"
	"github.com/abzibzi/jobOffers_API/api/responses"
	"github.com/gorilla/mux"
)

// GetRoles return all user roles from DB
func (server *Server) GetRoles(w http.ResponseWriter, r *http.Request) {
	role := enums.Role{}
	roles, err := role.FindAllRoles(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, roles)
}

// GetRole returns role by his ID
func (server *Server) GetRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	role := enums.Role{}
	roleGotten, err := role.FindRoleByID(server.DB, int(id))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, roleGotten)
}
