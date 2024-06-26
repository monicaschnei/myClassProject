// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"
)

type Availability struct {
	ID          int64  `json:"id"`
	Date        string `json:"date"`
	Start       string `json:"start"`
	EndTime     string `json:"end_time"`
	IsAvailable bool   `json:"is_available"`
	UserID      int64  `json:"user_id"`
	Username    string `json:"username"`
}

type Phone struct {
	ID          int64     `json:"id"`
	CountryCode int32     `json:"country_code"`
	AreaCore    int32     `json:"area_core"`
	Number      int32     `json:"number"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int64     `json:"user_id"`
}

type ProfessionalInformation struct {
	ID                 int64     `json:"id"`
	ExperiencePeriod   string    `json:"experience_period"`
	OcupationArea      string    `json:"ocupation_area"`
	University         string    `json:"university"`
	GraduationDiploma  string    `json:"graduation_diploma"`
	Validate           bool      `json:"validate"`
	GraduationCountry  string    `json:"graduation_country"`
	GraduationCity     string    `json:"graduation_city"`
	GraduationState    string    `json:"graduation_state"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	ProfessionalUserID int64     `json:"professional_user_id"`
}

type ProfessionalUser struct {
	ID                int64     `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	Name              string    `json:"name"`
	Username          string    `json:"username"`
	Gender            string    `json:"gender"`
	Email             string    `json:"email"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	Cpf               string    `json:"cpf"`
	ImageID           int64     `json:"image_id"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	HashedPassword    string    `json:"hashed_password"`
	UpdatedAt         time.Time `json:"updated_at"`
	ClassHourPrice    string    `json:"class_hour_price"`
}

type ProfessionalUserSubjectMatter struct {
	ProfessionalUserID int64 `json:"professionalUser_id"`
	SubjectMatterID    int64 `json:"subjectMatter_id"`
}

type ResponsibleStudent struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	Name              string    `json:"name"`
	Gender            string    `json:"gender"`
	Email             string    `json:"email"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	Cpf               string    `json:"cpf"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	HashedPassword    string    `json:"hashed_password"`
}

type StudentUser struct {
	ID                   int64     `json:"id"`
	Username             string    `json:"username"`
	Name                 string    `json:"name"`
	DateOfBirth          time.Time `json:"date_of_birth"`
	Gender               string    `json:"gender"`
	CreatedAt            time.Time `json:"created_at"`
	ResponsibleStudentID int64     `json:"responsible_student_id"`
	UpdatedAt            time.Time `json:"updated_at"`
	PasswordChangedAt    time.Time `json:"password_changed_at"`
	HashedPassword       string    `json:"hashed_password"`
}

type SubjectMatter struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Category  string    `json:"category"`
	Abstract  string    `json:"abstract"`
}

type SubjectMatterClass struct {
	ID                 int64     `json:"id"`
	CreatedAt          time.Time `json:"created_at"`
	SubjectMatterID    int32     `json:"subjectMatter_id"`
	Durantion          int32     `json:"durantion"`
	EnrollmentDate     time.Time `json:"enrollment_date"`
	EnrollmentTime     time.Time `json:"enrollment_time"`
	Cancellation       bool      `json:"cancellation"`
	CancellationReason string    `json:"cancellation_reason"`
	StudentAttendence  bool      `json:"student_attendence"`
	StudyMaterial      string    `json:"study_material"`
	TestingExam        string    `json:"testing_exam"`
	ProfessionalUserID int64     `json:"professional_user_id"`
	StudentUserID      int64     `json:"student_user_id"`
}
