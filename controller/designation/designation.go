package designation

import (
	"encoding/json"
	entitiesDesignation "golang-rest-api/entities/designation"
	response "golang-rest-api/helper"
	"golang-rest-api/model/designation"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Read(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	result := designation.Read(name)
	if result == nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, "Failed read data")
		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, "Success Read Data", result)
	return
}

func Create(w http.ResponseWriter, r *http.Request) {
	payload := entitiesDesignation.Mst_designation{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, "Failed to decode data")
		return
	}

	result := designation.Create(payload)
	if result == (entitiesDesignation.Mst_designation{}) {
		response.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to create data")
		return
	}
	response.WriteSuccessResponse(w, http.StatusOK, "Success Create Data", result)
	return

}

func Update(w http.ResponseWriter, r *http.Request) {
	var payload entitiesDesignation.Mst_designation
	//get id from url
	idStr := mux.Vars(r)["id"]
	if idStr == "" {
		response.WriteErrorResponse(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, "id parameter is not valid")
		return
	}

	// Get the body of the request
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, "Failed to decode data")
		return
	}
	result := designation.Update(id, payload)
	if result == (entitiesDesignation.Mst_designation{}) {
		response.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to update data")
	}
	response.WriteSuccessResponse(w, http.StatusOK, "Success Update Data", result)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	//get id from url
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, "id parameter is not valid")
		return
	}
	result := designation.Delete(id)
	response.WriteSuccessResponse(w, http.StatusOK, "Success Delete Data", result)
	return
}
