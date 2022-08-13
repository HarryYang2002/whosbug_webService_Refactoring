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

type UncountedDelete struct {
	TableId        int64  `gorm:"primaryKey,type:bigserial"`
	Parameters     string `gorm:"type:varchar(10000)"`
	StartLine      int    `gorm:"type:int"`
	EndLine        int    `gorm:"type:int"`
	ObjectPath     string `gorm:"type:varchar(1000)"`
	ObjectId       string `gorm:"type:varchar(1000)"`
	OldObjectId    string `gorm:"type:varchar(1000)"`
	OldLine        int    `gorm:"type:int"`
	NewLine        int    `gorm:"type:int"`
	DeleteLine     int    `gorm:"type:int"`
	AddedLine      int    `gorm:"type:int"`
	CommitId       int    `gorm:"type:int"`
	ReleaseVersion string `gorm:"type:varchar(200)"`
}

type ObjectType struct {
	TableId       int64   `gorm:"primaryKey,type:bigserial"`
	ObjectPath    string  `gorm:"type:varchar(1000)"`
	ObjectId      string  `gorm:"type:varchar(1000)"`
	ObjectInfo    string  `gorm:"type:jsonb"`
	OldConfidence float64 `gorm:"type:double"`
	NewConfidence float64 `gorm:"type:double"`
	StartLine     int     `gorm:"type:int"`
	EndLine       int     `gorm:"type:int"`
	Parameters    string  `gorm:"type:varchar(10000)"`
	OldObjectId   string  `gorm:"type:varchar(1000)"`
	CommitId      int     `gorm:"type:int"`
}

type ReleaseType struct {
}
