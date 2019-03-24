package mysql

import (
	"database/sql"
	"github.com/RyczkoDawid/school_app/pkg/models"
)

type StudentModel struct {
	DB *sql.DB
}

func (m *StudentModel) Insert(s *models.Student) (int, error) {
	stmt := `INSERT INTO students (name, surname, email, phone, class_id)
    VALUES(?, ?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, s.Name, s.Surname, s.Email, s.Phone, s.ClassID)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *StudentModel) Get(id int) (*models.Student, error) {
	stmt := `SELECT id, name, surname, email, phone, class_id FROM students
    WHERE id = ?`

	s := &models.Student{}
	err := m.DB.QueryRow(stmt, id).Scan(field(s)...)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *StudentModel) GetAllStudents() ([]*models.Student, error) {
	stmt := `SELECT id, name, surname, email, phone, class_id FROM students`
	return m.getAll(stmt)
}

func (m *StudentModel) GetAllByClass(classID int) ([]*models.Student, error) {
	stmt := `SELECT id, name, surname, email, phone, class_id FROM students WHERE class_id = ?`
	return m.getAll(stmt, classID)
}

func (m *StudentModel) getAll(stmt string, arg ...interface{}) ([]*models.Student, error) {

	rows, err := m.DB.Query(stmt, arg...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []*models.Student{}

	for rows.Next() {
		s := &models.Student{}
		err = rows.Scan(field(s)...)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return students, nil
}

func field(s *models.Student) []interface{} {
	return []interface{}{&s.ID, &s.Name, &s.Surname, &s.Email, &s.Phone, &s.ClassID}
}
