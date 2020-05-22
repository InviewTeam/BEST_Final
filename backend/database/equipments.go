package database

import (
	"log"
)

func Create(es Equipment) {
	GetDB().Create(&Equipment{Title: es.Title, Description: es.Description, Price: es.Price, Image: es.Image})
}

func Update(eq Equipment, ID string) error {
	var equipment Equipment
	err := GetDB().Where("ID = ?", ID).First(&equipment).Error
	if err != nil {
		log.Printf("Not found\n")
		return err
	}

	equipment.Title = eq.Title
	equipment.Description = eq.Description
	equipment.Price = eq.Price
	equipment.Image = eq.Image

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
