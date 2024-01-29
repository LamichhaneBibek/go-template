package models

type User struct {
	BaseModel
	Username string `gorm:"type:string;size:50;unique;not null;unique"`
	Email    string `gorm:"type:string;size:256;unique;not null;unique;default:null"`
	Password string `gorm:"type:string;size:256;not null"`
}
