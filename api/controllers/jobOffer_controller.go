package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/abzibzi/jobOffers_API/api/auth"
	"github.com/abzibzi/jobOffers_API/api/models"
	"github.com/abzibzi/jobOffers_API/api/responses"
	"github.com/gorilla/mux"
)

// GetJobOffert func gets one jobOffert by it's ID
func (server *Server) GetJobOffert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	jobOffert := models.JobOffer{}
	jobGotten, err := jobOffert.FindJobOffertByID(server.DB, int(id))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, jobGotten)
}

// GetJobOffers func gets all companies from DB
func (server *Server) GetJobOffers(w http.ResponseWriter, r *http.Request) {
	job := models.JobOffer{}
	jobs, err := job.FindAllJobOffers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, jobs)
}

// CreateJobOffert adds company tu DB
func (server *Server) CreateJobOffert(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	job := models.JobOffer{}
	err = json.Unmarshal(body, &job)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	job.Prepare()
	err = job.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized. Wrong token"))
		return
	}
	company := models.Company{}
	offertCompany, err := company.FindCompanyByID(server.DB, job.CompanyID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	if userID != offertCompany.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	jobCreated, err := job.SaveJobOffert(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, jobCreated.ID))
	responses.JSON(w, http.StatusCreated, jobCreated)
}

// UpdateJobOffert func updates existing job offert data
func (server *Server) UpdateJobOffert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	job := models.JobOffer{}
	jobGotten, err := job.FindJobOffertByID(server.DB, int(jobID))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Job offert not found"))
		return
	}
	company := models.Company{}
	offertCompany, err := company.FindCompanyByID(server.DB, jobGotten.CompanyID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	if userID != offertCompany.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	jobUpdate := models.JobOffer{}
	err = json.Unmarshal(body, &jobUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	newCompany := models.Company{}
	companyJobUpdated, err := newCompany.FindCompanyByID(server.DB, jobUpdate.CompanyID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	if userID != companyJobUpdated.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	jobUpdate.ID = jobGotten.ID
	jobUpdated, err := jobUpdate.UpdateJobOffert(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, jobUpdated)
}

// DeleteJobOffert removes job offert from DB
func (server *Server) DeleteJobOffert(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jobID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	job := models.JobOffer{}
	jobGotten, err := job.FindJobOffertByID(server.DB, int(jobID))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Job offert not found"))
		return
	}
	company := models.Company{}
	companyGotten, err := company.FindCompanyByID(server.DB, jobGotten.CompanyID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Job offert company not found"))
		return
	}
	if userID != companyGotten.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized. UserID does not match the company admin ID"))
		return
	}
	_, err = jobGotten.DeleteJobOffert(server.DB, int(jobID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", jobID))
	responses.JSON(w, http.StatusNoContent, "")
}
