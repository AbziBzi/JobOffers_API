package controllers

import (
	"net/http"
	"strconv"

	"github.com/abzibzi/jobOffers_API/api/models/enums"
	"github.com/abzibzi/jobOffers_API/api/responses"
	"github.com/gorilla/mux"
)

// GetCompanyTypes return all company types from DB
func (server *Server) GetCompanyTypes(w http.ResponseWriter, r *http.Request) {
	companyType := enums.CompanyType{}

	types, err := companyType.FindAllTypes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, types)
}

// GetCompanyType returns Company data by his ID
func (server *Server) GetCompanyType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	companyType := enums.CompanyType{}
	typeGotten, err := companyType.FindTypeByID(server.DB, int(id))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, typeGotten)
}
