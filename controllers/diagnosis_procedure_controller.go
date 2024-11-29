package controllers

import (
	"backendMedicalRecord/models"
	"backendMedicalRecord/repository"
	"encoding/json"
	"fmt"
	"net/http"
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


func generateReportsOfProcedures(w http.ResponseWriter, r *http.Request) {
    var requestBody struct {
        IDProcedure string `json:"IDProcedure"`
        IDPatient   string `json:"IDPatient"`
    }

    var (
        isTimeType              bool
        totalDays               int
        numberOfTimeProcedures  int
        diagnosisProceduresArray []models.DiagnosisProcedure
        currentProcedure        models.Procedure
        averageAge              float64
        averageGender           float64
        successfulTreatments    int64
        isSuccessfullSimilarPatients bool
    )

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

    patients, err := repository.GetAllPatients()
    if err != nil {
        http.Error(w, "Error fetching patients", http.StatusInternalServerError)
        return
    }

    patientsMap := make(map[string]models.Patient)
    for _, patient := range patients {
        patientsMap[string(patient.ID.Hex())] = patient
    }

    // Find the specified procedure
    for _, pro := range procedures {
        if pro.ID.Hex() == requestBody.IDProcedure {
            isTimeType = pro.IsTimeType
            currentProcedure = pro
            break
        }
    }

    currentPatient, exists := patientsMap[requestBody.IDPatient]
    if !exists {
        http.Error(w, "Patient not found", http.StatusNotFound)
        return
    }

    diagnosisProceduresMap := make(map[string][]models.DiagnosisProcedure)
    for _, diagProc := range diagnosisProcedures {
        id := diagProc.ID.Hex() 
        if _, exists := diagnosisProceduresMap[id]; exists {
            diagnosisProceduresMap[id] = append(diagnosisProceduresMap[id], diagProc)
        } else {
            diagnosisProceduresMap[id] = []models.DiagnosisProcedure{diagProc}
        }
    }

    var patientMoreCommon struct {
        TotalDays int
        Count int
        IsSuccessful bool
    }

    patientMoreCommon:= make(map[string] patientMoreCommon)

    agePatient, err: = calculateAge(currentPatient.BirthDate)
    if err!= nil {
        http.Error(w, "Error calculating age", http.StatusInternalServerError)
        return
    }


    for id, diagnosisProcedures := range diagnosisProceduresMap {
        patient, exists := patientsMap[id]
        if !exists {
            continue 
        }


        patientMoreCommon[id].Count += checkPatientSimilarity(currentPatient, patient)
        if age< agePatient && age >agePatient-5 {
            patientMoreCommon[patient.ID.Hex()]++
        }

        
        
        for _, diagProc := range diagnosisProcedures {
            for _, procedure := range diagProc.Procedures {
                if isTimeType {
                    if procedure.StartAt != nil && procedure.EndAt != nil && procedure.EndAt.After(*procedure.StartAt) {
                        duration := procedure.EndAt.Sub(*procedure.StartAt)
                        totalDays += int(duration.Hours() / 24)
                        patientMoreCommon[id].TotalDays = int(duration.Hours() / 24)

                        diagnosisProceduresArray = append(diagnosisProceduresArray, diagProc)
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
                    }
                    } else {
                        if procedure.IsCompleted {
                            diagnosisProceduresArray = append(diagnosisProceduresArray, diagProc)
                            successfulTreatments++
                            patientMoreCommon.IsSuccessful = true
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
                    }else {
                        patientMoreCommon.IsSuccessful = false
                    }
                }
            }
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

    var ArrayPatientMoreCommon struct{
        patient   models.Patient  `json:"patient"`
        TotalDays int     `json:"totalDays"`
        Count    int     `json:"count"`
        IsSuccessful bool `json:"isSuccessful"`
    }

    var mostCommonPatient int
    var highestCountPatient int
    for idPatient, count := range patientMoreCommon {
        if patient > highestCount {
            mostCommonPatient = patientsMap[patient]
            highestCount = count
        }

        if 
    }

    averageTime = totalDays / numberOfTimeProcedures
    var arrayPatientMoreCommon [] ArrayPatientMoreCommon

    for patient, count:= range patientMoreCommon{
        if count == highestCount {
            arrayPatientMoreCommon = append(arrayPatientMoreCommon, ArrayPatientMoreCommon{
                patient:   patientsMap[patient],
                TotalDays: count.TotalDays,
                Count:    count.Count,
                IsSuccessful: count.IsSuccessful,
            })
        }

        if isTimeType{
            if arrayPatientMoreCommon[idPatient].TotalDays <= averageTime {
                successfulTreatments++
            }
        }

    }

    fmt.Printf("The patient with the most common age is: %s with %d cases.\n", mostCommonPatient.Name, highestCount)


    fmt.Printf("The most common diagnosis is: %s with %d cases.\n", mostCommon, highestCount)
    



    if numberOfTimeProcedures > 0 {
        averageAge /= float64(numberOfTimeProcedures)
        averageGender /= float64(numberOfTimeProcedures)
    }

    response := struct {
        DiagnosisProcedures []models.DiagnosisProcedure `json:"diagnosisProcedures"`
        IsTimeType          bool                        `json:"isTimeType"`
        AverageTime         int                         `json:"averageTime"`
        AverageAge          float64                     `json:"averageAge"`
        AverageGender       float64                     `json:"averageGender"`
    ArrayPatientMoreCommon     ArrayPatientMoreCommon   `json:"ArrayPatientMoreCommon"`
    SuccessfulTreatments     int64                       `json:"successfulTreatments"`
    
    }{
        DiagnosisProcedures: diagnosisProceduresArray,
        IsTimeType:          isTimeType,
        AverageTime:         averageTime,
        AverageAge:          averageAge,
        AverageGender:       averageGender,
        SimilarPatient:      arrayPatientMoreCommon,
        SuccessfulTreatments: successfulTreatments,
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }
}


func calculateAge(birthDateStr string) (int, error) {
	const layout = "2006-01-02" // Formato ISO 8601: YYYY-MM-DD

	birthDate, err := time.Parse(layout, birthDateStr)
	if err != nil {
		return 0, fmt.Errorf("invalid date format: %v", err)
	}

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

    similarityChecks := []func() bool{
        patient.HasAllergies,
        patient.HasBloodGlucose,
        patient.HasHighBloodPressure,
        patient.HasDiabetes,
        patient.HasSmoking,
        patient.HasAsthma,
        patient.HasEpilepsy,
        patient.HasLungCancer,
        patient.HasHepatitis,
        patient.HasHeartDisease,
        patient.HasKidneyDisease,
        patient.HasThyroidDisease,
        patient.HasCancer,
        patient.HasHormonalDisorders,
        patient.HasRenalDisease,
        patient.HasObesity,
        patient.HasDiabeticRetinopathy,
        patient.HasCerebralPalsy,
        patient.HasCongenitalAnomalies,
        patient.HasMentalDisorders,
        patient.HasNeurologicalDisorders,
        patient.HasPhysicalDisabilities,
        patient.HasLearningDisabilities,
        patient.HasDevelopmentalDisabilities,
    }

    currentPatientChecks := []func() bool{
        currentPatient.HasAllergies,
        currentPatient.HasBloodGlucose,
        currentPatient.HasHighBloodPressure,
        currentPatient.HasDiabetes,
        currentPatient.HasSmoking,
        currentPatient.HasAsthma,
        currentPatient.HasEpilepsy,
        currentPatient.HasLungCancer,
        currentPatient.HasHepatitis,
        currentPatient.HasHeartDisease,
        currentPatient.HasKidneyDisease,
        currentPatient.HasThyroidDisease,
        currentPatient.HasCancer,
        currentPatient.HasHormonalDisorders,
        currentPatient.HasRenalDisease,
        currentPatient.HasObesity,
        currentPatient.HasDiabeticRetinopathy,
        currentPatient.HasCerebralPalsy,
        currentPatient.HasCongenitalAnomalies,
        currentPatient.HasMentalDisorders,
        currentPatient.HasNeurologicalDisorders,
        currentPatient.HasPhysicalDisabilities,
        currentPatient.HasLearningDisabilities,
        currentPatient.HasDevelopmentalDisabilities,
    }

    for i := 0; i < len(similarityChecks); i++ {
        if similarityChecks[i]() == currentPatientChecks[i]() {
            similarityCount++
        }
    }

    return similarityCount
}