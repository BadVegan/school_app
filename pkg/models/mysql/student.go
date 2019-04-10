package mysql

import (
	"database/sql"
	"gopkg.in/guregu/null.v3"
)

type StudentModel struct {
	DB *sql.DB
}

type Student struct {
	ID      int
	Name    string
	Surname string
	Email   string
	Phone   string
	ClassID null.Int
}

func (m *StudentModel) Insert(s *Student) (int, error) {
	stmt := `INSERT INTO students (name, surname, email, phone, class_id) VALUES(?, ?, ?, ?, ?)`
	return Insert(stmt, m.DB, s.Name, s.Surname, s.Email, s.Phone, s.ClassID)
}

func (m *StudentModel) Get(id int) (*Student, error) {
	stmt := `SELECT id, name, surname, email, phone, class_id FROM students WHERE id = ?`
	s := &Student{}
	err := Get(stmt, id, m.DB, field(s)...)
	return s, err
}

func (m *StudentModel) Update(s *Student) error {
	stmt := `UPDATE students SET name =?, surname=?, email = ?, phone = ?, class_id = ? WHERE id = ?`
	return Update(stmt, m.DB, s.Name, s.Surname, s.Email, s.Phone, s.ClassID, s.ID)

}

func (m *StudentModel) GetAllStudents() ([]*Student, error) {
	stmt := `SELECT id, name, surname, email, phone, class_id FROM students`
	return m.getAll(stmt)
}

func (m *StudentModel) GetAllByClass(classID int) ([]*Student, error) {
	stmt := `SELECT id, name, surname, email, phone, class_id FROM students WHERE class_id = ?`
	return m.getAll(stmt, classID)
}

func (m *StudentModel) getAll(stmt string, arg ...interface{}) ([]*Student, error) {

	rows, err := m.DB.Query(stmt, arg...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []*Student

	for rows.Next() {
		s := &Student{}
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

func field(s *Student) []interface{} {
	return []interface{}{&s.ID, &s.Name, &s.Surname, &s.Email, &s.Phone, &s.ClassID}
}
