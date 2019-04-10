package mysql

import (
	"database/sql"
	"gopkg.in/guregu/null.v3"
)

type ClassModel struct {
	DB *sql.DB
}

type Class struct {
	Id        int
	Name      string
	TeacherId null.Int
}

func (cm *ClassModel) Insert(c *Class) (int, error) {
	stmt := `INSERT INTO classes (name, teacher_id) VALUES(?, ?)`
	return Insert(stmt, cm.DB, c.Name, c.TeacherId)
}

func (cm *ClassModel) Get(id int) (*Class, error) {
	stmt := `SELECT id, name, teacher_id FROM classes WHERE id = ?`
	c := &Class{}
	err := Get(stmt, id, cm.DB, &c.Id, &c.Name, &c.TeacherId)
	return c, err
}

func (cm *ClassModel) Update(c *Class) error {
	stmt := `UPDATE classes SET name =?, teacher_id=? WHERE id = ?`
	return Update(stmt, cm.DB, &c.Name, &c.TeacherId, c.Id)

}

func (cm *ClassModel) GetClasses() ([]*Class, error) {
	stmt := `SELECT * FROM classes`
	return cm.getAll(stmt, cm.DB)
}

func (cm *ClassModel) getAll(stmt string, db *sql.DB, arg ...interface{}) ([]*Class, error) {

	rows, err := cm.DB.Query(stmt, arg...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []*Class

	for rows.Next() {
		c := &Class{}
		err = rows.Scan(&c.Id, &c.Name, &c.TeacherId)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return teachers, nil
}
