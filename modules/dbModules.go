package modules

// 这里时需要与数据库交互的结构体

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

type DbCreateProject struct {
	TableId          int64  `gorm:"primaryKey,type:bigserial"`
	Hash             string `gorm:"type:varchar(1000)"`
	Time             string `gorm:"type:varchar(1000)"`
	Author           string `gorm:"type:varchar(1000)"`
	Email            string `gorm:"type:varchar(1000)"`
	ReleaseVersion   string `gorm:"type:varchar(1000)"`
	ProjectId        int    `gorm:"type:int"`
	PvLastCommitHash string `gorm:"type:varchar(1000)"`
}
