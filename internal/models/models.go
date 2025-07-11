package models

import (
	"time"
)

// ================= USER =================
type User struct {
	ID          uint8   `gorm:"primaryKey"`
	PhoneNumber string  `gorm:"uniqueIndex;size:15"`
	Username    *string
	Password    string
	IsActive    bool    `gorm:"default:true"`
	IsStaff     bool    `gorm:"default:false"`
	DateJoined  time.Time `gorm:"autoCreateTime"`
	Role        string  `gorm:"size:20;default:user"`
	CreatedAt   time.Time
}

// ================= TEACHER =================
type Teacher struct {
	ID         uint8   `gorm:"primaryKey"`
	Fullname   *string
	TelegramID *string
	UserID     *uint8
	User       User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt  time.Time
}

// ================= GROUP =================
type Group struct {
	ID           uint8   `gorm:"primaryKey"`
	GroupName    *string
	Time         time.Time
	DurationHour *uint8
	WeekDay      string  `gorm:"size:15"`
	TeacherID    *uint8
	Teacher      Teacher `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time
}

// ================= STUDENT =================
type Student struct {
	ID          uint8   `gorm:"primaryKey"`
	Fullname    *string
	TelegramID  *string
	GroupID     *uint8
	Group       Group   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PhoneNumber *string
	IsActive    bool    `gorm:"default:false"`
	Age         *uint8
	CreatedAt   time.Time
}

// ================= PAYMENT =================
type Payment struct {
	ID         uint8   `gorm:"primaryKey"`
	StudentID  *uint8
	Student    Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaidAt     time.Time
	ValidUntil time.Time
}

// ================= MONTHLY FEE =================
type MonthlyFee struct {
	ID          uint8  `gorm:"primaryKey"`
	Amount      uint32
	Description string
}

// ================= STUDENT FEE =================
type StudentFee struct {
	ID           uint8       `gorm:"primaryKey"`
	StudentID    *uint8
	Student      Student     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MonthlyFeeID *uint8
	MonthlyFee   MonthlyFee  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// ================= ATTENDANCE =================
type Attendance struct {
	ID        uint8   `gorm:"primaryKey"`
	StudentID uint8
	Student   Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	GroupID   uint8
	Group     Group   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Date      time.Time
	Present   bool    `gorm:"default:false"`
}

// ================= LOCATION =================
type Location struct {
	ID          uint8   `gorm:"primaryKey"`
	Latitude    float32
	Longitude   float32
	Description string
}

// ================= STUDENT REGISTER CODE =================
type StudentCode struct {
	ID        uint8   `gorm:"primaryKey"`
	Code      string  `gorm:"size:6;uniqueIndex"`
	StudentID *uint8
	Student   Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsActive  bool    `gorm:"default:true"`
	CreatedAt time.Time
}
