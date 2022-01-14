package service

import (
	"api/database"
	"api/models"
)

func InsertOrderAndItems(order *models.Order) error {
	err := database.GetDB().Debug().Create(order).Error

	if err != nil {
		return err
	}

	return nil
}

func updateOrder(order *models.Order) error {
	err := database.GetDB().Table("orders").Where("id = ?", order.Id).
		Updates(map[string]interface{}{
			"customer_name": order.CustomerName,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteOrderAndItems(orderId uint64) error {
	err := database.GetDB().
		Where("order_id = ?", orderId).
		Delete(models.Item{}).Error

	if err != nil {
		return err
	}

	err = database.GetDB().Delete(&models.Order{}, orderId).Error

	if err != nil {
		return err
	}

	return nil
}

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	db := database.GetDB()

	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func GetOrderById(orderId uint64) (*models.Order, error) {
	var order models.Order
	order.Id = uint(orderId)

	db := database.GetDB()

	err := db.Preload("Items").Find(&order).Error

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func UpdateOrderAndItems(order *models.Order) error {
	err := updateOrder(order)

	if err != nil {
		return err
	}

	if len(order.Items) > 0 {
		for index := range order.Items {
			order.Items[index].OrderId = order.Id
			err = UpdateItem(&order.Items[index])
			if err != nil {
				return err
			}
		}
	}

	return nil
}
