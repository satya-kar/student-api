package main

import "sync"

type Student struct {
	StudentID string
	Name      string
	Age       int
	Class     string
	Subject   string
	IsDeleted bool
}

// In Memory Database
type StudentDatabase struct {
	Students map[string]Student // key -> StudentID , value -> Student
	mu       sync.Mutex         // mutex to handle concurrent read and write operations
}

var db StudentDatabase
