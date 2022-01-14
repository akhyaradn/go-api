package service

import (
	"api/database"
	"api/models"
)

func DeleteItem(item *models.Item) error {
	err := database.GetDB().Delete(item).Error

	if err != nil {
		return err
	}

	return nil
}

func BulkInsertItem(items *[]models.Item) error {
	err := database.GetDB().Create(items).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateItem(item *models.Item) error {
	err := database.GetDB().Save(item).Error

	if err != nil {
		return err
	}

	return nil
}
