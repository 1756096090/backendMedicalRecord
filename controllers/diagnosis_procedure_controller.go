package controllers

import (
	"backendMedicalRecord/models"
	"backendMedicalRecord/repository"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

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

func calculateAge(birthDate *time.Time) (int, error) {
	const layout = "2006-01-02" // Formato ISO 8601: YYYY-MM-DD


	now := time.Now()
	age := now.Year() - birthDate.Year()

	if now.Month() < birthDate.Month() || (now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}

	return age, nil
}

func checkPatientSimilarity(currentPatient, patient models.Patient) int {
	similarityCount := 0

	if patient.Gender == currentPatient.Gender {
		similarityCount++
	}

	// Create maps of boolean conditions for each patient
	patientConditions := map[string]bool{
		"HasAllergies":             patient.HasAllergies,
		"HasBloodGlucose":          patient.HasBloodGlucose,
		"HasDiabetes":              patient.HasDiabetes,
		"HasHeartDisease":          patient.HasHeartDisease,
		"HasNeurologicalDisorders": patient.HasNeurologicalDisorders,
		"HasBloodPressure":         patient.HasBloodPressure,
		"HasEndocrineDisorders":    patient.HasEndocrineDisorders,
	}

	currentPatientConditions := map[string]bool{
		"HasAllergies":             currentPatient.HasAllergies,
		"HasBloodGlucose":          currentPatient.HasBloodGlucose,
		"HasDiabetes":              currentPatient.HasDiabetes,
		"HasHeartDisease":          currentPatient.HasHeartDisease,
		"HasNeurologicalDisorders": currentPatient.HasNeurologicalDisorders,
		"HasBloodPressure":         currentPatient.HasBloodPressure,
		"HasEndocrineDisorders":    currentPatient.HasEndocrineDisorders,
	}

	// Compare conditions
	for condition, value := range patientConditions {
		if value == currentPatientConditions[condition] {
			similarityCount++
		}
	}

	return similarityCount
}

// models
var requestBodyReport struct {
	IDProcedure string `json:"IDProcedure"`
	IDPatient   string `json:"IDPatient"`
}

type MedicalData struct {
	DiagnosisProcedures []models.DiagnosisProcedure
	Procedures          []models.Procedure
	Patients            []models.Patient
}

func fetchMedicalData(w http.ResponseWriter) (*MedicalData, error) {
	diagnosisProcedures, err := repository.GetAllDiagnosisProcedures()
	if err != nil {
		http.Error(w, "Error fetching diagnosis procedures", http.StatusInternalServerError)
		return nil, fmt.Errorf("failed to fetch diagnosis procedures: %w", err)
	}

	procedures, err := repository.GetAllProcedures()
	if err != nil {
		http.Error(w, "Error fetching procedures", http.StatusInternalServerError)
		return nil, fmt.Errorf("failed to fetch procedures: %w", err)
	}

	patients, err := repository.GetAllPatients()
	if err != nil {
		http.Error(w, "Error fetching patients", http.StatusInternalServerError)
		return nil, fmt.Errorf("failed to fetch patients: %w", err)
	}

	return &MedicalData{
		DiagnosisProcedures: diagnosisProcedures,
		Procedures:          procedures,
		Patients:            patients,
	}, nil
}

func GenerateReportsOfProcedures(w http.ResponseWriter, r *http.Request) {

	var (
		isTimeType               bool
		totalDays                int
		numberOfTimeProcedures   int
		diagnosisProceduresArray []models.DiagnosisProcedure
		averageAge               float64
		averageTime              float64
		averageGender            float64
		successfulTreatments     int64
	)

	if err := json.NewDecoder(r.Body).Decode(&requestBodyReport); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	medicalData, err := fetchMedicalData(w)
	diagnosisProcedures := medicalData.DiagnosisProcedures
	patients := medicalData.Patients

	patientsMap := make(map[string]models.Patient)
	for _, patient := range patients {
		patientsMap[string(patient.ID.Hex())] = patient
	}

	diagnosisProceduresMap := make(map[string][]models.DiagnosisProcedure)
	for _, diagProc := range diagnosisProcedures {
		id := diagProc.IDPatient
		if _, exists := diagnosisProceduresMap[id]; exists {
			diagnosisProceduresMap[id] = append(diagnosisProceduresMap[id], diagProc)
		} else {
			diagnosisProceduresMap[id] = []models.DiagnosisProcedure{diagProc}
		}
	}

	currentPatient, exists := patientsMap[requestBodyReport.IDPatient]
	if !exists {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}

	type PatientMoreCommon struct {
		TotalDays    int
		Count        int
		IsSuccessful bool
		diagnosis    string
	}

	patientMoreCommon := make(map[string]PatientMoreCommon)

	agePatient, err := calculateAge(currentPatient.BirthDate)
	if err != nil {
		http.Error(w, "Error calculating age", http.StatusInternalServerError)
		return
	}

	for id, diagnosisProcedures := range diagnosisProceduresMap {

		var durationProcedure int
		if id == currentPatient.ID.Hex(){
			continue
		}
		patient, exists := patientsMap[id]
		if !exists {
			continue
		}
		age, err := calculateAge(patient.BirthDate)

		for i, diagProc := range diagnosisProcedures {
			index := id+"-"+diagProc.ID.Hex()+"-"+string(i)
			currentData := patientMoreCommon[index]
			fmt.Println("Key generated:", id+"-"+diagProc.ID.Hex()+"-",i) 
			log.Printf("info", id+"-"+diagProc.ID.Hex())
			for _, procedure := range diagProc.Procedures {
				if isTimeType {
					if procedure.StartAt != nil && procedure.EndAt != nil && procedure.EndAt.After(*procedure.StartAt) {

						duration := procedure.EndAt.Sub(*procedure.StartAt)
						durationProcedure = int(duration.Hours() / 24)
						totalDays += durationProcedure

						currentData.Count += checkPatientSimilarity(currentPatient, patient)
						if agePatient-5 < age && age < agePatient {
							currentData.Count++
						}

						currentData.TotalDays = int(duration.Hours() / 24)

						diagnosisProceduresArray = append(diagnosisProceduresArray, diagProc)
						if err != nil {
							http.Error(w, "Error calculating age", http.StatusInternalServerError)
							return
						}

						averageAge += float64(age)
						if patient.Gender {
							averageGender++
						}
						numberOfTimeProcedures++
					}
				} else {
					if procedure.IsCompleted {

						currentData.Count += checkPatientSimilarity(currentPatient, patient)
						if agePatient-5 < age && age < agePatient  +5{
							currentData.Count++
						}

						diagnosisProceduresArray = append(diagnosisProceduresArray, diagProc)
						successfulTreatments++
						currentData.IsSuccessful = true
						age, err := calculateAge(patient.BirthDate)
						if err != nil {
							http.Error(w, "Error calculating age", http.StatusInternalServerError)
							return
						}
						averageAge += float64(age)
						if patient.Gender {
							averageGender++
						}
						numberOfTimeProcedures++
					} else {
						currentData.IsSuccessful = false
					}
				}
			}

			patientMoreCommon[index] = currentData
		}
	}

	diagnosisCount := make(map[string]int)
	for _, diagProc := range diagnosisProceduresArray {
		diagnosisCount[diagProc.CodeUnderDiagnosis]++
	}

	var mostCommon string
	var highestCount int
	for diagnosis, count := range diagnosisCount {
		if count > highestCount {
			mostCommon = diagnosis
			highestCount = count
		}
	}

	type PatientMoreCommonModel struct {
		Patient      *models.Patient
		TotalDays    int
		Count        int
		IsSuccessful bool
	}

	var highestCountPatient int
	for _, stats := range patientMoreCommon {
		if stats.Count > highestCount {
			highestCountPatient = stats.Count
		}
	}

	averageTime = float64(totalDays) / float64(numberOfTimeProcedures)
	var arrayPatientMoreCommon []PatientMoreCommonModel

	for idPatientDiag, stats := range patientMoreCommon {
		var patientPtr *models.Patient
		idPatient, err := extractPatientID(idPatientDiag)
		if err != nil {
			http.Error(w, "Error extracting patient ID", http.StatusInternalServerError)
			return
		}

		if patient, exists := patientsMap[idPatient]; exists {
			patientPtr = &patient
			if stats.Count == highestCountPatient {

				arrayPatientMoreCommon = append(arrayPatientMoreCommon, PatientMoreCommonModel{
					Patient:      patientPtr,
					TotalDays:    stats.TotalDays,
					Count:        stats.Count,
					IsSuccessful: stats.IsSuccessful,
				})
			}

			if isTimeType {
				if float64(stats.TotalDays) <= averageTime {
					successfulTreatments++
				}
			}
		}

	}

	fmt.Printf("The most common diagnosis is: cases.\n", mostCommon, averageAge)

	if numberOfTimeProcedures > 0 {
		averageAge /= float64(numberOfTimeProcedures)
		averageGender /= float64(numberOfTimeProcedures)
	}

	response := struct {
		DiagnosisProcedures  []models.DiagnosisProcedure `json:"diagnosisProcedures"`
		IsTimeType           bool                        `json:"isTimeType"`
		AverageTime          float64                     `json:"averageTime"`
		AverageAge           float64                     `json:"averageAge"`
		AverageGender        float64                     `json:"averageGender"`
		SimilarPatients      []PatientMoreCommonModel    `json:"arrayPatientMoreCommon"`
		SuccessfulTreatments int64                       `json:"successfulTreatments"`
	}{
		DiagnosisProcedures:  diagnosisProceduresArray,
		IsTimeType:           isTimeType,
		AverageTime:          averageTime,
		AverageAge:           averageAge,
		AverageGender:        averageGender,
		SimilarPatients:      arrayPatientMoreCommon,
		SuccessfulTreatments: successfulTreatments,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func extractPatientID(combinedID string) (string, error) {
	r := regexp.MustCompile(`^(.+?)-`)

	matches := r.FindStringSubmatch(combinedID)
	if len(matches) < 2 {
		return "", fmt.Errorf("invalid ID format: %s", combinedID)
	}

	return matches[1], nil
}
