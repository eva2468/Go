package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/eva2468/student-management/pkg/models"
	"github.com/eva2468/student-management/pkg/utils"
	"github.com/gorilla/mux"
)

var NewStudent models.Student

func GetStudent(w http.ResponseWriter, r *http.Request) {

	NewStudent := models.GetStudent()
	res, _ := json.Marshal(NewStudent)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {

	CreateStudent := &models.Student{}
	utils.ParseBody(r, CreateStudent)
	b := CreateStudent.CreateStudent()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	student_id := vars["studentID"]
	ID, err := strconv.ParseInt(student_id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	stdentDetail, _ := models.GetStudentById(ID)
	res, _ := json.Marshal(stdentDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var updatestudent = &models.Student{}
	utils.ParseBody(r, updatestudent)
	vars := mux.Vars(r)
	student_id := vars["studentID"]
	ID, err := strconv.ParseInt(student_id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	studentdetails, db := models.GetStudentById(ID)
	if updatestudent.Name != "" {
		studentdetails.Name = updatestudent.Name
	}
	if updatestudent.Age != 0 {
		studentdetails.Age = updatestudent.Age
	}
	db.Save(&studentdetails)
	res, _ := json.Marshal(studentdetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	student_id := vars["studentID"]
	ID, err := strconv.ParseInt(student_id, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	stdentDetail, _ := models.DeleteStudent(ID)
	res, _ := json.Marshal(stdentDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
