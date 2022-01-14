package models

type Item struct {
	Id          uint   `gorm:"primaryKey" json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderId     uint   `json:"orderId"`
}

func (item *Item) ValidateCreate() bool {
	if item.ItemCode == "" {
		return false
	}

	if item.Description == "" {
		return false
	}

	if item.Quantity == 0 {
		return false
	}

	return true
}

func (item *Item) ValidateUpdate() bool {
	if item.Id == 0 {
		return false
	}

	if item.ItemCode == "" {
		return false
	}

	if item.Description == "" {
		return false
	}

	if item.Quantity == 0 {
		return false
	}

	return true
}
