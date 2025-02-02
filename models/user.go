// models/user.go
package models

type User struct {
	Base
	// Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Name string `json:"name"`
	// Password string `json:"-"`
	APIKEY string `gorm:"column:api_key;uniqueIndex;not null" json:"apikey"`
}
