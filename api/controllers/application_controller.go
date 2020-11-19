package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/abzibzi/jobOffers_API/api/auth"
	"github.com/abzibzi/jobOffers_API/api/models"
	"github.com/abzibzi/jobOffers_API/api/responses"
)

// CreateApplication creates Application
func (server *Server) CreateApplication(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	application := models.Application{}
	err = json.Unmarshal(body, &application)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fmt.Println(application)
	// application.Prepare()
	err = application.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized. Wrong token"))
		return
	}
	if userID != application.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized. Wrong User"))
		return
	}
	applicationCreated, err := application.SaveApplication(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d%d", r.Host, r.RequestURI, applicationCreated.UserID, applicationCreated.JobOfferID))
	responses.JSON(w, http.StatusCreated, applicationCreated)
}
