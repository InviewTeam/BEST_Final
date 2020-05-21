package database

import (
	"log"
)

func Create(es Equipment) {
	GetDB().Create(&Equipment{Name: es.Name, Description: es.Description, Price: es.Price})
}

func Update(eq Equipment, ID string) error {
	var equipment Equipment
	err := GetDB().Where("ID = ?", ID).First(&equipment).Error
	if err != nil {
		log.Printf("Not found\n")
		return err
	}

	equipment.Name = eq.Name
	equipment.Description = eq.Description
	equipment.Price = eq.Price

	GetDB().Save(&equipment)
	return nil
}

func Delete(ID string) error {
	var eq Equipment
	err := GetDB().Where("ID = ?", ID).First(&eq).Error
	if err != nil {
		log.Printf("Not found\n")
		return err
	}

	err = GetDB().Delete(eq).Error
	if err != nil {
		log.Printf("Cannot delete\n")
		return err
	}
	return nil
}

func Get() []Equipment {
	res := []Equipment{}
	GetDB().Find(&res)
	return res
}
