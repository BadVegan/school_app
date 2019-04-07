package main

import (
	"database/sql"
	"flag"
	"github.com/RyczkoDawid/school_app/cmd/web/structs"
	"github.com/RyczkoDawid/school_app/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql" // New import
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", "localhost:5000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn := flag.String("dsn", "admin:admin8811@tcp(golangdb.cxr4pybuxjuw.eu-west-1.rds.amazonaws.com:3333)/school?parseTime=true", "MySQL data source name")

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &structs.Application{
		ErrorLog:      errorLog,
		InfoLog:       infoLog,
		Student:       &mysql.StudentModel{DB: db},
		SummaryLesson: &mysql.SummaryLessonModel{DB: db},
		Teacher:       &mysql.TeacherModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  routes(app),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Actual connections to the database are established lazily,
	// as and when needed for the first time.
	// So to verify that everything is set up correctly we need to use the
	// db.Ping() method to create a connection and check for any errors.
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
