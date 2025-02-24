package models

type Customer struct {
	CustomerID string `json:"customer_id" gorm:"primary_key; column:id"`
	Name       string `json:"name" gorm:"column:customer_name; type:varchar(100);"`
	Email      string `json:"email" gorm:"column:customer_email; type:varchar(255);"`
	Phone      string `json:"phone" gorm:"column:customer_phone; type:varchar(20);"`
	Address    string `json:"address" gorm:"column:customer_address; type:varchar(255);"`
	LoyaltyPts int    `json:"loyalty_points" gorm:"column:loyalty_pts; type:int(11);"`
}
