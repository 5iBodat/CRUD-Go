package departement

import (
	"encoding/json"
	entitiesDepartement "golang-rest-api/entities/departement"
	response "golang-rest-api/helper"
	"golang-rest-api/model/departement"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Read(w http.ResponseWriter, r *http.Request) {
	type filter struct {
		Name string `json:"name"`
	}

	name := r.URL.Query().Get("name")
	result := departement.Read(name)
	if result == nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, "Failed read data")
		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, "Success Read Data", result)
	return
}
func Create(w http.ResponseWriter, r *http.Request) {
	var payload entitiesDepartement.Mst_departement
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, "Failed to decode data")
		return
	}

	result := departement.Create(payload)
	if result == (entitiesDepartement.Mst_departement{}) {
		response.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to create data")
		return
	}

	response.WriteSuccessResponse(w, http.StatusOK, "Success Create Data", result)
	return

}
func Update(w http.ResponseWriter, r *http.Request) {
	var payload entitiesDepartement.Mst_departement
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

	//get payload from body
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.WriteErrorResponse(w, http.StatusBadRequest, "Failed to decode data")
		return
	}

	result := departement.Update(id, payload)
	if result == (entitiesDepartement.Mst_departement{}) {
		response.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to update data")
	}
	response.WriteSuccessResponse(w, http.StatusOK, "Success Update Data", result)
	return
}
func Delete(w http.ResponseWriter, r *http.Request) {
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

	result := departement.Delete(id)

	response.WriteSuccessResponse(w, http.StatusOK, "Success Delete Data", result)
	return
}
