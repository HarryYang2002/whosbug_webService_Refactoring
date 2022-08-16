package modules

// 这里时需要与数据库交互的结构体

import (
	"github.com/google/uuid"
)

type UsersTable struct {
	UserId        uuid.UUID `gorm:"primaryKey"`
	UserName      string    `gorm:"type:varchar(150)"`
	UserFirstName string    `gorm:"type:varchar(150)"`
	UserLastName  string    `gorm:"type:varchar(150)"`
	UserEmail     string    `gorm:"type:varchar(255)"`
	UserToken     string    `gorm:"type:varchar(40)"`
	UserPassword  string    `gorm:"type:varchar(128)"`
}

type CommitsTable struct {
	TableId        uint64 `gorm:"primaryKey;auto_increment;type:bigserial"`
	Hash           string `gorm:"type:varchar(1000)"`
	Time           string `gorm:"type:varchar(1000)"`
	Author         string `gorm:"type:varchar(1000)"`
	Email          string `gorm:"type:varchar(1000)"`
	ReleaseTableId int    `gorm:"type:int"`
}

type ProjectsTable struct {
	TableId   uint64 `gorm:"primaryKey;auto_increment;type:serial"`
	ProjectId string `gorm:"type:varchar(200)"`
}

type ReleasesTable struct {
	TableId        uint64 `gorm:"primaryKey;auto_increment;type:serial"`
	ReleaseVersion string `gorm:"type:varchar(200)"`
	LastCommitHash string `gorm:"type:varchar(1000)"`
	ProjectId      string `gorm:"type:varchar(200)"`
}

type ObjectsTable struct {
	TableId         uint64 `gorm:"primaryKey;auto_increment;type:serial"`
	Parameters      string `gorm:"type:varchar(10000)"`
	Hash            string `gorm:"type:varchar(1000)"`
	StartLine       int    `gorm:"type:int"`
	EndLine         int    `gorm:"type:int"`
	ObjectPath      string `gorm:"type:varchar(1000)"`
	CurrentObjectId string `gorm:"type:varchar(1000)"`
	FatherObjectId  string `gorm:"type:varchar(1000)"`
	OldLine         int    `gorm:"type:int"`
	NewLine         int    `gorm:"type:int"`
	DeletedLine     int    `gorm:"type:int"`
	AddedLine       int    `gorm:"type:int"`
	ReleaseTableId  int    `gorm:"type:int"`
	CommitTableId   int    `gorm:"type:varchar(200)"`
}

type NodesTable struct {
	TableId          uint64  `gorm:"primaryKey;auto_increment;type:serial"`
	ObjectPath       string  `gorm:"type:varchar(1000)"`
	ObjectParameters string  `gorm:"type:varchar(10000)"`
	CurrentObjectId  string  `gorm:"type:varchar(1000)"`
	FatherObjectId   string  `gorm:"type:varchar(1000)"`
	OldConfidence    float64 `gorm:"type:double"`
	NewConfidence    float64 `gorm:"type:double"`
	CommitTableId    int     `gorm:"type:varchar(200)"`
	ObjectTableId    int     `gorm:"type:int"`
}
