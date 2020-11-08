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

// GetOffice func gets one office by it's ID
func (server *Server) GetOffice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	office := models.Office{}
	officeGotten, err := office.FindOfficeByID(server.DB, int(id))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, officeGotten)
}

// GetOffices func gets all offices
func (server *Server) GetOffices(w http.ResponseWriter, r *http.Request) {
	office := models.Office{}
	offices, err := office.FindAllOffices(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, offices)
}

// CreateOffice adds office to DB
func (server *Server) CreateOffice(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	office := models.Office{}
	err = json.Unmarshal(body, &office)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	office.Prepare()
	err = office.Validate()
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
	officeCompany, err := company.FindCompanyByID(server.DB, office.CompanyID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	if userID != officeCompany.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	officeCreated, err := office.SaveOffice(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, officeCreated.ID))
	responses.JSON(w, http.StatusCreated, officeCreated)
}

// UpdateOffice func updates existing office data
func (server *Server) UpdateOffice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	officeID, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	office := models.Office{}
	officeGotten, err := office.FindOfficeByID(server.DB, int(officeID))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Office not found"))
		return
	}
	company := models.Company{}
	officeCompany, err := company.FindCompanyByID(server.DB, officeGotten.CompanyID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	if userID != officeCompany.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	officeUpdate := models.Office{}
	err = json.Unmarshal(body, &officeUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	officeCompany, err = company.FindCompanyByID(server.DB, officeUpdate.CompanyID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	if userID != officeCompany.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	officeUpdate.ID = officeGotten.ID
	officeUpdated, err := officeUpdate.UpdateOffice(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, officeUpdated)
}

// DeleteOffice removes office from DB
func (server *Server) DeleteOffice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	officeID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	office := models.Office{}
	officeGotten, err := office.FindOfficeByID(server.DB, int(officeID))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Office not found"))
		return
	}
	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	company := models.Company{}
	officeCompany, err := company.FindCompanyByID(server.DB, office.CompanyID)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	if userID != officeCompany.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = officeGotten.DeleteOffice(server.DB, int(officeID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", officeID))
	responses.JSON(w, http.StatusNoContent, "")
}
