package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	ID                  uint `gorm:"primarykey"`
	UserID              uint
	DetailTransactionID uint
	DetailTransaction   DetailTransaction
	Total               int
	Link                string
	TransactionStatus   string
	FraudStatus         string
	PaymentType         string
	Provider            string
	CreatedAt           *time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

type User struct {
	ID           uint `gorm:"primarykey"`
	Role         string
	Username     string
	Password     string
	Email        string `gorm:"unique"`
	PhoneNumber  string
	IsVerified   bool
	Transactions []Transaction
	CreatedAt    *time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type DetailTransaction struct {
	ID        uint `gorm:"primarykey"`
	ProductID uint
	Discount  int
	Subtotal  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Product struct {
	ID            uint   `gorm:"primarykey"`
	Name          string `gorm:"unique"`
	Description   string
	CategoryID    uint
	Category      Category
	Transaction   []DetailTransaction
	Price         int
	Stocks        int
	Sold          int
	SubCategoryID uint
	SubCategory   SubCategory
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type SubCategory struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `gorm:"unique"`
	Tax      int
	ImageURL string
	//Product []Product
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CSVModels struct {
	UserID      uint   `csv:"ID User"`
	PaymentType string `csv:"Metode Pembayaran"`
	Provider    string `csv:"Provider"`
	ProductID   uint   `csv:"ID Product"`
	Discount    int    `csv:"Discount"`
	Subtotal    int    `csv:"Sub Total"`
	Tax         int    `csv:"Pajak"`
	Total       int    `csv:"Total"`
}

type Voucher struct {
	ID        uint   `gorm:"primarykey"`
	Code      string `gorm:"unique"`
	Value     int
	Valid     time.Time
	CreatedAt *time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t *Transaction) ToCsVModel() CSVModels {
	return CSVModels{
		UserID:      t.UserID,
		PaymentType: t.PaymentType,
		Provider:    t.Provider,
		ProductID:   t.DetailTransaction.ProductID,
		Discount:    t.DetailTransaction.Discount,
		Subtotal:    t.DetailTransaction.Subtotal,
		Tax:         t.Total - (t.DetailTransaction.Subtotal - t.DetailTransaction.Discount),
		Total:       t.Total,
	}
}

func CsvModelList(tx []Transaction) []CSVModels {
	var out []CSVModels
	for x := range tx {
		item := tx[x].ToCsVModel()
		out = append(out, item)
	}
	return out
}
