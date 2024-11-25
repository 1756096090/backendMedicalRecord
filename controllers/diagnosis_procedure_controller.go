package controllers

import (
    "encoding/json"
    "net/http"
    "backendMedicalRecord/models"
    "backendMedicalRecord/repository"
	"github.com/gorilla/mux"
)

func CreateDiagnosisProcedure(w http.ResponseWriter, r *http.Request) {
    var diagnosisProcedure models.DiagnosisProcedure
    if err := json.NewDecoder(r.Body).Decode(&diagnosisProcedure); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    
    result, err := repository.CreateDiagnosisProcedure(diagnosisProcedure)
    if err != nil {
        http.Error(w, "Error creating diagnosisProcedure", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)
}

func GetDiagnosisProcedure(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    diagnosisProcedureID := vars["id"]

    diagnosisProcedure, err := repository.GetDiagnosisProcedureByID(diagnosisProcedureID)
    if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    if diagnosisProcedure == nil {
        http.Error(w, "DiagnosisProcedure not found", http.StatusNotFound)
        return
    }
    
    json.NewEncoder(w).Encode(diagnosisProcedure)
}

func UpdateDiagnosisProcedure(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    diagnosisProcedureID := vars["id"]

    var diagnosisProcedure models.DiagnosisProcedure
    err := json.NewDecoder(r.Body).Decode(&diagnosisProcedure)
    if err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    result, err := repository.UpdateDiagnosisProcedure(diagnosisProcedureID, diagnosisProcedure)
    if err != nil {
        http.Error(w, "Error updating diagnosisProcedure", http.StatusInternalServerError)
        return
    }

    if result.ModifiedCount == 0 {
        http.Error(w, "No hay cambios", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(result)
}

func DeleteDiagnosisProcedure(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    diagnosisProcedureID := vars["id"]
    
    result, err := repository.DeleteDiagnosisProcedure(diagnosisProcedureID)
    if err != nil {
        http.Error(w, "Error deleting diagnosisProcedure", http.StatusInternalServerError)
        return
    }

    if result.DeletedCount == 0 {
        http.Error(w, "DiagnosisProcedure not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent) 
}

func GetAllDiagnosisProcedures(w http.ResponseWriter, r *http.Request) {
    diagnosisProcedures, err := repository.GetAllDiagnosisProcedures()
    if err != nil {
        http.Error(w, "Error getting diagnosisProcedures", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(diagnosisProcedures)
}

func GetAllDiagnosisProceduresByID(w http.ResponseWriter, r *http.Request){
	 var requestBody struct {
        IDPatient string `json:"IDPatient"`
    }

    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    diagnosisProcedures, err := repository.GetAllDiagnosisProceduresByID(requestBody.IDPatient)
    if err!= nil {
        http.Error(w, "Error getting diagnosisProcedures", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(diagnosisProcedures)
}


