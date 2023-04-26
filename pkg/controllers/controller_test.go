package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateStudent(t *testing.T) {
	data, _ := ioutil.ReadFile("create.json")
	req := httptest.NewRequest("POST", "http://localhost:9090/student", bytes.NewReader(data))
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	CreateStudent(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}
	expected := `{"ID":1,"name":"ABHISHEK E V","Age":21}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}

func TestGetStudent(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:9090/student", nil)
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	GetStudent(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	expected := `[{"ID":1,"name":"ABHISHEK E V","Age":21}]`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}

func TestGetStudentByID(t *testing.T) {

	req := httptest.NewRequest("GET", "http://localhost:9090/student/studentID", nil)
	req = mux.SetURLVars(req, map[string]string{"studentID": "1"})
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	GetStudentById(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	expected := `{"ID":1,"name":"ABHISHEK E V","Age":21}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}

func TestUpdateStudent(t *testing.T) {
	//reader := strings.NewReader("name=abhishek", "age=21")
	data, _ := ioutil.ReadFile("update.json")
	req := httptest.NewRequest("PUT", "http://localhost:9090/student/studentID", bytes.NewReader(data))

	req = mux.SetURLVars(req, map[string]string{"studentID": "1"})
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	UpdateStudent(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	expected := `{"ID":1,"name":"ABHISHEK E V","Age":21}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}

func TestDeleteStudent(t *testing.T) {

	req := httptest.NewRequest("DELETE", "http://localhost:9090/student/studentID", nil)
	req = mux.SetURLVars(req, map[string]string{"studentID": "1"})
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	DeleteStudent(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}

	expected := `{"ID":1,"name":"ABHISHEK E V","Age":21}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}
