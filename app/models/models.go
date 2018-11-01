package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Dob         time.Time
	PhoneNumber string `json:"phonenumber"`
	Email       string `json:"email"`
	Status      bool   `json:"status"`
}

func (e *Person) Disable() {
	e.Status = false
}

func (p *Person) Enable() {
	p.Status = true
}

// DBMigrate will create and migrate the tables , and then make the same relationship if necessary

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Person{})
	return db
}
