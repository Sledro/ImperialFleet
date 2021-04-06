package api

import (
	"ImperialFleet/core"
	"ImperialFleet/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// spaceshipCreateHandler - Create a spaceship
func spaceshipCreateHandler(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create spaceship object
	spaceship := core.Spaceship{}
	err := json.NewDecoder(r.Body).Decode(&spaceship)
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipCreateHandler", err, w)
		return
	}

	// Validate input
	validate := validator.New()
	err = validate.Struct(spaceship)
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipCreateHandler", err, w)
		return
	}

	// Store
	err = spaceship.CreateSpaceship()
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipCreateHandler", err, w)
		return
	}

	// Sort the data to be returned
	returnData := map[string]interface{}{
		"success": true,
	}
	json.NewEncoder(w).Encode(returnData)
}

// spaceshipGetHandler - Get a spaceship
func spaceshipGetHandler(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get url param
	vars := mux.Vars(r)

	//validate param
	v := validator.New()
	err := v.Var(vars["id"], "required,max=999999,numeric")
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipGetHandler", errors.New("Must be numeric and less than 999999"), w)
		return
	}

	// convert from string to int
	spaceshipID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipGetHandler", err, w)
		return
	}

	// Get spaceship
	spaceship, err := core.GetSpaceship(spaceshipID)
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipGetHandler", err, w)
		return
	}

	// Sort the data to be returned
	json.NewEncoder(w).Encode(spaceship)
}

// spaceshipListHandler - List spaceships based on (name, class, status)
func spaceshipListHandler(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create spaceship object
	spaceship := core.Spaceship{}
	err := json.NewDecoder(r.Body).Decode(&spaceship)
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipListHandler", err, w)
		return
	}

	// Validate input
	validate := validator.New()
	err = validate.Struct(spaceship)
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipListHandler", err, w)
		return
	}

	// Get List
	data, err := core.ListSpaceships(spaceship.Name, spaceship.Class, spaceship.Status)
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipListHandler", err, w)
		return
	}

	// Sort the data to be returned
	returnData := map[string]interface{}{
		"data": data,
	}
	json.NewEncoder(w).Encode(returnData)
}

// spaceshipUpdateHandler - Update a spaceships
func spaceshipUpdateHandler(w http.ResponseWriter, r *http.Request) {

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get url param
	// vars := mux.Vars(r)

	// Create spaceship object
	spaceship := core.Spaceship{}
	err := json.NewDecoder(r.Body).Decode(&spaceship)
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipCreateHandler", err, w)
		return
	}

	// Update
	// err := core.UpdateSpaceship(vars["id"], spaceship.Name, spaceship.Class, spaceship.Crew, spaceship.Image, spacespaceship.Value)

	// Sort the data to be returned
	returnData := map[string]interface{}{
		"success": true,
	}
	json.NewEncoder(w).Encode(returnData)
}

// spaceshipDeleteHandler - List spaceships based on (name, class, status)
func spaceshipDeleteHandler(w http.ResponseWriter, r *http.Request) {

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get url param
	vars := mux.Vars(r)

	//validate param
	v := validator.New()
	err := v.Var(vars["id"], "required,max=999999,numeric")
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipDeleteHandler", errors.New("Must be numeric and less than 999999"), w)
		return
	}

	// convert from string to int
	spaceshipID, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipDeleteHandler", err, w)
		return
	}

	// Delete
	err = core.DeleteSpaceship(spaceshipID)
	if err != nil {
		utils.NewAPIError("spaceshipHandlers", "spaceshipDeleteHandler", err, w)
		return
	}

	// Sort the data to be returned
	returnData := map[string]interface{}{
		"success": true,
	}
	json.NewEncoder(w).Encode(returnData)
}
