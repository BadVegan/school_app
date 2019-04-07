package mysql

import (
	"database/sql"
	"errors"
	"fmt"
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
	stmt := `INSERT INTO teachers (name, surname, email, phone)
    VALUES(?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, s.Name, s.Surname, s.Email, s.Phone)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TeacherModel) Get(id int) (*Teacher, error) {
	stmt := `SELECT * FROM teachers WHERE id = ?`

	s := &Teacher{}
	err := m.DB.QueryRow(stmt, id).Scan(&s.Id, &s.Name, &s.Surname, &s.Email, &s.Phone)
	if err == sql.ErrNoRows {
		return nil, ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	return s, nil
}

func (m *TeacherModel) GetAllTeachers() ([]*Teacher, error) {
	stmt := `SELECT * FROM teachers`
	return m.getAll(stmt)
}

func (m *TeacherModel) Update(t *Teacher) error {
	stmt := `UPDATE teachers SET name =?, surname=?, email = ?, phone = ? WHERE id = ?`
	_, err := m.DB.Exec(stmt, t.Name, t.Surname, t.Email, t.Phone, t.Id)

	if err == sql.ErrNoRows {
		return ErrNoRecord
	} else if err != nil {
		return err
	}

	fmt.Printf("%v", t)
	return nil
}

func (m *TeacherModel) getAll(stmt string, arg ...interface{}) ([]*Teacher, error) {

	rows, err := m.DB.Query(stmt, arg...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teacher []*Teacher

	for rows.Next() {
		s := &Teacher{}
		err = rows.Scan(&s.Id, &s.Name, &s.Surname, &s.Email, &s.Phone)
		if err != nil {
			return nil, err
		}
		teacher = append(teacher, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return teacher, nil
}
