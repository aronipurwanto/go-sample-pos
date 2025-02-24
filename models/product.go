package models

type Product struct {
	ProductID   string  `json:"product_id" gorm:"primaryKey;column:id"`
	Name        string  `json:"name" gorm:"column:product_name; length:255"`
	Description string  `json:"description" gorm:"column:product_description; length:255"`
	Price       float64 `json:"price" gorm:"column:product_price"`
	StockQty    int     `json:"stock_qty" gorm:"column:stock_qty"`
	Category    string  `json:"category" gorm:"column:product_category"`
	SKU         string  `json:"sku" gorm:"column:product_sku"`
	TaxRate     float64 `json:"tax_rate" gorm:"column:tax_rate"`
}

type ProductError struct {
	Product Product `json:"product"`
	Error   error   `json:"error"`
}
