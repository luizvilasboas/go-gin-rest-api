package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero,required"`
	Age  int    `json:"age" validate:"nonzero,required"`
	CPF  string `json:"cpf" validate:"len=11,unique,regexp=^[0-9]*$"`
	RG   string `json:"rg" validate:"len=9,unique,regexp=^[0-9]*$"`
}

func ValidateStudentData(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}

	return nil
}
