package modules

//createUser
import (
	"github.com/google/uuid"
)

type DbCreateUser struct {
	UserId        uuid.UUID `gorm:"primaryKey"`
	UserName      string    `gorm:"type:varchar(150)"`
	UserFirstName string    `gorm:"type:varchar(150)"`
	UserLastName  string    `gorm:"type:varchar(150)"`
	UserEmail     string    `gorm:"type:varchar(255)"`
	UserToken     string    `gorm:"type:varchar(40)"`
	UserPassword  string
}
