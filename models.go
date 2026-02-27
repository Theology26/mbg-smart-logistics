package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Role     string `gorm:"type:enum('vendor', 'guru');not null" json:"role"`

	IngredientQCs []IngredientQC `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"ingredient_qcs"`
	DeliveryPlans []DeliveryPlan `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"delivery_plans"`

	Confirmations []Confirmation `gorm:"foreignKey:GuruID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"confirmations"`
}

type IngredientQC struct {
	gorm.Model
	UserID         uint   `gorm:"not null" json:"user_id"`
	IngredientName string `gorm:"type:varchar(100);not null" json:"ingredient_name"`
	PhotoURL       string `gorm:"type:varchar(255);not null" json:"photo_url"`
	Status         string `gorm:"type:enum('SAFE', 'UNSAFE');not null" json:"status"`
	AiAnalysis     string `gorm:"type:text" json:"ai_analysis"`
}

type DeliveryPlan struct {
	gorm.Model
	UserID       uint      `gorm:"not null" json:"user_id"`
	MenuName     string    `gorm:"type:varchar(150);not null" json:"menu_name"`
	SchoolName   string    `gorm:"type:varchar(150);not null" json:"school_name"`
	DistanceKm   float64   `gorm:"not null" json:"distance_km"`
	TravelTimeHr float64   `gorm:"not null" json:"travel_time_hr"`
	Status       string    `gorm:"type:enum('APPROVED', 'REJECTED');not null" json:"status"`
	AiReason     string    `gorm:"type:text" json:"ai_reason"`
	DeliveryDate time.Time `gorm:"not null" json:"delivery_date"`

	Confirmation Confirmation `gorm:"foreignKey:DeliveryPlanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"confirmation"`
	Feedback     Feedback     `gorm:"foreignKey:DeliveryPlanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"feedback"`
}

type Confirmation struct {
	gorm.Model
	DeliveryPlanID uint      `gorm:"uniqueIndex;not null" json:"delivery_plan_id"`
	GuruID         uint      `gorm:"not null" json:"guru_id"`
	PhotoProof     string    `gorm:"type:varchar(255)" json:"photo_proof"`
	ReceivedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"received_at"`
}

type Feedback struct {
	gorm.Model
	DeliveryPlanID uint   `gorm:"uniqueIndex;not null" json:"delivery_plan_id"`
	Rating         int    `gorm:"not null" json:"rating"`
	Comment        string `gorm:"type:text" json:"comment"`
	IsFresh        bool   `gorm:"not null" json:"is_fresh"`
}
