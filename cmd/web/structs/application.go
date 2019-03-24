package structs

import (
	"github.com/RyczkoDawid/school_app/pkg/models/mysql"
	"log"
)

type Application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Student  *mysql.StudentModel
}
