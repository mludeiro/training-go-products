package entity

type Product struct {
	ID   uint   `gorm:"primarykey" json:"id"`
	Name string `gorm:"name" json:"name"`
}

func (Product) TableName() string {
	return "Product"
}
