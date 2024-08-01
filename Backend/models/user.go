package models
import "github.com/lib/pq"


type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	Role  string `json:"role"`
	Permissions pq.StringArray `gorm:"type:text[]"`
	Departments pq.StringArray `gorm:"type:text[]"`
}



	

