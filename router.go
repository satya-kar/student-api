package main

import "github.com/gorilla/mux"

func StudentApiRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/student/v1/students", CreateStudent).Methods("POST")
	router.HandleFunc("/student/v1/students/{studentID}", DeleteStudent).Methods("DELETE")
	router.HandleFunc("/student/v1/students/{studentID}", GetStudent).Methods("GET")
	router.HandleFunc("/student/v1/students", GetAllStudents).Methods("GET")

	return router
}
