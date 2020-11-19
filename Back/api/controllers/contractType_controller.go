package controllers

import (
	"net/http"
	"strconv"

	"github.com/abzibzi/jobOffers_API/api/models/enums"
	"github.com/abzibzi/jobOffers_API/api/responses"
	"github.com/gorilla/mux"
)

// GetConntractTypes return all contract types from DB
func (server *Server) GetConntractTypes(w http.ResponseWriter, r *http.Request) {
	contractType := enums.ContractType{}
	types, err := contractType.FindAllTypes(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, types)
}

// GetContractType returns Company data by his ID
func (server *Server) GetContractType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	contractType := enums.ContractType{}
	typeGotten, err := contractType.FindTypeByID(server.DB, int(id))
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, typeGotten)
}
