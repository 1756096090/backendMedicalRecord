package controllers

import (
	"backendMedicalRecord/models"
	"backendMedicalRecord/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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

func GetAllDiagnosisProceduresByID(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		IDPatient string `json:"IDPatient"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	diagnosisProcedures, err := repository.GetAllDiagnosisProceduresByID(requestBody.IDPatient)
	if err != nil {
		http.Error(w, "Error getting diagnosisProcedures", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(diagnosisProcedures)
}

func generateReportsOfProcedures(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		IDProcedure string `json:"IDProcedure"`
		IDPatient   string `json:"IDPatient"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if requestBody.IDProcedure == "" {
		http.Error(w, "IDProcedure is required", http.StatusBadRequest)
		return
	}

	diagnosisProcedures, err := repository.GetAllDiagnosisProcedures()
	if err != nil {
		http.Error(w, "Error fetching diagnosis procedures", http.StatusInternalServerError)
		return
	}

	procedures, err := repository.GetAllProcedures()
	if err != nil {
		http.Error(w, "Error fetching procedures", http.StatusInternalServerError)
		return
	}

    var (
		isTimeType            bool
		arrayIsTimeType       []models.DiagnosisProcedure
		totalDays            int
		numberOfTimeProcedures int
		diagnosisProceduresArray []models.DiagnosisProcedure
	)

    patients, err:= repository.GetAllPatients()
    if err != nil {
        http.Error(w, "Error fetching patients")
    }

    patientsMap:= make(map[string]models.Patient)
    for _,patient:= range patients{
        patientsMap[patient.ID.Hex()]
    }

	procedureMap := make(map[string]models.Procedure)
	for _, proc := range procedures {
		procedureMap[proc.ID.Hex()] = proc
	}

	for _, diagProc := range diagnosisProcedures {
		for _, procedure := range diagProc.Procedures {
			if procedure.IDProcedure != requestBody.IDProcedure {
				continue
			}

			proc, exists := procedureMap[requestBody.IDProcedure]
			if !exists {
				continue
			}

			diagnosisProceduresArray = append(diagnosisProceduresArray, diagProc)

			if proc.IsTimeType {
				isTimeType = true

				if procedure.StartAt != nil && procedure.EndAt != nil && procedure.EndAt.After(*procedure.StartAt) {
					duration := procedure.EndAt.Sub(*procedure.StartAt)
					totalDays += int(duration.Hours() / 24)
					numberOfTimeProcedures++
					arrayIsTimeType = append(arrayIsTimeType, diagProc)
				}
			}
		}
	}

	mediaTime := 0
	if numberOfTimeProcedures > 0 {
		mediaTime = totalDays / numberOfTimeProcedures
	}

	response := struct {
		DiagnosisProcedures []models.DiagnosisProcedure `json:"diagnosisProcedures"`
		IsTimeType          bool                        `json:"isTimeType"`
		MediaTime           int                         `json:"mediaTime"`
	}{
		DiagnosisProcedures: diagnosisProceduresArray,
		IsTimeType:          isTimeType,
		MediaTime:           mediaTime,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}