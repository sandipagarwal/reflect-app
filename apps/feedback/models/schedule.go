package models

import (
	"time"

	"github.com/jinzhu/gorm"

	userModels "github.com/iReflect/reflect-app/apps/user/models"
)

// Schedule at which feedback events would be created for a team
type Schedule struct {
	gorm.Model
	Team          userModels.Team
	TeamID        uint      `gorm:"not null"`
	PeriodValue   uint      `gorm:"not null"`                   // consecutive period units, after which event need to be recreated
	PeriodUnit    string    `gorm:"type:varchar(15); not null"` // Unit of period, eg: year, week, days
	PeriodOffset  uint      `gorm:"default:1; not null"`        // Offset in days between the period end and actual event date
	FeedbackTitle string    `gorm:"type:varchar(255); not null"`
	ExpireInDays  uint      `gorm:"default:10; not null"`
	NextEventAt   time.Time `gorm:"not null"`
	Active        bool      `gorm:"default:true; not null"`
}
