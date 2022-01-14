package models

import "time"

type Order struct {
	Id           uint       `gorm:"primaryKey" json:"orderId"`
	CustomerName string     `json:"customerName"`
	OrderedAt    *time.Time `json:"orderedAt"`
	Items        []Item     `gorm:"ForeignKey:OrderId" json:"items"`
}

func (order *Order) ValidateCreate() bool {
	if order.CustomerName == "" {
		return false
	}

	if order.OrderedAt == nil {
		return false
	}

	if len(order.Items) > 0 {
		for index := range order.Items {
			ok := order.Items[index].ValidateCreate()
			if !ok {
				return false
			}
		}
	}

	return true
}

func (order *Order) ValidateUpdate() bool {
	if order.CustomerName == "" {
		return false
	}

	if len(order.Items) > 0 {
		for index := range order.Items {
			ok := order.Items[index].ValidateUpdate()
			if !ok {
				return false
			}
		}
	}

	return true
}
