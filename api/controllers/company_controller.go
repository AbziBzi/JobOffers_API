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

// GetCompany func gets one company by it's ID
func (server *Server) GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	company := models.Company{}
	companyGotten, err := company.FindCompanyByID(server.DB, int(id))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, companyGotten)
}

// GetCompanies func gets all companies from DB
func (server *Server) GetCompanies(w http.ResponseWriter, r *http.Request) {
	company := models.Company{}

	companies, err := company.FindAllCompanies(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, companies)
}

// CreateCompany adds company tu DB
func (server *Server) CreateCompany(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	company := models.Company{}
	err = json.Unmarshal(body, &company)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	company.Prepare()
	err = company.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized. Wrong token"))
		return
	}
	if userID != company.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	user := models.User{}
	admin, err := user.FindUserByID(server.DB, userID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	if admin.RoleID != 2 {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized. User is a developer"))
		return
	}
	companyCreated, err := company.SaveCompany(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, companyCreated.ID))
	responses.JSON(w, http.StatusCreated, companyCreated)
}

// UpdateCompany func updates existing company data
func (server *Server) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id, err := strconv.ParseUint(vars["id"], 10, 32)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusBadRequest, err)
	// 	return
	// }
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	company := models.Company{}
	err = json.Unmarshal(body, &company)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	company.Prepare()
	err = company.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedCompany, err := company.UpdateCompany(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, updatedCompany)
}
