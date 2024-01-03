package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
