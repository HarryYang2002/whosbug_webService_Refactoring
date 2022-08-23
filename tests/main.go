package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	. "webService_Refactoring/modules"
)

func main() {
	dsn := "host=localhost user=postgres password= dbname= port= " +
		"sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		err.Error()
	}
	temp := CommitsTable{}
	db.Table("commits").Where("project_id = ?", 123).Find(&temp)
	// fmt.Println(temp.ReleaseTableId)

}
