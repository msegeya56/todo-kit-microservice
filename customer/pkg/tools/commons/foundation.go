package commons


import (
	"database/sql"
	"time"
)

type DeletedAt sql.NullTime

type Foundation struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Type        string
	Stage       string
	Status      int
	Data        []byte
	Document    []byte
	Events      []byte
	CreateAt    time.Time
	UpdatedAt   time.Time
	DeletedAt   DeletedAt `gorm:"index"`
	Timestamp   time.Time
}

type FoundationEntity struct {
	ID          uint 
	Name        string
	Description string
	Type        string
	Stage       string
	Status      int
	Data        []byte
	Document    []byte
	Events      []byte
	CreateAt    time.Time
	UpdatedAt   time.Time
	DeletedAt   DeletedAt 
	Timestamp   time.Time
}
