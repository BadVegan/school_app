package mysql

import (
	"database/sql"
	"errors"
)

var ErrNoRecordTeacher = errors.New(" models: no matching record found")

type TeacherModel struct {
	DB *sql.DB
}

type Teacher struct {
	Id      int
	Name    string
	Surname string
	Email   string
	Phone   string
}

func (m *TeacherModel) Insert(s *Teacher) (int, error) {
	stmt := `INSERT INTO teachers (name, surname, email, phone) VALUES(?, ?, ?, ?)`
	return Insert(stmt, m.DB, &s.Name, &s.Surname, &s.Email, &s.Phone)
}

func (m *TeacherModel) Get(id int) (*Teacher, error) {
	stmt := `SELECT * FROM teachers WHERE id = ?`
	s := &Teacher{}

	err := Get(stmt, id, m.DB, &s.Id, &s.Name, &s.Surname, &s.Email, &s.Phone)
	return s, err
}

func (m *TeacherModel) Update(t *Teacher) error {
	stmt := `UPDATE teachers SET name =?, surname=?, email = ?, phone = ? WHERE id = ?`
	return Update(stmt, m.DB, t.Name, t.Surname, t.Email, t.Phone, t.Id)

}

func (m *TeacherModel) GetTeachers() ([]*Teacher, error) {
	stmt := `SELECT * FROM teachers`
	teacher, err := m.getAll(stmt, readRows)
	s, isTeacher := teacher.([]Teacher)
	if err != nil {
		return teacher, err
	}
	return teacher, err
}

func (m *TeacherModel) getAll(stmt string, reader rowsReader, params ...interface{}) ([]*interface{}, error) {
	rows, err := GetAll(stmt, m.DB, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teacher, err := reader(rows)

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return teacher, nil
}

type rowsReader func(rows *sql.Rows) ([]*interface{}, error)

func readRows(rows *sql.Rows) ([]*interface{}, error) {
	var teacher []*Teacher
	for rows.Next() {
		s := &Teacher{}
		err := rows.Scan(&s.Id, &s.Name, &s.Surname, &s.Email, &s.Phone)
		if err != nil {
			return nil, err
		}
		teacher = append(teacher, s)
	}
	return teacher, nil
}
