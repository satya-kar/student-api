package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Struct used to respond with unique student ID on creation of new record

type ResponseStudentID struct {
	StudentID string
}

// Handler for creation of new records -> POST: /student/v1/students

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student Student

	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		logError("Unable to read request body", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(reqBody, &student); err != nil {
		logError("Unmarshal Failed", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student.StudentID = uuid.NewString()
	student.IsDeleted = false

	db.mu.Lock()
	db.Students[student.StudentID] = student
	db.mu.Unlock()

	logInfo("Student record created with StudentID: " + student.StudentID)
	w.WriteHeader(http.StatusCreated)

	var resID ResponseStudentID
	resID.StudentID = student.StudentID

	json.NewEncoder(w).Encode(resID)
}

// Handler for retrieval of specific records using student ID -> Get /student/v1/students/{studentID}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	db.mu.Lock()
	student, exists := db.Students[studentID]
	db.mu.Unlock()

	if !exists || student.IsDeleted {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(student)
}

// Handler for retrieval of all records in Database -> Get /student/v1/students

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()
	var result []Student
	for _, student := range db.Students {
		if !student.IsDeleted {
			result = append(result, student)
		}
	}

	json.NewEncoder(w).Encode(result)
}

// Handler for deletion of specific records using student ID -> Delete /student/v1/students/{studentID}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["studentID"]

	db.mu.Lock()
	student, exists := db.Students[studentID]
	if exists && !student.IsDeleted {
		student.IsDeleted = true
		db.Students[studentID] = student
		logInfo("Student record soft-deleted with StudentID: " + student.StudentID)
	}
	db.mu.Unlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
