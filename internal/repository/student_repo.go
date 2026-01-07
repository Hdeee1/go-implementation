package repository

import (
	"database/sql"

	"github.com/Hdeee1/go-implementation/internal/model"
)

type StudentRepository interface {
	Create(student *model.Student) error
	GetAll() ([]*model.Student, error)
}

type studentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) StudentRepository {
	return &studentRepo{db: db}
}

func (r *studentRepo) Create(s *model.Student) error {
	query := `INSERT INTO student (name, age, date_of_birth) VALUES($1,$2, $3) RETURNING id`
	return r.db.QueryRow(query, s.Name, s.Age, s.DateOfBirth).Scan(&s.ID)
}

func (r *studentRepo) GetAll() ([]*model.Student, error) {
	rows, err := r.db.Query("SELECT id, name, age, date_os_birth FROM student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentList []*model.Student

	for rows.Next() {
		student := &model.Student{}
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.DateOfBirth)
		if err != nil {
			return nil, err
		}
		studentList = append(studentList, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return studentList, nil
}