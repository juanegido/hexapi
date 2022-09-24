package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	mooc "github.com/juanegido/hexapi/internal"
)

// CourseRepository is a MySQL mooc.CourseRepository implementation.
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.CourseRepository.
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Save implements the mooc.CourseRepository interface.
func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID().String(),
		Name:     course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

// GetAll implements the mooc.CourseRepository interface.
func (r *CourseRepository) GetAll(ctx context.Context) (courses []mooc.Course, err error) {
	courseSQLStruct := sqlbuilder.NewSelectBuilder()
	courseSQLStruct.Select("id", "name", "duration")
	courseSQLStruct.From(sqlCourseTable)

	sqlQuery, args := courseSQLStruct.Build()

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("error trying to get course on database: %v", err)
	}
	defer rows.Close()
	courses = []mooc.Course{}
	for rows.Next() {
		var sqlCourse sqlCourse
		err := rows.Scan(sqlCourse.ID, sqlCourse.Name, sqlCourse.Duration)
		if err != nil {
			return nil, err
		}
		course, err := mooc.NewCourse(sqlCourse.ID, sqlCourse.Name, sqlCourse.Duration)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
