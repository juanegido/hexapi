package internal

import "context"

// CourseRepository is the interface that wraps the basic methods to persist a course.
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

// Course is the domain entity that represents a course.
type Course struct {
	id       string
	name     string
	duration string
}

// NewCourse creates a new course.
func NewCourse(id, name, duration string) Course {
	return Course{
		id:       id,
		name:     name,
		duration: duration,
	}
}

// ID Getter
func (c Course) ID() string {
	return c.id
}

// Name Getter
func (c Course) Name() string {
	return c.name
}

// Duration Getter
func (c Course) Duration() string {
	return c.duration
}
