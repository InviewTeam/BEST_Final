package database

import "log"

func CreateApp(app Application) {
	GetDB().Create(&ApplicationDB{Title: app.Title, Description: app.Description, Status: "Not ready", Executor: "admin"})
}

func UpdateApp(app ApplicationDB, ID string) error {
	var appDB ApplicationDB
	err := GetDB().Where("ID = ?", ID).First(&appDB).Error
	if err != nil {
		log.Printf("Not found\n")
		return err
	}

	appDB.Title = app.Title
	appDB.Description = app.Description
	appDB.Status = app.Status
	appDB.Executor = app.Executor

	GetDB().Save(&appDB)
	return nil
}

func DeleteApp(ID string) error {
	var app ApplicationDB
	err := GetDB().Where("ID = ?", ID).First(&app).Error
	if err != nil {
		log.Printf("Not found\n")
		return err
	}

	err = GetDB().Delete(app).Error
	if err != nil {
		log.Printf("Cannot delete\n")
		return err
	}
	return nil
}

func GetApps() []ApplicationDB {
	res := []ApplicationDB{}
	GetDB().Find(&res)
	return res
}
