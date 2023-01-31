package models

import (
	"github.com/eva2468/student-management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Student struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Age  int64  `json: "age" `
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}

func (s *Student) CreateStudent() *Student {

	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetStudent() []Student {
	var Student []Student
	db.Find(&Student)
	return Student

}

func GetStudentById(id int64) (*Student, *gorm.DB) {
	var getStudent Student
	db := db.Where("ID=?", id).Find(&getStudent)
	return &getStudent, db
}

func DeleteStudent(id int64) Student {

	var student Student
	db.Where("id=?", id).Delete(&student)
	return student

}
