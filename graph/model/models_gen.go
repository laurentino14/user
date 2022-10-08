// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthenticationInput struct {
	Token    *string `json:"token"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type Course struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Slug        string        `json:"slug"`
	Description string        `json:"description"`
	Image       string        `json:"image"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
	Lessons     []*Lesson     `json:"Lessons"`
	Steps       []*Step       `json:"Steps"`
	Enrollments []*Enrollment `json:"Enrollments"`
}

type Enrollment struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	UserID    string `json:"userId"`
	CourseID  string `json:"courseId"`
}

type GetUserAuthInput struct {
	Token *string `json:"token"`
}

type Lesson struct {
	ID        *string `json:"id"`
	Title     string  `json:"title"`
	Slug      string  `json:"slug"`
	Link      string  `json:"link"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	StepID    string  `json:"stepId"`
}

type NewCourse struct {
	Title       *string `json:"title"`
	Slug        *string `json:"slug"`
	Image       *string `json:"image"`
	Description *string `json:"description"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
}

type NewEnrollment struct {
	UserID   string `json:"userId"`
	CourseID string `json:"courseId"`
}

type NewLesson struct {
	Title  string `json:"title"`
	Slug   string `json:"slug"`
	Link   string `json:"link"`
	StepID string `json:"stepId"`
}

type NewStep struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Slug        string  `json:"slug"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
	CourseID    string  `json:"courseId"`
}

type NewUser struct {
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
	Cellphone *string `json:"cellphone"`
}

type Step struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	Lessons     []*Lesson `json:"lessons"`
	Course      *Course   `json:"Course"`
	CourseID    string    `json:"courseId"`
}

type User struct {
	ID         string        `json:"id"`
	Firstname  string        `json:"firstname"`
	Lastname   string        `json:"lastname"`
	Email      string        `json:"email"`
	Password   string        `json:"password"`
	Cellphone  string        `json:"cellphone"`
	TokenUser  string        `json:"token_user"`
	Enrollment []*Enrollment `json:"Enrollment"`
}

type UserAuthenticated struct {
	ID        *string `json:"id"`
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Email     *string `json:"email"`
	Cellphone *string `json:"cellphone"`
	TokenUser *string `json:"token_user"`
}
