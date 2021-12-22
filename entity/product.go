package entity

type Product struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"name"`
}

func (Product) TableName() string {
	return "Product"
}
