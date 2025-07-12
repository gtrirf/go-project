package models

import (
	"time"
)

// ================= USER =================
type User struct {
	ID             uint       `gorm:"primaryKey;autoIncrement;column:id"`
	Password       string     `gorm:"column:password;size:128;not null"`             // AbstractBaseUser
	LastLogin      *time.Time `gorm:"column:last_login"`                             // AbstractBaseUser
	IsSuperuser    bool       `gorm:"column:is_superuser;default:false"`             // PermissionsMixin
	PhoneNumber    string     `gorm:"column:phone_number;size:15;unique;not null"`   // Custom
	Username       *string    `gorm:"column:username;size:150"`                      // Custom
	IsStaff        bool       `gorm:"column:is_staff;default:false"`                 // Custom
	IsActive       bool       `gorm:"column:is_active;default:true"`                 // AbstractBaseUser
	DateJoined     time.Time  `gorm:"column:date_joined;autoCreateTime"`             // Custom
	Role           string     `gorm:"column:role;size:20;default:user"`              // Custom

	// NOTE: Permissions ManyToMany table "users_custom_user_permissions" is not explicitly modeled here.
	// You can manage it via raw SQL or define another model if needed.
}

func (User) TableName() string {
	return "users"
}

type Permission struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	Codename    string `gorm:"size:100"`
	ContentTypeID uint `gorm:"column:content_type_id"`
}

func (Permission) TableName() string {
	return "auth_permission"
}

type ContentType struct {
	ID        uint   `gorm:"primaryKey"`
	AppLabel  string `gorm:"size:100"`
	Model     string `gorm:"size:100"`
}

func (ContentType) TableName() string {
	return "django_content_type"
}

type UserPermission struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint `gorm:"column:user_id"`
	PermissionID uint `gorm:"column:permission_id"`
}

func (UserPermission) TableName() string {
	return "users_custom_user_permissions"
}

// ================= TEACHER =================
type Teacher struct {
	ID         uint8    `gorm:"primaryKey;column:id"`
	Fullname   *string  `gorm:"column:fullname"`
	TelegramID *string  `gorm:"column:telegram_id"`
	UserID     *uint32  `gorm:"column:user_id"`
	User       *User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Teacher) TableName() string {
	return "teachers"
}

// ================= GROUP =================
type Group struct {
	ID           uint8     `gorm:"primaryKey;column:id"`
	GroupName    *string   `gorm:"column:group_name"`
	Time         time.Time `gorm:"column:time"`
	DurationHour *uint8    `gorm:"column:duration_hour"`
	WeekDay      string    `gorm:"column:week_day;size:15"`
	TeacherID    *uint8    `gorm:"column:teacher_id"`
	Teacher      *Teacher  `gorm:"foreignKey:TeacherID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Group) TableName() string {
	return "groups"
}

// ================= STUDENT =================
type Student struct {
	ID          uint32    `gorm:"primaryKey;column:id"`
	Fullname    *string   `gorm:"column:fullname"`
	TelegramID  *string   `gorm:"column:telegram_id"`
	GroupID     *uint8    `gorm:"column:group_id"`
	Group       *Group    `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PhoneNumber *string   `gorm:"column:phone_number;size:15"`
	IsActive    bool      `gorm:"column:is_active;default:false"`
	Age         *uint8    `gorm:"column:age"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (Student) TableName() string {
	return "students"
}

// ================= PAYMENT =================
type Payment struct {
	ID         uint32    `gorm:"primaryKey;autoIncrement;column:id"`
	StudentID  *uint32   `gorm:"column:student_id"`
	Student    *Student  `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaidAt     time.Time `gorm:"column:paid_at"`
	ValidUntil time.Time `gorm:"column:valid_until"`
}

func (Payment) TableName() string {
	return "payments"
}

// ================= MONTHLY FEE =================
type MonthlyFee struct {
	ID          uint32 `gorm:"primaryKey;autoIncrement;column:id"`
	Amount      uint32 `gorm:"column:amount"`
	Description string `gorm:"column:description"`
}

func (MonthlyFee) TableName() string {
	return "monthly_fees"
}

// ================= STUDENT FEE =================
type StudentFee struct {
	ID           uint32      `gorm:"primaryKey;autoIncrement;column:id"`
	StudentID    *uint32     `gorm:"column:student_id"`
	Student      *Student    `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MonthlyFeeID *uint32     `gorm:"column:monthly_fee_id"`
	MonthlyFee   *MonthlyFee `gorm:"foreignKey:MonthlyFeeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (StudentFee) TableName() string {
	return "student_fees"
}

// ================= ATTENDANCE =================
type Attendance struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement;column:id"`
	StudentID *uint32   `gorm:"column:student_id"`
	Student   *Student  `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GroupID   *uint8    `gorm:"column:group_id"`
	Group     *Group    `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Date      time.Time `gorm:"column:date;autoCreateTime"`
	Present   bool      `gorm:"column:present;default:false"`
}

func (Attendance) TableName() string {
	return "attendances"
}

// ================= LOCATION =================
type Location struct {
	ID          uint8   `gorm:"primaryKey;column:id"`
	Latitude    float32 `gorm:"column:latitude"`
	Longitude   float32 `gorm:"column:longitude"`
	Description string  `gorm:"column:description"`
}

func (Location) TableName() string {
	return "locations"
}

// ================= STUDENT REGISTER CODE =================
type StudentCode struct {
	ID        uint8     `gorm:"primaryKey;column:id"`
	Code      string    `gorm:"column:code;size:6;uniqueIndex"`
	StudentID *uint8    `gorm:"column:student_id"`
	Student   *Student  `gorm:"foreignKey:StudentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsActive  bool      `gorm:"column:is_active;default:true"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (StudentCode) TableName() string {
	return "student_codes"
}