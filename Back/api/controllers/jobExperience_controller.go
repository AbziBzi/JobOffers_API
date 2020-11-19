package controllers

import (
	"net/http"
	"strconv"

	"github.com/abzibzi/jobOffers_API/api/models/enums"
	"github.com/abzibzi/jobOffers_API/api/responses"
	"github.com/gorilla/mux"
)

// GetJobExperiences return all job experiences from DB
func (server *Server) GetJobExperiences(w http.ResponseWriter, r *http.Request) {
	experience := enums.JobExperience{}
	experiences, err := experience.FindAllExperiences(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, experiences)
}

// GetJobExperience returns Company data by his ID
func (server *Server) GetJobExperience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	experience := enums.JobExperience{}
	experienceGotten, err := experience.FindExperienceByID(server.DB, int(id))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, experienceGotten)
}
