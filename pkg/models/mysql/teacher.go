package mysql

import (
	"database/sql"
)

type DataBaseType interface {
	GetAll(stmt string, db *sql.DB, params ...interface{}) ([]*DataBaseType, error)
}

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
	return Insert(stmt, m.DB, s.Name, s.Surname, s.Email, s.Phone)
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
	return m.getAll(stmt, m.DB)
}

func (m *TeacherModel) getAll(stmt string, db *sql.DB, arg ...interface{}) ([]*Teacher, error) {

	rows, err := m.DB.Query(stmt, arg...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []*Teacher

	for rows.Next() {
		t := &Teacher{}
		err = rows.Scan(&t.Id, &t.Name, &t.Surname, &t.Email, &t.Phone)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return teachers, nil
}
