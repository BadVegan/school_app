package models

import (
	"errors"
	"gopkg.in/guregu/null.v3"
)

var ErrNoRecord = errors.New(" models: no matching record found")

type Student struct {
	ID      int
	Name    string
	Surname string
	Email   string
	Phone   string
	ClassID null.Int
}

type Class struct {
	ID       int
	Name     string
	Students []Student
}
