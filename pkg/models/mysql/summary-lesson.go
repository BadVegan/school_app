package mysql

import (
	"database/sql"
	"errors"
	"time"
)

var ErrNoRecordSummaryLesson = errors.New(" models: no matching record found")

type SummaryLessonModel struct {
	DB *sql.DB
}

type SummaryLesson struct {
	Id      int
	Created time.Time
	Summary string
	ClassID int
}

func (m *SummaryLessonModel) Get(id int) (*SummaryLesson, error) {
	stmt := `SELECT *  FROM summary_lesson WHERE id = ?`
	sl := &SummaryLesson{}
	err := m.DB.QueryRow(stmt, id).Scan(&sl.Id, &sl.Created, &sl.Summary, &sl.ClassID)
	if err == sql.ErrNoRows {
		return nil, ErrNoRecordSummaryLesson
	} else if err != nil {
		return nil, err
	}
	return sl, nil
}
